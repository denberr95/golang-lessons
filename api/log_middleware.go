package api

import (
	"bytes"
	"io"
	"main/util"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	isWritable bool
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	if w.isWritable {
		w.body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}

func logMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logBaseEntry := createBaseLogEntry(c)
		logRequestEntry := logBaseEntry.WithField(util.LogrusFieldHttpDirection, util.HTTPRequest)
		logResponseEntry := logBaseEntry.WithField(util.LogrusFieldHttpDirection, util.HTTPResponse)

		logHeaders(logRequestEntry, util.FormatHttpHeaders(c.Request.Header))

		if isMultipartForm(c) {
			logMultipartRequest(logRequestEntry, c)
		} else {
			logStandardRequestBody(logRequestEntry, c)
		}

		bw := wrapWriter(c)
		c.Next()
		logResponse(logResponseEntry, c, bw)
	}
}

func createBaseLogEntry(c *gin.Context) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		util.LogrusFieldHttpMethod: c.Request.Method,
		util.LogrusFieldHttpURL:    c.Request.URL.String(),
	})
}

func logHeaders(entry *logrus.Entry, headers string) {
	entry.WithField(util.LogrusFieldHttpType, util.HTTPHeaders).Debugf("%s", headers)
}

func isMultipartForm(c *gin.Context) bool {
	return strings.HasPrefix(c.GetHeader(util.HeaderContentType), "multipart/form-data")
}

func logMultipartRequest(entry *logrus.Entry, c *gin.Context) {
	if err := c.Request.ParseMultipartForm(webServerConfig.HTTP.GetMaxMultipartMemoryMB()); err == nil && c.Request.MultipartForm != nil {
		entry.WithField(util.LogrusFieldHttpType, util.HTTPBody).Debugf("%s", util.FormatMultipartForm(c.Request.MultipartForm))
		entry.WithField(util.LogrusFieldHttpType, util.HTTPAttachment).Debugf("%s", util.FormatMultipartFiles(c.Request.MultipartForm))
	} else {
		entry.WithField(util.LogrusFieldHttpType, util.HTTPBody).Debugf("Errore durante la lettura del form: %v", err)
	}
}

func logStandardRequestBody(entry *logrus.Entry, c *gin.Context) {
	if c.Request.Body != nil {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		if len(bodyBytes) > 0 {
			entry.WithField(util.LogrusFieldHttpType, util.HTTPBody).Debugf("%s", string(bodyBytes))
		} else {
			entry.WithField(util.LogrusFieldHttpType, util.HTTPBody).Debug("Body vuoto")
		}
	}
}

func wrapWriter(c *gin.Context) *bodyWriter {
	bw := &bodyWriter{
		ResponseWriter: c.Writer,
		body:           bytes.NewBufferString(util.EmptyString),
		isWritable:     true,
	}
	c.Writer = bw
	return bw
}

func logResponse(entry *logrus.Entry, c *gin.Context, bw *bodyWriter) {
	if transferEncoding := c.Writer.Header().Get(util.HeaderTransferEncoding); strings.Contains(strings.ToLower(transferEncoding), "chunked") {
		bw.isWritable = false
	}

	logHeaders(entry, util.FormatHttpHeaders(c.Writer.Header()))

	if !bw.isWritable {
		entry.WithFields(logrus.Fields{
			util.LogrusFieldHttpType:   util.HTTPBody,
			util.LogrusFieldHttpStatus: c.Writer.Status(),
		}).Debug("Contenuto streaming/chunked non loggato")
		return
	}

	contentType := c.Writer.Header().Get(util.HeaderContentType)
	if strings.HasPrefix(contentType, "multipart/") {
		entry.WithFields(logrus.Fields{
			util.LogrusFieldHttpType:   util.HTTPAttachment,
			util.LogrusFieldHttpStatus: c.Writer.Status(),
		}).Debugf("%s", parseMultipartResponseFiles(contentType, bw.body.Bytes()))
	} else {
		entry.WithFields(logrus.Fields{
			util.LogrusFieldHttpType:   util.HTTPBody,
			util.LogrusFieldHttpStatus: c.Writer.Status(),
		}).Debugf("%s", bw.body.String())
	}
}

func parseMultipartResponseFiles(contentType string, body []byte) string {
	var sb strings.Builder
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil || params["boundary"] == "" {
		return "No multipart boundary"
	}

	reader := multipart.NewReader(bytes.NewReader(body), params["boundary"])

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "Errore lettura multipart response"
		}
		cd := part.Header.Get(util.HeaderContentDisposition)
		_, params, err := mime.ParseMediaType(cd)
		if err != nil {
			continue
		}
		if filename := params["filename"]; filename != "" {
			sb.WriteString(filename + util.Semicolon)
		}
		part.Close()
	}
	return sb.String()
}

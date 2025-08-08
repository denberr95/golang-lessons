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
		var bodyBytes []byte

		logBaseEntry := log.WithFields(logrus.Fields{
			util.LogrusFieldHttpMethod: c.Request.Method,
			util.LogrusFieldHttpURL:    c.Request.URL.String(),
		})

		logRequestEntry := logBaseEntry.WithFields(logrus.Fields{
			util.LogrusFieldHttpDirection: util.HTTPRequest,
		})

		logResponseEntry := logBaseEntry.WithFields(logrus.Fields{
			util.LogrusFieldHttpDirection: util.HTTPResponse,
		})

		logRequestEntry.WithFields(logrus.Fields{
			util.LogrusFieldHttpType: util.HTTPHeaders,
		}).Debugf("%v", util.FormatHttpHeaders(c.Request.Header))

		contentType := c.GetHeader(util.HeaderContentType)

		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := c.Request.ParseMultipartForm(cfg.GetMaxHeaderSizeMB()); err == nil && c.Request.MultipartForm != nil {

				logRequestEntry.WithFields(logrus.Fields{
					util.LogrusFieldHttpType: util.HTTPBody,
				}).Debugf("%s", util.FormatMultipartForm(c.Request.MultipartForm))

				logRequestEntry.WithFields(logrus.Fields{
					util.LogrusFieldHttpType: util.HTTPAttachment,
				}).Debugf("%s", util.FormatMultipartFiles(c.Request.MultipartForm))

			} else {
				logRequestEntry.WithFields(logrus.Fields{
					util.LogrusFieldHttpType: util.HTTPBody,
				}).Debugf("Errore durante la lettura del form: %v", err)
			}
		} else {
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			logRequestEntry.WithFields(logrus.Fields{
				util.LogrusFieldHttpType: util.HTTPBody,
			}).Debugf("%s", string(bodyBytes))
		}
		bodyWriter := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(util.EmptyString),
			isWritable:     true,
		}

		c.Writer = bodyWriter
		c.Next()

		if transferEncoding := c.Writer.Header().Get(util.HeaderTransferEncoding); strings.Contains(strings.ToLower(transferEncoding), "chunked") {
			bodyWriter.isWritable = false
		}

		logResponseEntry.WithFields(logrus.Fields{
			util.LogrusFieldHttpType: util.HTTPHeaders,
		}).Debugf("%s", util.FormatHttpHeaders(c.Writer.Header()))

		if !bodyWriter.isWritable {
			logResponseEntry.WithFields(logrus.Fields{
				util.LogrusFieldHttpType:   util.HTTPBody,
				util.LogrusFieldHttpStatus: c.Writer.Status(),
			}).Debug("Contenuto streaming/chunked non loggato")
		} else {
			contentType := c.Writer.Header().Get(util.HeaderContentType)
			if strings.HasPrefix(contentType, "multipart/") {
				logResponseEntry.WithFields(logrus.Fields{
					util.LogrusFieldHttpType:   util.HTTPAttachment,
					util.LogrusFieldHttpStatus: c.Writer.Status(),
				}).Debugf("%s", parseMultipartResponseFiles(contentType, bodyWriter.body.Bytes()))
			} else {
				logResponseEntry.WithFields(logrus.Fields{
					util.LogrusFieldHttpType:   util.HTTPBody,
					util.LogrusFieldHttpStatus: c.Writer.Status(),
				}).Debugf("%s", bodyWriter.body.String())
			}
		}
	}
}

func parseMultipartResponseFiles(contentType string, body []byte) string {
	var sb strings.Builder
	boundary := ""
	if _, params, err := mime.ParseMediaType(contentType); err == nil {
		boundary = params["boundary"]
	}
	if boundary == "" {
		return "No multipart boundary"
	}

	reader := multipart.NewReader(bytes.NewReader(body), boundary)

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
		filename := params["filename"]
		if filename != "" {
			sb.WriteString(filename)
			sb.WriteString(util.Semicolon)
		}
		part.Close()
	}
	return sb.String()
}

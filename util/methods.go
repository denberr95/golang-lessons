package util

import (
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strings"
)

func FormatHttpHeaders(headers http.Header) string {
	var builder strings.Builder
	var result string
	for key, value := range headers {
		builder.WriteString(key)
		builder.WriteString(Equals)
		builder.WriteString(strings.Join(value, Comma))
		builder.WriteString(Semicolon)
		builder.WriteString(Space)
	}
	result = strings.TrimSuffix(builder.String(), Semicolon+Space)
	return result
}

func FormatMultipartForm(multipartForm *multipart.Form) string {
	var builder strings.Builder
	var result string

	for key, value := range multipartForm.Value {
		builder.WriteString(key)
		builder.WriteString(Equals)
		builder.WriteString(strings.Join(value, Comma))
		builder.WriteString(Semicolon)
		builder.WriteString(Space)
	}
	result = strings.TrimSuffix(builder.String(), Semicolon+Space)
	return result
}

func FormatMultipartFiles(multipartForm *multipart.Form) string {
	var builder strings.Builder
	var result string

	for key, files := range multipartForm.File {
		for _, fileHeader := range files {
			builder.WriteString(key)
			builder.WriteString(Equals)
			builder.WriteString(fileHeader.Filename)
			builder.WriteString(Semicolon)
			builder.WriteString(Space)
		}
	}
	result = strings.TrimSuffix(builder.String(), Semicolon+Space)
	return result
}

func ConvertMegabitesToBytes(v int) int {
	return v * 1024 * 1024
}

func ShiftMB(v int) int64 {
	return int64(v) << 20
}

func PrintEnvVars() []string {
	result := os.Environ()
	sort.Strings(result)
	return result
}

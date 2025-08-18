package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func configureAccessLog() {
	if webServerConfig.Log.EnableAccessLog {
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] - [%s] - [%s] - [%s] - [%s] - [%d] - [%s] [%s] - [%s]\n",
				param.TimeStamp.Format(time.RFC3339),
				param.ClientIP,
				param.Method,
				param.Request.RequestURI,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
	}
}

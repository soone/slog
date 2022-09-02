package slog

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLoggerFormatter() gin.HandlerFunc {
	mode := os.Getenv("mode")
	if mode == ProduceMode {
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("{\"common\":\"GIN\", \"code\":\"%3d\", \"latency\":\"%v\", \"ip\":\"%s\", \"method\":\"%s\", \"path\":\"%v\", \"msg\":\"%s\", \"time\":\"%s\"}\n",
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
				param.ErrorMessage,
				param.TimeStamp.Format(time.RFC3339),
			)
		})
	}

	return gin.Logger()
}

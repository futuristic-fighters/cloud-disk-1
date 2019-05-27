package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//自定义gin的日志中间件
func (l *logger) GinLogger() gin.HandlerFunc {
	formatter := l.GinLoggerFormatter()

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		l.checkLoggerWriter()

		_, _ = fmt.Fprint(l.out, formatter(param))
	}
}

//自定义gin的loggerFormatter
func (l *logger) GinLoggerFormatter() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor, keysString, ErrWrap, KeysWrap string
		var keysBuffer bytes.Buffer

		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		b, _ := json.Marshal(param.Keys)
		_ = json.Indent(&keysBuffer, b, "", "    ")
		keysString = keysBuffer.String()
		keysString = strings.Replace(keysString, "{\n", "", 1)
		keysString = strings.Replace(keysString, "\n}", "", 1)
		if keysString == "null" {
			keysString = ""
		} else {
			KeysWrap = "\n"
		}

		if param.ErrorMessage != "" {
			ErrWrap = "\n    "
		}

		return fmt.Sprintf("[GIN] %s | %s %s %s %s |%s %d %s| %v | %s %s%s %s%s\n",
			param.TimeStamp.Format("15:04:05.999999999"),
			methodColor, param.Method, resetColor,
			param.Path,
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			ErrWrap,
			param.ErrorMessage,
			KeysWrap,
			keysString,
		)
	}
}

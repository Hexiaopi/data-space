package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/pkg/metrics"
)

// type ResponseWithRecorder struct {
// 	gin.ResponseWriter
// 	statusCode int
// 	body       bytes.Buffer
// }

// func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
// 	rec.ResponseWriter.WriteHeader(statusCode)
// 	rec.statusCode = statusCode
// }

// func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
// 	n, err = rec.ResponseWriter.Write(d)
// 	if err != nil {
// 		return
// 	}
// 	rec.body.Write(d)
// 	return
// }

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		metrics.HttpRequestConcurrency.WithLabelValues(path, method).Inc()
		defer func() {
			metrics.HttpRequestConcurrency.WithLabelValues(path, method).Dec()
		}()
		start := time.Now()
		c.Next()
		code := c.Writer.Status()
		duration := time.Since(start)
		metrics.HttpRequestLatency.WithLabelValues(path, method).Observe(duration.Seconds())
		metrics.HttpRequestCounter.WithLabelValues(path, method, strconv.Itoa(code)).Inc()
	}
}

package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	log "github.com/hexiaopi/data-space/pkg/logger"
)

type ResponseWithRecorder struct {
	gin.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)
	return
}

// Logger 日志记录
func Logger(logger log.Logger, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		start := time.Now()
		//记录请求包
		buf, _ := io.ReadAll(c.Request.Body)
		rdr := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr //rewrite

		logger.Info("receive request",
			log.String("path", c.Request.URL.Path),
			log.String("param", c.Request.URL.RawQuery),
			log.String("method", c.Request.Method),
			log.String("reqpkg", string(buf)),
		)

		//记录返回包
		wc := &ResponseWithRecorder{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}
		c.Writer = wc

		c.Next()

		defer func() { //日志记录扫尾工作
			logger.Info("done request",
				log.String("path", c.Request.URL.Path),
				log.Int("status", wc.statusCode),
				log.String("respkg", wc.body.String()),
				log.String("utm", time.Since(start).String()),
			)
		}()
	}
}

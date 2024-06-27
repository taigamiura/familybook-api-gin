package middlewares

import (
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
  "log"
  "time"
)

func RecordUaAndTime(c *gin.Context) {
  logger, err := zap.NewProduction()
  if err != nil {
    log.Fatal(err.Error())
  }
  defer logger.Sync()

  start := time.Now()
  ua := c.GetHeader("User-Agent")
  c.Next()
  end := time.Now()

  // logger.Info("request",
  //   zap.String("path", c.Request.URL.Path),
  //   zap.String("ua", ua),
  //   zap.Int("status", c.Writer.Status()),
  //   zap.Duration("latency", end.Sub(start)),
  // )
  var errorMessage string
	if len(c.Errors) > 0 {
		errorMessage = c.Errors.String()
	}
  logger.Info("request",
		zap.String("path", c.Request.URL.Path),
		zap.String("ua", ua),
		zap.String("method", c.Request.Method),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("latency", end.Sub(start)),
		zap.String("client_ip", c.ClientIP()),
		zap.String("error_message", errorMessage),
		zap.String("query", c.Request.URL.RawQuery),
	)
}
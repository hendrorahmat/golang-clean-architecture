package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"net/http"
	"time"
)

func TimeoutHandler(config config.HttpConf) func(c *gin.Context) {
	timeout := time.Duration(config.Timeout) * time.Second
	responseBodyTimeout := gin.H{
		"code":    http.StatusRequestTimeout,
		"message": "request timeout, response is sent from middleware",
	}

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.JSON(http.StatusRequestTimeout, responseBodyTimeout)
				c.Abort()
			}

			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

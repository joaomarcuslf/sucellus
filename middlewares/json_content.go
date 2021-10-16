package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.GetHeader("Content-Type")

		if contentType == "application/json" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}

	}
}

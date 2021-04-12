package middlewares

import (
	"net/http"
	"strings"

	"github.com/LarryKapija/shoppinglist_api/utils"
	"github.com/gin-gonic/gin"
)

func CacheControlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public")
		c.Header("Etag", utils.GenerateEtag(c.Request.URL.Path))
		if match := c.Request.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(utils.Etags[c.Request.URL.Path], match) {
				c.AbortWithStatus(http.StatusNotModified)
				return
			}
		}
		if match := c.Request.Header.Get("If-Match"); match != "" {
			if !strings.Contains(utils.Etags[c.Request.URL.Path], match) {
				c.AbortWithStatus(http.StatusConflict)
				return
			}
		}

		c.Next()
	}
}

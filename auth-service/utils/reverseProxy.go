package utils

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid target URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		c.Request.URL.Path = c.Param("proxyPath")

		// Ensure the path is correctly formatted
		proxyPath := c.Param("proxyPath")
		c.Request.URL.Path = proxyPath
		log.Println("Forwarding request to:", remote.ResolveReference(c.Request.URL))
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	routerGroups = append(routerGroups, registerEchoRouter)
}

func registerEchoRouter(g *gin.RouterGroup) {
	r := g.Group("/api/v1")
	r.GET("echo", func(c *gin.Context) {
		data := c.Query("data")
		c.String(http.StatusOK, data)
	})
}

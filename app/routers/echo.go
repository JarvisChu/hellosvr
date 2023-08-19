package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/hellosvr/config"
	"net/http"
)

func init() {
	routerGroups = append(routerGroups, registerEchoRouter)
}

func registerEchoRouter(g *gin.RouterGroup) {
	r := g.Group("/api/v1")
	r.GET("echo", func(c *gin.Context) {
		c.String(http.StatusOK, config.GetConfig().App.Echo)
	})
}

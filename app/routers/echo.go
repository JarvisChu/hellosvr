package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/hellosvr/config"
	"net/http"
	"net/http/httputil"
)

func init() {
	routerGroups = append(routerGroups, registerEchoRouter)
}

func registerEchoRouter(g *gin.RouterGroup) {
	r := g.Group("/api/v1")
	r.GET("echo", func(c *gin.Context) {

		// show raw request http headers
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		c.String(http.StatusOK, config.GetConfig().App.Echo)
	})
}

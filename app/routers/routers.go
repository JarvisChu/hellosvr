package routers

import "github.com/gin-gonic/gin"

var (
	routerGroups = make([]func(g *gin.RouterGroup), 0)
)

func InitRouters(r *gin.Engine) *gin.Engine {
	g := r.Group("")
	for _, f := range routerGroups {
		f(g)
	}

	return r
}

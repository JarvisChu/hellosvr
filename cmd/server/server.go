package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/hellosvr/app/routers"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Start a web server",
	Example: "hellosvr server -c etc/configs.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func run() error {
	fmt.Println(" server running...")

	r := gin.Default()
	routers.InitRouters(r)
	return r.Run(":8080")
}

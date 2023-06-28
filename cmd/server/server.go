package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/hellosvr/app/routers"
	"github.com/jarvischu/hellosvr/config"
	"github.com/spf13/cobra"
)

var (
	cfgPath string
	Cmd     = &cobra.Command{
		Use:     "server",
		Short:   "Start a web server",
		Example: "hellosvr server -c etc/configs.yml",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "hellosvr server -c etc/configs.yml")
}

func setup() error {
	return config.InitConfig(cfgPath)
}

func run() error {
	addr := fmt.Sprintf(":%v", config.GetConfig().App.Port)
	fmt.Printf("server is running on %v ...\n", addr)

	r := gin.Default()
	routers.InitRouters(r)
	return r.Run(addr)
}

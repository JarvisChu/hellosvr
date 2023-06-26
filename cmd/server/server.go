package server

import (
	"fmt"
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
	return nil
}

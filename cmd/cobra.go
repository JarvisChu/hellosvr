package cmd

import (
	"fmt"
	"github.com/jarvischu/hellosvr/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hellosvr",
	Short: "A example webserver for learning k8s",
	Run: func(cmd *cobra.Command, args []string) {
		help()
	},
}

func help() {
	fmt.Println("hellosvr is a example project for learning k8s, use -h for more information")
}

func init() {
	rootCmd.AddCommand(server.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

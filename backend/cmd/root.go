package main

import (
	"github.com/spf13/cobra"

	"github.com/hexiaopi/data-space/internal/app"
	"github.com/hexiaopi/data-space/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func main() {
	if err := config.Init("app.yaml"); err != nil {
		panic(err)
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

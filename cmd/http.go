package cmd

import (
	"gin-seed/gosdk"
	"gin-seed/machine/server"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func newApp() gosdk.Application {
	return gosdk.New(
		gosdk.WithName("gin-seed"),
		gosdk.WithVersion("0.0.1"),
		gosdk.WithInitRunnable(server.NewGinServer("gin-server", "gin")),
	)
}

var startServer = &cobra.Command{
	Use:   "app",
	Short: "Gin seed application",
	Run: func(cmd *cobra.Command, args []string) {
		app := newApp()
		serviceLogger := app.Logger("service")

		if err := app.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		if err := app.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	startServer.AddCommand(outEnvCmd)
	if err := startServer.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

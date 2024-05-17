package cmd

import (
	"app/adapter"
	"app/addons"
	"app/addons/database"
	"app/addons/server"
	"app/gosdk"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func newApp() gosdk.Application {
	return gosdk.New(
		gosdk.WithName("gin-seed"),
		gosdk.WithVersion("0.0.1"),
		gosdk.WithInitRunnable(server.NewGinServer(addons.GinServerName, addons.GinServerPrefix)),
		gosdk.WithInitRunnable(database.NewPgDatabase(addons.PgDatabaseName, addons.PgDatabasePrefix)),
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

		adapter.NewAdapter(app).Start()

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

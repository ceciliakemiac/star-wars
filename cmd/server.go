package cmd

import (
	"log"
	"os"
	"star-wars/api"
	"star-wars/helper"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",

	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	db := helper.ConnectDB()

	server, err := api.NewServer(os.Getenv("API_ADDR"), db)
	if err != nil {
		log.Fatal("startServer() Error Creating New Server: ", err)
	}

	if err = server.Run(); err != nil {
		log.Fatal("startServer() Error Running Server: ", err)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

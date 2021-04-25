package cmd

import (
	"star-wars/database"
	"star-wars/helper"

	"github.com/spf13/cobra"
)

// insertPlanetsCmd represents the insertPlanets command
var insertPlanetsCmd = &cobra.Command{
	Use: "insertPlanets",

	Run: func(cmd *cobra.Command, args []string) {
		insertPlanets()
	},
}

func insertPlanets() {
	db := helper.ConnectDB()
	database.InsertPlanets(db)
}

func init() {
	rootCmd.AddCommand(insertPlanetsCmd)
}

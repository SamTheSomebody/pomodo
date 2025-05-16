/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"pomodo/helpers"
	"pomodo/ui"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all of your tasks!`,
	Run: func(cmd *cobra.Command, args []string) {
		db := helpers.GetDBQueries()
		tasks, err := db.GetTasks(cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		for _, t := range tasks {
			ui.PrintTask(t)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

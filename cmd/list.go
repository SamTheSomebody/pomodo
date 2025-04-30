/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	db "pomodo/database"
	"pomodo/ui"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all of your tasks!`,
	Run: func(cmd *cobra.Command, args []string) {
		db := db.GetDBQueries()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

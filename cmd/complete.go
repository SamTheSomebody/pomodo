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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		db := helpers.GetDBQueries()
		id := helpers.GetTask(cmd.Context(), args[0]).ID
		task, err := db.CompleteTask(cmd.Context(), id)
		if err != nil {
			log.Fatalf("Error completing task: %v", err)
		}
		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

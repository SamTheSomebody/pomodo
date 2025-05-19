/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a timer in the background",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// TODO background flag
	// TODO time flag - goes directly to timer
	// TODO no time flag - goes to start page
}

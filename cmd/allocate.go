/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// allocateCmd represents the allocate command
var allocateCmd = &cobra.Command{
	Use:   "allocate",
	Short: "Allocate the amount of time you'd like to work today.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("allocate called")
		// TODO Add allocation
	},
}

func init() {
	rootCmd.AddCommand(allocateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allocateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allocateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

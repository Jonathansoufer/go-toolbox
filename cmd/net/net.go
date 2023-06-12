/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package net

import (

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net is a palette to manage network",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	NetCmd.AddCommand(pingCmd)
}

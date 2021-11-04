package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var rootCmd=&cobra.Command{
	Use: "images",
	Short: "Print images information",
	Long: "Print all images information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd)
	},
}

func Execute()  {
	rootCmd.Execute()
}
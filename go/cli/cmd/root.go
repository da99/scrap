/*
Copyright © 2025 da99/diego <null@null>

*/
package cmd

import (
	// "os"
	// "fmt"
	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Compile .go.html files",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() error {
	return rootCmd.Execute()
	// err := rootCmd.Execute()
	// if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	// }
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



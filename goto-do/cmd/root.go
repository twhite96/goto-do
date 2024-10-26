/*
Copyright © 2024 tiff github@tifflabs.org

*/
package cmd

import (
	"os"
  "encoding/csv"
  "strconv"
  "text/tabwriter"
  "github.com/mergestat/timediff"
	"github.com/spf13/cobra"
  "errors"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   `goto-do`,
	Short: `CLI todo CRUD app`,
  Long:  `CLI todo app where you can create, update, and delete tasks all 
  from the command line. Written in Go`
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goto-do.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



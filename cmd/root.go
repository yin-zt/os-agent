package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cb",
	Short: "cb agent",
	Long:  `A tool to collect and report os info`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

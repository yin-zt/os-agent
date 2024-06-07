package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yin-zt/os-agent/core"
)

var cloudbootAgent = &cobra.Command{
	Use:   "cloudboot-agent",
	Short: "run cloudboot-agent",
	Long:  `run cloudboot-agent`,
	Run: func(cmd *cobra.Command, args []string) {
		core.RunAgent()
	},
}

func init() {
	rootCmd.AddCommand(cloudbootAgent)
}

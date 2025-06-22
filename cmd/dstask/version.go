package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show dstask version information",
	Long:  `Display version, build date, and git commit information.`,
	Run: func(cmd *cobra.Command, args []string) {
		dstask.CommandVersion()
	},
}

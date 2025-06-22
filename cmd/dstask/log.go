package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [task summary]",
	Short: "Log a task (already resolved)",
	Long:  `Create a task that is already marked as resolved. Useful for recording completed work.`,
	Example: `  dstask log fixed critical bug +work
  dstask log completed project documentation`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandLog(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

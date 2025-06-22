package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next [filter]",
	Short: "Show most important tasks (priority, creation date -- truncated and default)",
	Long: `Show the most important tasks sorted by priority then creation date.
This is the default command when no command is specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandNext(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

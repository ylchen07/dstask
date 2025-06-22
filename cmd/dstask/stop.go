package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [id...]",
	Short: "Change task status to pending",
	Long:  `Stop working on one or more tasks, changing their status back to pending.`,
	Example: `  dstask stop 1
  dstask stop 1 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandStop(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

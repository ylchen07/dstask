package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [id...]",
	Short: "Change task status to active",
	Long:  `Mark one or more tasks as actively being worked on.`,
	Example: `  dstask start 1
  dstask start 1 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandStart(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

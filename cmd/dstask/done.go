package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done [id...]",
	Aliases: []string{"resolve"},
	Short:   "Resolve a task",
	Long:    `Mark one or more tasks as completed/resolved.`,
	Example: `  dstask done 1
  dstask resolve 1 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandDone(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

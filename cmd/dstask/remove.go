package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [id...]",
	Aliases: []string{"rm"},
	Short:   "Remove a task (use to remove tasks added by mistake)",
	Long:    `Remove one or more tasks by ID. This permanently deletes the task.`,
	Example: `  dstask remove 1
  dstask rm 1 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandRemove(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

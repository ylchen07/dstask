package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "context [filter]",
	Short: "Set global context for task list and new tasks",
	Long: `Set or view the global context that filters tasks and applies to new tasks.
Use "none" to clear the context.`,
	Example: `  dstask context +work
  dstask context project:dstask
  dstask context none
  dstask context  # Show current context`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandContext(conf, state, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

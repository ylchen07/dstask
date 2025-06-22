package main

import (
	"os"

	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo last action with git revert",
	Long:  `Undo the last action by reverting the most recent git commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandUndo(conf, os.Args, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

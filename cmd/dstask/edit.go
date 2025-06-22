package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit [id]",
	Short:   "Edit task with text editor",
	Long:    `Open the specified task in your default text editor ($EDITOR).`,
	Example: `  dstask edit 1`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandEdit(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

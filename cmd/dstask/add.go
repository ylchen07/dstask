package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task summary/filter]",
	Short: "Add a task",
	Long: `Add a new task with tags, projects, and priority specified.
Tags are specified with + (eg: +work), projects with project: prefix (eg: project:dstask),
and priorities from P3 (low), P2 (default) to P1 (high) and P0 (critical).`,
	Example: `  dstask add fix server +work
  dstask add project:dstask improve CLI P1
  dstask add template:24  # Create from template`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandAdd(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

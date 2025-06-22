package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:   "modify [id...] [attributes]",
	Short: "Set attributes for a task",
	Long: `Modify task attributes like tags, project, and priority.
Can modify multiple tasks at once.`,
	Example: `  dstask modify 1 +urgent
  dstask modify 1 2 project:newproject
  dstask modify 1 P1`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandModify(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

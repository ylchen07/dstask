package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template [task summary/filter]",
	Short: "Add a task template",
	Long:  `Create a task template that can be reused with the add command.`,
	Example: `  dstask template weekly review +work
  dstask add template:1  # Use template with ID 1`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandTemplate(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

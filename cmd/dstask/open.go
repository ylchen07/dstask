package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [id...]",
	Short: "Open all URLs found in summary/annotations",
	Long:  `Open all URLs found in the specified tasks in the default browser.`,
	Example: `  dstask open 1
  dstask open 1 2 3`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandOpen(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

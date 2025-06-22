package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Pull then push to git repository, automatic merge commit",
	Long: `Synchronize with the remote git repository by pulling changes,
merging automatically, and pushing local changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := dstask.CommandSync(conf.Repo); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

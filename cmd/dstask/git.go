package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git [git-command]",
	Short: "Pass a command to git in the repository",
	Long:  `Execute git commands in the dstask repository. Used for push/pull and other git operations.`,
	Example: `  dstask git status
  dstask git push
  dstask git log`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		dstask.MustRunGitCmd(conf.Repo, args...)
	},
}

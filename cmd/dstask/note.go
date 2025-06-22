package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:     "note [id] [note text]",
	Aliases: []string{"notes"},
	Short:   "Append to or edit note for a task",
	Long: `Add or edit markdown notes for a task. If no note text is provided,
opens the note in your default editor.`,
	Example: `  dstask note 1 "Important detail about this task"
  dstask note 1  # Opens editor`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandNote(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

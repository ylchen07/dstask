package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var (
	conf  dstask.Config
	state dstask.State
	ctx   dstask.Query
)

var rootCmd = &cobra.Command{
	Use:   "dstask",
	Short: "Single binary terminal-based TODO manager with git-based sync",
	Long: `dstask is a personal task tracker designed to help you focus.
It uses git to synchronise tasks and supports contexts, priorities, and markdown notes.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Default command is "next"
		if err := runNextCommand(args); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

func Execute() {
	// Custom pre-processing for dstask's flexible argument system
	// If the first argument is not a known command, treat it as arguments to the default "next" command
	args := os.Args[1:]
	if len(args) > 0 {
		cmdName := args[0]

		// Skip custom processing for help flags
		if cmdName == "--help" || cmdName == "-h" {
			if err := rootCmd.Execute(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			return
		}

		// Check if this is a known command
		isKnownCmd := false
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == cmdName || contains(cmd.Aliases, cmdName) {
				isKnownCmd = true
				break
			}
		}

		// If it's not a known command, treat everything as arguments to the default command
		if !isKnownCmd {
			if err := runNextCommand(args); err != nil {
				dstask.ExitFail(err.Error())
			}
			return
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add all commands
	rootCmd.AddCommand(nextCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(contextCmd)
	rootCmd.AddCommand(modifyCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(noteCmd)
	rootCmd.AddCommand(undoCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(gitCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(showActiveCmd)
	rootCmd.AddCommand(showPausedCmd)
	rootCmd.AddCommand(showOpenCmd)
	rootCmd.AddCommand(showProjectsCmd)
	rootCmd.AddCommand(showTagsCmd)
	rootCmd.AddCommand(showTemplatesCmd)
	rootCmd.AddCommand(showResolvedCmd)
	rootCmd.AddCommand(showUnorganisedCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(bashCompletionCmd)
	rootCmd.AddCommand(zshCompletionCmd)
	rootCmd.AddCommand(fishCompletionCmd)
	rootCmd.AddCommand(completionsCmd)
}

func initConfig() {
	conf = dstask.NewConfig()
	dstask.EnsureRepoExists(conf.Repo)

	// Load state for getting and setting ctx
	state = dstask.LoadState(conf.StateFile)
	ctx = state.Context

	// Check if we have a context override.
	if conf.CtxFromEnvVar != "" {
		splitted := strings.Fields(conf.CtxFromEnvVar)
		ctx = dstask.ParseQuery(splitted...)
	}
}

// Helper function to parse query from args
func parseQueryFromArgs(args []string) dstask.Query {
	query := dstask.ParseQuery(args...)

	// Check if we ignore context with the "--" token
	if query.IgnoreContext {
		return query
	}

	return query.Merge(ctx)
}

func runNextCommand(args []string) error {
	query := parseQueryFromArgs(args)
	return dstask.CommandNext(conf, ctx, query)
}

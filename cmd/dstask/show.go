package main

import (
	"github.com/naggie/dstask"
	"github.com/spf13/cobra"
)

var showActiveCmd = &cobra.Command{
	Use:   "show-active [filter]",
	Short: "Show tasks that have been started",
	Long:  `Display all tasks that are currently in active status.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowActive(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showPausedCmd = &cobra.Command{
	Use:   "show-paused [filter]",
	Short: "Show tasks that have been started then stopped",
	Long:  `Display all tasks that were started but then paused.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowPaused(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showOpenCmd = &cobra.Command{
	Use:   "show-open [filter]",
	Short: "Show all non-resolved tasks (without truncation)",
	Long:  `Display all open tasks without truncating the output.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowOpen(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showProjectsCmd = &cobra.Command{
	Use:   "show-projects [filter]",
	Short: "List projects with completion status",
	Long:  `Display all projects with their completion statistics.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowProjects(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showTagsCmd = &cobra.Command{
	Use:   "show-tags [filter]",
	Short: "List tags in use",
	Long:  `Display all tags currently in use across tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowTags(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showTemplatesCmd = &cobra.Command{
	Use:   "show-templates [filter]",
	Short: "Show task templates",
	Long:  `Display all available task templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowTemplates(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showResolvedCmd = &cobra.Command{
	Use:   "show-resolved [filter]",
	Short: "Show resolved tasks",
	Long:  `Display tasks that have been completed/resolved.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowResolved(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

var showUnorganisedCmd = &cobra.Command{
	Use:   "show-unorganised [filter]",
	Short: "Show untagged tasks with no projects (global context)",
	Long:  `Display tasks that have no tags and no project assigned.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := parseQueryFromArgs(args)
		if err := dstask.CommandShowUnorganised(conf, ctx, query); err != nil {
			dstask.ExitFail(err.Error())
		}
	},
}

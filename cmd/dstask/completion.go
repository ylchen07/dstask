package main

import (
	"fmt"
	"os"

	"github.com/naggie/dstask/completions"
	"github.com/spf13/cobra"
)

var bashCompletionCmd = &cobra.Command{
	Use:   "bash-completion",
	Short: "Print bash completion script to stdout",
	Long:  `Generate bash completion script. Add to your .bashrc: source <(dstask bash-completion)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(completions.Bash)
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh-completion",
	Short: "Print zsh completion script to stdout",
	Long:  `Generate zsh completion script. Add to your .zshrc: source <(dstask zsh-completion)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(completions.Zsh)
	},
}

var fishCompletionCmd = &cobra.Command{
	Use:   "fish-completion",
	Short: "Print fish completion script to stdout",
	Long:  `Generate fish completion script.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(completions.Fish)
	},
}

// Legacy completions command for backward compatibility
var completionsCmd = &cobra.Command{
	Use:    "_completions",
	Short:  "Internal completions command",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		completions.Completions(conf, os.Args, ctx)
	},
}

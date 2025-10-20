package main

import (
	"path/filepath"

	cli "github.com/isaqueveras/jangada/internal"
	newapp "github.com/isaqueveras/jangada/internal/new-app"
	"github.com/isaqueveras/jangada/internal/sail"
	"github.com/spf13/cobra"
)

func main() {
	dirBase, _ := filepath.Abs("")

	cli.Init(dirBase)
	cli.SetFullDirectoryPath()

	cmd := &cobra.Command{
		Use:               "jangada",
		Short:             "Jangada is a CLI tool for project scaffolding and code generation.",
		Example:           "jangada new my-app --module github.com/username/my-app --database postgres",
		ValidArgsFunction: cobra.FixedCompletions([]cobra.Completion{"new", "sail"}, cobra.ShellCompDirective(0)),
		SuggestFor:        []string{"new", "sail"},
		Version:           "v0.1.0-beta",
	}

	cmd.AddCommand(
		// Command `jangada new my-app`
		newapp.Command(),

		// Command `jangada sail ...`
		sail.Command(),
	)

	cmd.Execute()
}

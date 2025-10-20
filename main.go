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

	root := &cobra.Command{
		Use:               "jangada",
		Short:             "Jangada is a CLI tool for project scaffolding and code generation.",
		Example:           "jangada new my-app --module github.com/username/my-app --database postgres",
		ValidArgsFunction: cobra.FixedCompletions([]cobra.Completion{"new", "sail"}, cobra.ShellCompDirective(0)),
		SuggestFor:        []string{"new", "sail"},
		Version:           "v0.1.0-beta",
	}

	commandNew := &cobra.Command{
		Use:     "new [name]",
		Short:   "Create a new app",
		Args:    cobra.ExactArgs(1),
		Example: "jangada new myapp",
		Run:     newapp.Execute,
		Aliases: []string{"n"},
	}

	commandSail := &cobra.Command{
		Use:        "sail",
		Short:      "Create layer for bounded context",
		Example:    "jangada sail",
		ArgAliases: []string{"transport"},
		Aliases:    []string{"s"},
	}

	commandTransport := &cobra.Command{
		Use:     "transport",
		Short:   "Create transport layer",
		Args:    cobra.RangeArgs(1, 2),
		Example: "jangada sail transport catalog/company",
		Run:     sail.Execute,
		ValidArgs: []cobra.Completion{
			"web",
			// "rest",
			// "grpc",
			// "graphql",
			// "webhook",
			// "all",
		},
	}

	commandSail.AddCommand(commandTransport)
	root.AddCommand(commandNew, commandSail)

	root.Execute()
}

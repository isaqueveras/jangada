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
		Example:           "jangada new myapp --mod=github.com/username/myapp --db=postgres",
		ValidArgsFunction: cobra.FixedCompletions([]cobra.Completion{"new", "sail"}, cobra.ShellCompDirective(0)),
		SuggestFor:        []string{"new", "sail"},
		Version:           "v0.1.0-beta",
	}

	commandNew := &cobra.Command{
		Use:     "new [name]",
		Short:   "Create a new app",
		Args:    cobra.ExactArgs(1),
		Example: "jangada new myapp --mod=github.com/username/myapp --db=postgres",
		Run:     newapp.Execute,
	}

	commandNew.Flags().String("host", ":8080", "host")
	commandNew.Flags().String("mod", cli.GetConfig().AppName, "mod")
	commandNew.Flags().String("db", "postgres", "db")

	commandNew.PreRun = func(cmd *cobra.Command, args []string) {
		cli.SetAppName(args[0])
		cli.SetDefaultHost(cmd.Flag("host").Value.String())
		cli.SetModuleName(cmd.Flag("mod").Value.String())
		cli.SetDatabase(cmd.Flag("db").Value.String())
	}

	commandSail := &cobra.Command{
		Use:        "sail",
		Short:      "Create layer for bounded context",
		Example:    "jangada sail",
		ArgAliases: []string{"transport"},
		Aliases:    []string{"s"},
	}

	commandTransport := &cobra.Command{
		Use:       "transport",
		Short:     "Create transport layer",
		Args:      cobra.RangeArgs(1, 2),
		Example:   "jangada sail transport catalog/company --layer={web,rest}",
		Run:       sail.Execute,
		ValidArgs: []cobra.Completion{"web", "rest"},
	}

	commandTransport.Flags().String("layer", "web", "choose transport layer")
	commandTransport.Flags().String("name", "", "create a method/router in controller")

	commandSail.AddCommand(commandTransport)
	root.AddCommand(commandNew, commandSail)

	root.Execute()
}

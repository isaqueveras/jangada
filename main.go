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

	commandNew.Flags().String("host", ":8080", "--host=localhost:8080")
	commandNew.Flags().String("mod", "", "--mod=github.com/username/myapp")
	commandNew.Flags().String("db", "postgres", "--db=postgres")

	commandSail := &cobra.Command{
		Use:     "sail",
		Short:   "Create layer for bounded context",
		Example: "jangada sail",
		Aliases: []string{"s"},
		ArgAliases: []string{
			"transport",
			"application",
			"domain",
		},
	}

	{
		commandTransport := &cobra.Command{
			Use:       "transport",
			Short:     "Create transport layer",
			Args:      cobra.RangeArgs(1, 2),
			Example:   "jangada sail transport crm/customer --layer={web,rest}",
			Run:       sail.Transport,
			ValidArgs: []cobra.Completion{"web", "rest"},
		}

		commandTransport.Flags().String("layer", "web", "choose transport layer")
		commandTransport.Flags().String("name", "", "create a method/router in controller")

		commandSail.AddCommand(commandTransport)
	}

	{
		commandApplication := &cobra.Command{
			Use:     "application",
			Short:   "Create application layer",
			Aliases: []string{"app"},
			Args:    cobra.RangeArgs(1, 2),
			Example: "jangada sail application crm/customer --service=CreateCompany",
			Run:     sail.Application,
		}

		commandApplication.Flags().String("service", "", "create a new service in application")
		commandSail.AddCommand(commandApplication)
	}

	commandSail.AddCommand(&cobra.Command{
		Use:     "domain",
		Short:   "Create domain layer",
		Args:    cobra.RangeArgs(1, 2),
		Example: "jangada sail domain crm/customer",
		Run:     sail.Domain,
	})

	root.AddCommand(commandNew, commandSail)
	root.Execute()
}

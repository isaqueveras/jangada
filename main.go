package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	cli "github.com/isaqueveras/jangada/internal"
	newapp "github.com/isaqueveras/jangada/internal/new-app"
	"github.com/isaqueveras/jangada/internal/sail"
	"github.com/spf13/cobra"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go gracefullyStop(cancel)

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
		Example: "jangada new myapp --mod=github.com/username/myapp --db=postgres --host=localhost:8081",
		Run:     newapp.Execute,
	}

	commandNew.Flags().String("host", "localhost", "--host=localhost")
	commandNew.Flags().String("port", "8080", "--port=8080")
	commandNew.Flags().String("mod", "", "--mod=github.com/username/myapp")
	commandNew.Flags().String("db", "postgres", "--db=postgres")

	commandSail := &cobra.Command{
		Use:        "sail",
		Short:      "Create layer for bounded context",
		Example:    "jangada sail",
		Aliases:    []string{"s"},
		ArgAliases: []string{"transport", "application", "domain", "infrastructure"},
	}

	{
		commandTransport := &cobra.Command{
			Use:       "transport",
			Short:     "Create transport layer",
			Args:      cobra.RangeArgs(1, 2),
			Aliases:   []string{"t"},
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
			Aliases: []string{"app", "a"},
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
		Aliases: []string{"d"},
		Args:    cobra.RangeArgs(1, 2),
		Example: "jangada sail domain crm/customer",
		Run:     sail.Domain,
	})

	commandSail.AddCommand(&cobra.Command{
		Use:     "infrastructure",
		Short:   "Create infrastructure layer",
		Aliases: []string{"infra", "i"},
		Args:    cobra.RangeArgs(1, 2),
		Example: "jangada sail infrastructure crm/customer",
		Run:     sail.Infrastructure,
	})

	root.AddCommand(commandNew, commandSail)
	root.Execute()
}

func gracefullyStop(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 2)
	defer close(stop)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(stop)

	<-stop
	cancel()

	fmt.Fprintln(os.Stdout, "\ninterrupt received, wait for exit or ^C to terminate")
	<-stop

	os.Exit(1)
}

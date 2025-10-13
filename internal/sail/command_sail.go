// Package sail provides commands to create layers for a bounded context.
package sail

import "github.com/spf13/cobra"

// Sail defines the Sail command structure.
type Sail struct {
	CommandTransport SailTransport
}

// Command represents the 'sail' command.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sail",
		Short: "Create layer for bounded context",
	}

	cmd.AddCommand(
		transportCommand(),
	)

	return cmd
}

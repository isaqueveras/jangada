// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"github.com/spf13/cobra"
)

// Sail defines the Sail command structure.
type Sail struct {
	PathDir          string
	CommandTransport SailTransport
}

// NewCommand creates a new instance of the Sail command.
func NewCommand(pathDir string) *Sail {
	return &Sail{
		PathDir: pathDir,
		CommandTransport: SailTransport{
			pathDir: pathDir,
		},
	}
}

// Execute represents the 'sail' command.
func (s *Sail) Execute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sail",
		Short: "Create layer for bounded context",
	}

	cmd.AddCommand(&cobra.Command{
		Use:     "transport",
		Short:   "Create transport layer",
		Args:    cobra.RangeArgs(1, 2),
		Example: exampleCreateTransportText,
		Run:     s.CommandTransport.Execute,
	})

	return cmd
}

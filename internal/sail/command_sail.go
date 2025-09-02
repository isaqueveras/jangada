// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"github.com/spf13/cobra"
)

// Sail defines the Sail command structure.
type Sail struct {
	PathDir          string
	CommandInterface SailInterface
}

// NewCommand creates a new instance of the Sail command.
func NewCommand() *Sail {
	return &Sail{}
}

// Execute represents the 'sail' command.
func (s *Sail) Execute(pathDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sail",
		Short: "Create layer for bounded context",
	}

	cmd.AddCommand(&cobra.Command{
		Use:     "interface",
		Short:   "Create interface for bounded context",
		Args:    cobra.RangeArgs(1, 2),
		Example: exampleCreateInterfaceText,
		Run: func(_ *cobra.Command, args []string) {
			s.CommandInterface.Execute(pathDir, args)
		},
	})

	return cmd
}

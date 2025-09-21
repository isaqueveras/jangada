package template

// ApplicationOrchestrator is a template for an application orchestrator
const ApplicationOrchestrator = `// Package orchestrator defines an orchestrator for {{ .Entity }}
package orchestrator

import "context"

// {{ .Entity }}Orchestrator exposes the orchestrator methods for {{ ToLower .Entity }}
type {{ .Entity }}Orchestrator interface {
	// GetByID returns a {{ ToLower .Entity }}
	GetByID(ctx context.Context, id string) (any, error)
	
	// List returns a list of {{ ToLower .Entity }}
	List(ctx context.Context, filters map[string]string) ([]any, error)
	
	// Create creates a new {{ ToLower .Entity }}
	Create(ctx context.Context, params *command.{{ .Entity }}CreateParams) (any, error)
	
	// Update updates a {{ ToLower .Entity }}
	Update(ctx context.Context, params *command.{{ .Entity }}UpdateParams) (any, error)

	// Delete deletes a {{ .Entity }}
	Delete(ctx context.Context, id string) error
}
`

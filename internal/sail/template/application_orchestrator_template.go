package template

// ApplicationOrchestrator is a template for an application orchestrator
const ApplicationOrchestrator = `// Package orchestrator defines an orchestrator for {{ ToLower .Entity }}
package orchestrator

import (
	"context"

	"{{ .Module }}/internal/application/{{ .Folder }}/command"
	"{{ .Module }}/internal/application/{{ .Folder }}/query"
)

// {{ .Entity }}Orchestrator exposes the orchestrator methods for {{ ToLower .Entity }}
type {{ .Entity }}Orchestrator interface {
	// {{ .Entity }}OrchestratorReader is a reader for {{ ToLower .Entity }}
	{{ .Entity }}OrchestratorReader

	// {{ .Entity }}OrchestratorWriter is a writer for {{ ToLower .Entity }}
	{{ .Entity }}OrchestratorWriter
}

// {{ .Entity }}OrchestratorReader exposes the reader methods for {{ ToLower .Entity }}
type {{ .Entity }}OrchestratorReader interface {
	// Get returns a {{ ToLower .Entity }} by id
	Get(ctx context.Context, id *string) (*query.{{ .Entity }}Item, error)

	// List returns a list of {{ ToLower .Entity }}s
	List(ctx context.Context, filters map[string]any) (*query.List{{ .Entity }}Item, error)
}

// {{ .Entity }}OrchestratorWriter exposes the writer methods for {{ ToLower .Entity }}
type {{ .Entity }}OrchestratorWriter interface {
	// Create creates a new {{ ToLower .Entity }}
	Create(ctx context.Context, params *command.Create{{ .Entity }}) (*string, error)

	// Update updates a {{ ToLower .Entity }}
	Update(ctx context.Context, params *command.Update{{ .Entity }}) error

	// Delete deletes a {{ ToLower .Entity }}
	Delete(ctx context.Context, id *string) error
}
`

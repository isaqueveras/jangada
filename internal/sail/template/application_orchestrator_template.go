package template

// ApplicationOrchestrator is a template for an application orchestrator
const ApplicationOrchestrator = `// Package {{ ToLower .Entity }} defines an orchestrator for {{ ToLower .Entity }}
package {{ ToLower .Entity }}

import (
	"context"

	"{{ .Module }}/pkg/database"
)

// Orchestrator exposes the orchestrator methods for {{ ToLower .Entity }}
type Orchestrator interface {
	// GetByID returns a {{ ToLower .Entity }} name by id
	GetByID(ctx context.Context, id string) (string, error)
}

// orchestrator is a orchestrator for {{ ToLower .Entity }} that implements the Orchestrator interface
type orchestrator struct {
	conn database.ConnectionPool
	// srv  domain.ServiceBuilder
}

// NewOrchestrator creates a new instance of orchestrator with the given repository
func NewOrchestrator(conn database.ConnectionPool /*srv domain.ServiceBuilder*/) Orchestrator {
	return &orchestrator{
		conn: conn,
		// srv:  srv,
	}
}

// GetByID returns a {{ ToLower .Entity }} name by id
func (o *orchestrator) GetByID(ctx context.Context, id string) (string, error) {
	return "", nil
}
`

const ApplicationService = `// Package {{ ToLower .Entity }} defines an service for {{ ToLower .Entity }}
package {{ ToLower .Entity }}

import "context"

// service is a service for company
type service struct {
	// repository {{ ToLower .Entity }}Domain.Repository
}

// NewService creates a new instance of service with the given repository
func NewService( /* repository {{ ToLower .Entity }}Domain.Repository */ ) *service {
	return &service{
		// repository: repository
	}
}

// GetByID finds a {{ ToLower .Entity }} by its ID
func (s *service) GetByID(ctx context.Context, id string) (int64, error) {
	// return s.repository.GetByID(ctx, id)
	return 0, nil
}
`

const ApplicationBuilder = `// Package application defines a builder for application
package application

type builder struct {
	// infra domain.RepositoryBuilder
}

// NewBuilder creates a new instance of ServiceBuilder
func NewBuilder() *builder {
	return &builder{
		// infra: infrastructure.NewBuilder(),
	}
}
`

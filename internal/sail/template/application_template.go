package template

// ApplicationOrchestrator is a template for an application orchestrator
const ApplicationOrchestrator = `// Package {{ ToLower .Entity }} defines an orchestrator for {{ ToLower .Entity }}
package {{ ToLower .Entity }}

import (
	"{{ .Module }}/internal/domain"
	"{{ .Module }}/pkg/database"
)

// Orchestrator exposes the orchestrator methods for {{ ToLower .Entity }}
type Orchestrator interface{}

// orchestrator is a orchestrator for {{ ToLower .Entity }} that implements the Orchestrator interface
type orchestrator struct {
	conn database.ConnectionPool
	srv  domain.ServiceBuilder
}

// NewOrchestrator creates a new instance of orchestrator with the given repository
func NewOrchestrator(conn database.ConnectionPool, srv domain.ServiceBuilder) Orchestrator {
	return &orchestrator{conn: conn, srv: srv}
}
`

const ApplicationService = `// Package {{ ToLower .Entity }} defines an service for {{ ToLower .Entity }}
package {{ ToLower .Entity }}

type service struct {
	// repository {{ ToLower .Entity }}.Repository
}

// NewService creates a new instance of service with the given repository
func NewService( /* repository {{ ToLower .Entity }}.Repository */ ) *service {
	return &service{
		// repository: repository
	}
}
`

const ApplicationBuilder = `// Package application provides application management
package application

import (
	"{{ .ModuleName }}/internal/domain"
	"{{ .ModuleName }}/internal/infrastructure"
)

type builder struct {
	repository domain.RepositoryBuilder
}

// New creates a new instance of builder with the given repository
func New() *builder {
	return &builder{repository: infrastructure.New()}
}
`

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
	srv  domain.Services
}

// NewOrchestrator creates a new instance of orchestrator with the given repository
func NewOrchestrator(conn database.ConnectionPool, srv domain.Services) Orchestrator {
	return &orchestrator{conn: conn, srv: srv}
}
`

// DomainService is a template for a domain service
const DomainService = `// Package {{ ToLower .Entity }} contains the business logic for the {{ ToLower .Entity }} domain
package {{ ToLower .Entity }}

// Service is the interface for the {{ ToLower .Entity }} service
type Service interface{}

type service struct {
	repository Repository
}

// NewService creates a new instance of service with the given repository
func NewService(repository Repository) *service {
	return &service{repository: repository}
}
`

// DomainRepository is a template for a domain repository
const DomainRepository = `// Package {{ ToLower .Entity }} contains the repository for the {{ ToLower .Entity }} domain
package {{ ToLower .Entity }}

// Repository defines the interface for the {{ ToLower .Entity }} repository
type Repository interface{}
`

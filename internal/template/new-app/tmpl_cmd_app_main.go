// Package newapp contains templates for new app
package newapp

const tmplAppMain = `// Package main contains the main function of the app
package main

import (
	"{{ .ModuleName }}/core"
	"{{ .ModuleName }}/pkg/database"
)

func main() {
	server := core.New()

	conn := database.NewConnectionPool(server.Config().GetDatabases()...)
	defer conn.CloseConnections()

	// Uncomment the code below to enable the transport layer
	// transport.Handler(server, conn)

	server.Init()
}
`

const tmplApplicationBuilder = `// Package application provides application management
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
	return &builder{
		repository: infrastructure.New(),
	}
}
`

const tmplDomainBuilder = `// Package domain defines the interfaces for building repositories and services.
package domain

// RepositoryBuilder is a builder for repository
type RepositoryBuilder interface{}

// ServiceBuilder is a builder for service
type ServiceBuilder interface{}
`

const tmplInfrastructureBuilder = `// Package infrastructure provides infrastructure services and builders
package infrastructure

type builder struct{}

// New creates a new instance of builder for infrastructure
func New() *builder {
	return &builder{}
}
`

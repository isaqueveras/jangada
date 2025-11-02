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

const tmplApplicationService = `// Package application provides application management
package application

import (
	"{{ .ModuleName }}/internal/domain"
	"{{ .ModuleName }}/internal/infrastructure"
)

type builder struct {
	repository domain.Repositories
}

// NewService creates a new instance of builder with the given repository
func NewService() *builder {
	return &builder{repository: infrastructure.NewRepository()}
}
`

const tmplDomainInterface = `// Package domain defines the interfaces for building repositories and services.
package domain

// Repositories is a builder for repository
type Repositories interface{}

// Services is a builder for service
type Services interface{}
`

const tmplInfrastructureRepository = `// Package infrastructure provides infrastructure services and builders
package infrastructure

type builder struct{}

// New creates a new instance of builder for infrastructure
func NewRepository() *builder {
	return &builder{}
}
`

// Package newapp contains templates for new app
package newapp

const tmplAppMain = `// Package main contains the main function of the app
package main

import (
	"{{ .ModuleName }}/core"
	"{{ .ModuleName }}/pkg/database"
)

func main() {
	app := core.New()

	conn := database.NewConnectionPool(app.Config().GetDatabases()...)
	defer conn.CloseConnections()

	if err := app.Metrics(); err != nil {
		app.Log().Error("failed to register metrics: " + err.Error())
	}

	// Uncomment the code below to enable the transport layer
	// transport.Handler(app, conn)

	app.Run()
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

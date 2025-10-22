package newapp

const tmplCoreCore string = `// Package core defines the core framework
package core

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"{{ .ModuleName }}/config"
	"{{ .ModuleName }}/web/layouts"
)

const (
	_defaultHost = "{{ .DefaultHost }}"
)

// Core defines the core framework
type Core struct {
	router  *gin.Engine
	log     *slog.Logger
	address string
}

// New creates a new core framework
func New() *Core {
	config.NewConfig()

	server := &Core{
		address: _defaultHost,
		router:  gin.Default(),
		log:     slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}

	server.router.Static("public", "public")

	server.router.NoRoute(func(ctx *gin.Context) {
		layouts.NotFound().Render(ctx.Request.Context(), ctx.Writer)
	})

	server.router.GET("/", func(ctx *gin.Context) {
		layouts.Welcome().Render(ctx.Request.Context(), ctx.Writer)
	})

	return server
}

func (c *Core) Init() error         { return c.router.Run(c.address) }
func (c *Core) Router() *gin.Engine { return c.router }
func (c *Core) Log() *slog.Logger   { return c.log }
`

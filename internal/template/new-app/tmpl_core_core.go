// Package newapp contains templates for new app
package newapp

const tmplCoreCore string = `// Package core defines the core framework
package core

import (
	"log"
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
	cfg			*config.Config
	router  *gin.Engine
	log     *slog.Logger
	address string
}

// New creates a new core framework
func New() *Core {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := &Core{
		address: _defaultHost,
		router:  gin.Default(),
		log:     slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
		cfg:     cfg,
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

// Config returns the config
func (c *Core) Config() *config.Config { return c.cfg }
`

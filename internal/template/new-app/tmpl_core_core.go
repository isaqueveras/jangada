// Package newapp contains templates for new app
package newapp

const tmplCoreCore string = `// Package core defines the core framework
package core

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"{{ .ModuleName }}/config"
	"{{ .ModuleName }}/pkg/metric"
	"{{ .ModuleName }}/web/layouts"
)

// Core defines the core framework
type Core struct {
	cfg    *config.Config
	router *gin.Engine
	log    *slog.Logger
}

// New creates a new core framework
func New() *Core {
	logg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	if err := godotenv.Load(); err != nil {
		logg.Info("No .env file found, continuing with environment variables...")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		logg.Error("Failed to load config: " + err.Error())
		return nil
	}

	if err = cfg.LoadDatabase("{{ ToLower .AppName }}"); err != nil {
		logg.Error("Failed to load databases: " + err.Error())
		return nil
	}

	server := &Core{
		router: gin.Default(),
		log:    logg,
		cfg:    cfg,
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

// Run starts the server
func (c *Core) Run() error {
	return c.router.Run(c.cfg.GetApplication().Address)
}

// Router returns the router
func (c *Core) Router() *gin.Engine {
	return c.router
}

// Log returns the logger
func (c *Core) Log() *slog.Logger {
	return c.log
}

// Config returns the config
func (c *Core) Config() *config.Config {
	return c.cfg
}

// Metrics returns the metrics middleware and router
func (c *Core) Metrics() error {
	nps, err := metric.NewPrometheusMetrics(c.Config())
	if err != nil {
		return err
	}

	c.Router().Use(metric.Middleware(nps))
	c.Router().GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})

	return nil
}
`

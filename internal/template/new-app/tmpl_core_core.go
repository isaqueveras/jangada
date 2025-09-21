package newapp

const tmplCoreCore string = `// Package core defines the core framework
package core

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"{{ .ModuleName }}/config"
	"{{ .ModuleName }}/web/layouts"
)

const (
	_defaultAddr = "{{ .DefaultAddrPort }}"
)

type Core struct {
	router  *gin.Engine
	log     *slog.Logger
	db      *gorm.DB
	address string

	tls *TLS
}

type TLS struct {
	CertFile string
	KeyFile  string
}

func New() *Core {
	config.NewConfig()

	server := &Core{
		address: _defaultAddr,
		router:  gin.Default(),
		log:     slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
		tls:     &TLS{},
	}

	server.router.Static("public", "public")

	server.router.NoRoute(func(ctx *gin.Context) {
		layouts.NotFound().Render(ctx.Request.Context(), ctx.Writer)
	})

	server.router.GET("/", func(ctx *gin.Context) {
		layouts.Welcome().Render(ctx.Request.Context(), ctx.Writer)
	})

	if err := server.database(); err != nil {
		server.Log().Error("failed to connect to database", "error", err)
		return nil
	}

	return server
}

func (c *Core) Init() error {
	return c.router.Run(c.address)
}

func (c *Core) InitTLS() error {
	return c.router.RunTLS(c.address, c.tls.CertFile, c.tls.KeyFile)
}

func (c *Core) Router() *gin.Engine {
	return c.router
}

func (c *Core) Log() *slog.Logger {
	return c.log
}

func (c *Core) DB() *gorm.DB {
	return c.db
}

func (c *Core) database() (err error) {
	c.db, err = gorm.Open(sqlite.Open("db/jangada_development.db"))
	return err
}`

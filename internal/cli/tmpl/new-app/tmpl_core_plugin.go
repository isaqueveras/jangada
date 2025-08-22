package newapp

const tmplCorePlugin = `package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Plugin interface {
	Name() string
	Migrate(*gorm.DB) error
	Register(*gin.Engine)
}

func (c *Core) SetPlugin(plug Plugin) {
	c.log.Info(fmt.Sprintf("Registering plugin: %s", plug.Name()))
	if err := plug.Migrate(c.db); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	plug.Register(c.router)
	c.log.Info(fmt.Sprintf("Plugin %s registered", plug.Name()))
}`

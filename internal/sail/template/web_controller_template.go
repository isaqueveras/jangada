// Package template contains templates for a web controller
package template

// WebController is a template for a web controller
const WebController = `package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{ .Module }}/internal/application/{{ .Folder }}/command"
	"{{ .Module }}/internal/application/{{ .Folder }}/orchestrator"
	"jangada/core"
)

// {{ .Entity }}Controller ...
type {{ .Entity }}Controller struct {
	orchestrator orchestrator.{{ .Entity }}Orchestrator
}

// New{{ .Entity }}Controller create a new {{ .Entity }}Controller instance and register routes
func New{{ .Entity }}Controller(core *core.Core, app orchestrator.{{ .Entity }}Orchestrator) *{{ .Entity }}Controller {
	ctrl := &{{ .Entity }}Controller{
		orchestrator: app,
	}

	r := core.Router().Group("/v1/{{ .Folder }}")
	r.GET(":id", ctrl.GetByID)

	return ctrl
}

// GetByID define a method to get by id
func (c *{{ .Entity }}Controller) GetByID(ctx *gin.Context) {
	params := command.{{ .Entity }}Params{
		ID: ctx.Param("id"),
	}

	data, err := c.orchestrator.GetByID(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, data)
}
`

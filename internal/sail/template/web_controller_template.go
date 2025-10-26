// Package template contains templates for a {{ .Layer }} controller
package template

// ControllerTemplate is a template for a {{ .Layer }} controller
const ControllerTemplate = `// Package {{ ToLower .Entity }} defines a {{ .Layer }} controller for {{ .Entity }}Controller 
package {{ ToLower .Entity }}

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{ .Module }}/core"
	"{{ .Module }}/internal/application/{{ .Folder }}/orchestrator"
	"{{ .Module }}/internal/transport/{{ .Layer }}/{{ .Folder }}/request"
)

// Controller is a controller for {{ ToLower .Entity }}
type Controller struct {
	orchestrator orchestrator.Orchestrator
}

// NewController register routes for {{ ToLower .Entity }}
func NewController(core *core.Core, app orchestrator.Orchestrator) {
	ctrl := &Controller{
		orchestrator: app,
	}

	// Create a new group of resource. 
	r := core.Router().Group("/v1/{{ .Folder }}")

	// Display the specified resource.
	r.GET("{{ ToLower .Entity }}/:id", ctrl.GetByID)

	// Display a listing of the resource.
	r.GET("{{ ToLower .Entity }}", ctrl.GetAll)

	// Store a newly created resource.
	r.POST("{{ ToLower .Entity }}", ctrl.Create)

	// Update the specified resource.
	r.PUT("{{ ToLower .Entity }}/:id", ctrl.Update)

	// Remove the specified resource.
	r.DELETE("{{ ToLower .Entity }}/:id", ctrl.Delete)
}

// GetByID define a method to get a resource by id
func (c *{{ .Entity }}Controller) GetByID(ctx *gin.Context) {
	params := new(request.Get{{ .Entity }}Params)
	if err := ctx.ShouldBindUri(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := c.orchestrator.Get(ctx, params.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
		
	// mapper.To{{ .Entity }}Response(data)
	ctx.JSON(http.StatusOK, data)
}

// GetAll define a method to get all resources
func (c *{{ .Entity }}Controller) GetAll(ctx *gin.Context) {
	params := new(request.{{ .Entity }}Filters)
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := c.orchestrator.List(ctx, params.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
		
	ctx.JSON(http.StatusOK, data)
}

// Create define a method to create a resource
func (c *{{ .Entity }}Controller) Create(ctx *gin.Context) {
	params := new(request.{{ .Entity }}CreateParams)
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := c.orchestrator.Create(ctx, params.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

// Update define a method to update a resource
func (c *{{ .Entity }}Controller) Update(ctx *gin.Context) {
	params := new(request.{{ .Entity }}UpdateParams)
	if err := ctx.ShouldBindUri(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.orchestrator.Update(ctx, params.ToCommand()); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Delete define a method to delete a resource
func (c *{{ .Entity }}Controller) Delete(ctx *gin.Context) {
	params := new(request.{{ .Entity }}DeleteParams)
	if err := ctx.ShouldBindUri(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.orchestrator.Delete(ctx, params.ToCommand()); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
`

const ControllerMethod string = `// {{ .Method }} ...
func (c *Controller) {{ .Method }}(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
`

// ControllerTemplateClean is a template for a {{ .Layer }} controller
const ControllerTemplateClean = `// Package {{ ToLower .Entity }} defines a {{ .Layer }} controller 
package {{ ToLower .Entity }}

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{ .Module }}/core"
)

// Controller is a controller for {{ ToLower .Entity }}
type Controller struct {
	// orchestrator orchestrator.Orchestrator
}

// NewController register routes for {{ ToLower .Entity }}
func NewController(core *core.Core /* app orchestrator.Orchestrator */) {
	ctrl := &Controller{
		// orchestrator: app,
	}

	// Create a new group of resource. 
	r := core.Router().Group("/v1/api/{{ .Folder }}")

	// Create a new resource and its routes.
	r.GET("me", ctrl.Me)
}

func (c *Controller) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
`

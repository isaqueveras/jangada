// Package template contains templates for transport
package template

// ControllerTemplateRest is a template for a {{ .Layer }} controller
const ControllerTemplateRest = `// Package {{ ToLower .Entity }} defines a {{ .Layer }} controller 
package {{ ToLower .Entity }}

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{ .Module }}/core"
	"{{ .Module }}/internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
)

// Controller is a controller for {{ ToLower .Entity }}
type Controller struct {
	app {{ ToLower .Entity }}.Orchestrator
}

// NewController register routes for {{ ToLower .Entity }}
func NewController(core *core.Core, app {{ ToLower .Entity }}.Orchestrator) {
	ctrl := &Controller{
		app: app,
	}

	// Create a new group of resource.
	r := core.Router().Group("/v1/api/{{ .Folder }}")

	// Create a new resource and its routes.
	r.GET("ping", ctrl.Pong)
}

// Pong define a method to get a resource by id
func (c *Controller) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ping-pong"})
}
`

const ControllerTestTemplate = `// Package {{ ToLower .Entity }}_test defines a {{ .Layer }} controller 
package {{ ToLower .Entity }}_test
`

// HandlerController is a template for all controllers
const HandlerController = `// Package transport defines the transport layer handlers and routes.
package transport

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/transport/{{ .Layer }}"
	"{{ .Module }}/pkg/database"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core, conn database.ConnectionPool) {
	{{ .Layer }}.Handler(core, conn)
}
`

// HandlerLayerController is a template for a rest/web controller
const HandlerLayerController = `// Package {{ .Layer }} defines the transport layer handlers and routes.
package {{ ToLower .Layer }}

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/application"
	"{{ .Module }}/pkg/database"

	{{ ToLower .Entity }}App "{{ .Module }}/internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
	{{ ToLower .Layer }}{{ .Entity }}Transport "{{ .Module }}/internal/transport/{{ ToLower .Layer }}/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core, conn database.ConnectionPool) {
	{{ ToLower .Layer }}{{ .Entity }}Transport.NewController(core, {{ ToLower .Entity }}App.NewOrchestrator(conn, application.NewService()))
}
`

// ControllerMethod is a template for a rest controller
const ControllerMethod string = `// {{ .Method }} ...
func (c *Controller) {{ .Method }}(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
`

// ControllerTemplateWeb is a template for a web controller
const ControllerTemplateWeb = `// Package {{ ToLower .Entity }} defines a {{ .Layer }} controller 
package {{ ToLower .Entity }}

import (
	"github.com/gin-gonic/gin"

	"{{ .Module }}/core"
	"{{ .Module }}/core/helpers"
	"{{ .Module }}/internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
)

// Controller is a controller for {{ ToLower .Entity }}
type Controller struct {
	app {{ ToLower .Entity }}.Orchestrator
}

// NewController register routes for {{ ToLower .Entity }}
func NewController(core *core.Core, app {{ ToLower .Entity }}.Orchestrator) {
	ctrl := &Controller{
		app: app,
	}

	// Create a new group of resource.
	r := core.Router().Group("w/{{ .Folder }}")

	// Register route for {{ ToLower .Entity }} page
	r.GET("{{ ToLower .Entity }}", ctrl.{{ .Entity }})
}

// {{ .Entity }} is a handler for {{ ToLower .Entity }} page
func (c *Controller) {{ .Entity }}(ctx *gin.Context) {
	helpers.View(ctx, welcome())
}
`

// ControllerTemplateHTMLWeb is a template for a web controller
const ControllerTemplateHTMLWeb = `// Package {{ ToLower .Entity }} defines a template
package {{ ToLower .Entity }}

import "{{ .Module }}/web/layouts"

templ welcome() {
	@layouts.Page("Welcome to {{ ToLower .Entity }} page") {
		<h1>Welcome to {{ ToLower .Entity }} page!</h1>
	}
}`

// CoreHelpersView is a template for core helpers view
const CoreHelpersView = `// Package helpers provides helper functions for views and templates
package helpers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// View renders a templ component to the gin context
func View(ctx *gin.Context, component templ.Component) {
	component.Render(ctx.Request.Context(), ctx.Writer)
}`

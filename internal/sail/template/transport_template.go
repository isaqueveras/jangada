// Package template contains templates for a {{ .Layer }} controller
package template

// ControllerTemplate is a template for a {{ .Layer }} controller
const ControllerTemplate = `// Package {{ ToLower .Entity }} defines a {{ .Layer }} controller 
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

// HandlerLayerController is a template for a {{ .Layer }} controller
const HandlerLayerController = `// Package {{ .Layer }} defines the transport layer handlers and routes.
package {{ ToLower .Layer }}

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/application"
	"{{ .Module }}/pkg/database"

	customerApp "{{ .Module }}/internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
	customerTransport "{{ .Module }}/internal/transport/{{ ToLower .Layer }}/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core, conn database.ConnectionPool) {
	{{ ToLower .Entity }}Transport.NewController(core, customerApp.NewOrchestrator(conn, application.New()))
}
`

// ControllerMethod is a template for a {{ .Layer }} controller
const ControllerMethod string = `// {{ .Method }} ...
func (c *Controller) {{ .Method }}(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
`

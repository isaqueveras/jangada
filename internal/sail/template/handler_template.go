// Package template contains templates for all controllers
package template

// HandlerController is a template for all controllers
const HandlerController = `// Package transport defines the transport layer handlers and routes.
package transport

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/transport/{{ .Layer }}"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core) {
	{{ .Layer }}.Handler(core)
}
`

// HandlerLayerController is a template for a {{ .Layer }} controller
const HandlerLayerController = `// Package {{ .Layer }} defines the transport layer handlers and routes.
package {{ .Layer }}

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/transport/{{ .Layer }}/{{ .Folder }}/controller"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core) {
	controller.New{{ .Entity }}Controller(core, nil)
}
`

// Package template contains templates for a web controller
package template

// WebHandlerController is a template for a web controller
const WebHandlerController = `// Package web defines the transport layer handlers and routes.
package web

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/transport/{{ .Layer }}/{{ .Folder }}/controller"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core) {
	controller.New{{ .Entity }}Controller(core, nil)
}
`

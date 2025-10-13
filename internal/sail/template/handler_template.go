// Package template contains templates for all controllers
package template

// HandlerController is a template for all controllers
const HandlerController = `// Package transport defines the transport layer handlers and routes.
package transport

import (
	"{{ .Module }}/core"
	"{{ .Module }}/internal/transport/web"
)

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core) {
	web.Handler(core)
}
`

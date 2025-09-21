// Package template contains templates for a web controller
package template

// WebHandlerController is a template for a web controller
const WebHandlerController = `// Package web defines the transport layer handlers and routes.
package web

import "{{ .Module }}/core"

// Handler builds transport handlers and registers them to the core router.
func Handler(core *core.Core) {
	v1 := core.Router().Group("")
	_ = v1
}
`

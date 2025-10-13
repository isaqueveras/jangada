// Package sail provides commands to create layers for a bounded context.
package sail

import "github.com/isaqueveras/jangada/internal/sail/template"

// Template defines the template structure
type Template struct {
	Path    string
	Content string
}

// WebTransportTemplate holds templates for generating transport layer files
var WebTransportTemplate = []Template{
	{"internal/transport/handler.go", template.HandlerController},
	{"internal/transport/web/handler.go", template.WebHandlerController},

	{"internal/transport/{{ .Layer }}/{{ .Folder }}/controller/{{ ToLower .Entity }}_controller.go", template.WebController},
	{"internal/transport/{{ .Layer }}/{{ .Folder }}/request/{{ ToLower .Entity }}_request.go", template.WebRequest},

	{"internal/application/{{ .Folder }}/orchestrator/{{ ToLower .Entity }}_orchestrator.go", template.ApplicationOrchestrator},
	{"internal/application/{{ .Folder }}/command/{{ ToLower .Entity }}_command.go", template.ApplicationCommand},
	{"internal/application/{{ .Folder }}/query/{{ ToLower .Entity }}_query.go", template.ApplicationQuery},

	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/mapper/{{ ToLower .Entity }}_mapper.go", `package mapper`},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/response/{{ ToLower .Entity }}_response.go", `package response`},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/view/{{ ToLower .Entity }}_view.go", `package view`},
}

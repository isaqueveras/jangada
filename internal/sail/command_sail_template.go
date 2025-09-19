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
	{"internal/transport/{{ .Layer }}/{{ .Folder }}/controller/{{ .Entity }}_controller.go", template.WebController},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/controller/{{ .Entity }}_controller_test.go", template.WebControllerTest},

	{"internal/transport/{{ .Layer }}/{{ .Folder }}/mapper/{{ .Entity }}_mapper.go", `package mapper`},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/mapper/{{ .Entity }}_mapper_test.go", `package mapper`},

	{"internal/transport/{{ .Layer }}/{{ .Folder }}/response/{{ .Entity }}_response.go", `package response`},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/response/{{ .Entity }}_response_test.go", `package response`},

	{"internal/transport/{{ .Layer }}/{{ .Folder }}/request/{{ .Entity }}_request.go", `package request`},
	// {"internal/transport/{{ .Layer }}/{{ .Folder }}/request/{{ .Entity }}_request_test.go", `package request`},

	{"internal/transport/{{ .Layer }}/{{ .Folder }}/view/{{ .Entity }}_view.go", `package view`},
}

// Package sail provides commands to create layers for a bounded context.
package sail

// Template defines the template structure
type Template struct {
	Path    string
	Content string
}

// TemplateTransport holds templates for generating transport layer files
var TemplateTransport = []Template{
	{"internal/transport/{{ .layer }}/{{ .folder }}/controller/{{ .entity }}_controller.go", `package controller`},
	{"internal/transport/{{ .layer }}/{{ .folder }}/controller/{{ .entity }}_controller_test.go", `package controller`},

	{"internal/transport/{{ .layer }}/{{ .folder }}/mapper/{{ .entity }}_mapper.go", `package mapper`},
	{"internal/transport/{{ .layer }}/{{ .folder }}/mapper/{{ .entity }}_mapper_test.go", `package mapper`},

	{"internal/transport/{{ .layer }}/{{ .folder }}/response/{{ .entity }}_response.go", `package response`},
	{"internal/transport/{{ .layer }}/{{ .folder }}/response/{{ .entity }}_response_test.go", `package response`},

	{"internal/transport/{{ .layer }}/{{ .folder }}/request/{{ .entity }}_request.go", `package request`},
	{"internal/transport/{{ .layer }}/{{ .folder }}/request/{{ .entity }}_request_test.go", `package request`},

	{"internal/transport/{{ .layer }}/{{ .folder }}/view/{{ .entity }}_view.go", `package view`},
}

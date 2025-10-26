// Package template contains templates for a web request
package template

const TransportModel = `// Package request defines a request for {{ .Entity }}
package {{ ToLower .Entity }}

import "{{ .Module }}/internal/application/{{ .Folder }}/command"

// Get{{ .Entity }}Params is a request to get a {{ ToLower .Entity }} by id
type Get{{ .Entity }}Params struct {
	// ID is the id of the {{ ToLower .Entity }}
	ID *string ` + "`uri:\"id\" binding:\"required\"`" + `
}

// ToCommand converts the {{ .Entity }}Params to a command
func (p *Get{{ .Entity }}Params) ToCommand() *string {
	return p.ID
}

// {{ .Entity }}Filters is a request for filtering {{ ToLower .Entity }}
type {{ .Entity }}Filters struct {
	ID   *string ` + "`form:\"id\"`" + `
	Name *string ` + "`form:\"name\"`" + `
}

// ToCommand converts the {{ .Entity }}Filters to a command
func (p *{{ .Entity }}Filters) ToCommand() map[string]any {
	return map[string]any{
		"id":   p.ID,
		"name": p.Name,
	}
}

// {{ .Entity }}CreateParams is a request for creating a {{ ToLower .Entity }}
type {{ .Entity }}CreateParams struct {
	Name *string ` + "`json:\"name\" binding:\"required\"`" + `
}

// ToCommand converts the {{ .Entity }}CreateParams to a command
func (p *{{ .Entity }}CreateParams) ToCommand() *command.Create{{ .Entity }} {
	return &command.Create{{ .Entity }}{
		Name: p.Name,
	}
}

// {{ .Entity }}UpdateParams is a request for updating a {{ ToLower .Entity }}
type {{ .Entity }}UpdateParams struct {
	ID   *string ` + "`uri:\"id\" binding:\"required\"`" + `
	Name *string ` + "`json:\"name\" binding:\"required\"`" + `
}

// ToCommand converts the {{ .Entity }}UpdateParams to a command
func (p *{{ .Entity }}UpdateParams) ToCommand() *command.Update{{ .Entity }} {
	return &command.Update{{ .Entity }}{
		ID:   p.ID,
		Name: p.Name,
	}
}

// {{ .Entity }}DeleteParams is a request for deleting a {{ ToLower .Entity }}
type {{ .Entity }}DeleteParams struct {
	ID *string ` + "`uri:\"id\" binding:\"required\"`" + `
}

// ToCommand converts the {{ .Entity }}DeleteParams to a command
func (p *{{ .Entity }}DeleteParams) ToCommand() *string { 
	return p.ID
}
`

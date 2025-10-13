package template

// ApplicationCommand is a template for an application command
const ApplicationCommand = `// Package command defines a command for {{ .Entity }}
package command

// Create{{ .Entity }} is a command to create a new {{ ToLower .Entity }}
type Create{{ .Entity }} struct {
	ID   *string
	Name *string
}

// Update{{ .Entity }} is a command to update a {{ ToLower .Entity }}
type Update{{ .Entity }} struct {
	ID   *string
	Name *string
}
`

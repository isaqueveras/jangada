package template

// ApplicationQuery is a template for an application query
const ApplicationQuery = `// Package query defines a query for {{ ToLower .Entity }}
package query

// {{ .Entity }}Item is a {{ ToLower .Entity }} item
type {{ .Entity }}Item struct {
	ID   *string
	Name *string
}

// List{{ .Entity }}Item is a list of {{ ToLower .Entity }} items
type List{{ .Entity }}Item struct {
	Items []*{{ .Entity }}Item
}
`

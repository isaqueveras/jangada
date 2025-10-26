// Package newapp contains templates for new app
package newapp

const tmplAppMain = `// Package main contains the main function of the app
package main

import (
	"{{ .ModuleName }}/core"
	"{{ .ModuleName }}/pkg/database"
)

func main() {
	server := core.New()

	conn := database.NewConnectionPool(server.Config().GetDatabases()...)
	defer conn.CloseConnections()

	// Uncomment the code below to enable the transport layer
	// transport.Handler(server)

	server.Init()
}
`

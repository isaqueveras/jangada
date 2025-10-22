package newapp

const tmplAppMain = `package main

import (
	"{{ .ModuleName }}/core"
	"{{ .ModuleName }}/services/database"
)

func main() {
	server := core.New()

	conn := database.NewConnectionPool(server.Config().GetDatabases()...)
	defer conn.CloseConnections()

	server.Init()
}`

package newapp

const tmplAppMain = `package main

import "{{ .ModuleName }}/core"

func main() {
	server := core.New()
	server.Init()
}`

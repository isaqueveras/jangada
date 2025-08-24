package newapp

import "strings"

var tmplReadme = strings.ReplaceAll(tmplReadmeTemp, "'", "`")

const tmplReadmeTemp string = `# Jangada Framework

A full-stack web framework in Go for building modern web applications, RESTful APIs, and gRPC with integrated frontend and backend.

## Setup

'''bash
# Install Jangada
go install github.com/isaqueveras/jangada@latest
'''

## Run Jangada

'''bash
# Run development server
jangada serve
'''

## Run Jangada with Docker

'''bash
# Run development server
docker compose up
'''

## Tests and Coverage

'''bash
# Run tests
jangada test
'''

## Documentation

[Docs](https://jangada-framework.com/docs)
`

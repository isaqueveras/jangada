// Package sail provides commands to create layers for a bounded context.
package sail

import "github.com/isaqueveras/jangada/internal/sail/template"

type info struct{ Layer, Folder, Entity, Module, Method string }

// Template defines the template structure
type Template struct {
	path, content string
	canModify     bool
}

var applicationTemplate = []Template{
	{path: "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/orchestrator.go", content: template.ApplicationOrchestrator},
}

var (
	domainPath     = "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}"
	domainTemplate = []Template{
		{path: domainPath + "/entity.go", content: "package {{ ToLower .Entity }}"},
		{path: domainPath + "/factory.go", content: "package {{ ToLower .Entity }}"},
		{path: domainPath + "/valueobject.go", content: "package {{ ToLower .Entity }}"},
		{path: domainPath + "/repository.go", content: template.DomainRepository},
		{path: domainPath + "/service.go", content: template.DomainService},
	}
)

var infrastructureTemplate = []Template{
	{path: "internal/infrastructure/{{ ToLower .Folder }}/{{ ToLower .Entity }}/repository.go", content: template.InfrastructureRepository},
	{path: "internal/infrastructure/{{ ToLower .Folder }}/{{ ToLower .Entity }}/postgres/data.go", content: template.InfrastructurePostgresData},
	{path: "internal/infrastructure/{{ ToLower .Folder }}/{{ ToLower .Entity }}/postgres/model.go", content: template.InfrastructurePostgresModel},
}

var transportTemplateRest = []Template{
	{path: "internal/transport/handler.go", content: template.HandlerController},
	{path: "internal/transport/rest/handler.go", content: template.HandlerLayerController},
	{path: "internal/transport/rest/{{ ToLower .Folder }}/{{ ToLower .Entity }}/controller.go", content: template.ControllerTemplateRest, canModify: true},
}

var transportTemplateWeb = []Template{
	{path: "core/helpers/view.go", content: template.CoreHelpersView},
	{path: "internal/transport/handler.go", content: template.HandlerController},
	{path: "internal/transport/web/handler.go", content: template.HandlerLayerController},
	{path: "internal/transport/web/{{ ToLower .Folder }}/{{ ToLower .Entity }}/controller.go", content: template.ControllerTemplateWeb, canModify: true},
	{path: "internal/transport/web/{{ ToLower .Folder }}/{{ ToLower .Entity }}/model.go", content: "package {{ ToLower .Entity }}"},
	{path: "internal/transport/web/{{ ToLower .Folder }}/{{ ToLower .Entity }}/template.templ", content: template.ControllerTemplateHTMLWeb},
}

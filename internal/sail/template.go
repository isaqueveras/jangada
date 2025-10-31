// Package sail provides commands to create layers for a bounded context.
package sail

import "github.com/isaqueveras/jangada/internal/sail/template"

type info struct {
	Layer, Folder, Entity, Module, Method string
}

// Template defines the template structure
type Template struct {
	Path, Content string
	CanModify     bool
}

var transportTemplateRest = []Template{
	{
		Path:    "internal/transport/handler.go",
		Content: template.HandlerController,
	},
	{
		Path:    "internal/transport/rest/handler.go",
		Content: template.HandlerLayerController,
	},
	{
		Path:      "internal/transport/rest/{{ ToLower .Folder }}/{{ ToLower .Entity }}/controller.go",
		Content:   template.ControllerTemplate,
		CanModify: true,
	},
	// {
	// 	Path:    "internal/transport/rest/{{ ToLower .Folder }}/{{ ToLower .Entity }}/controller_test.go",
	// 	Content: template.ControllerTestTemplate,
	// },
}

var applicationTemplate = []Template{
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/service.go",
		Content: template.ApplicationService,
	},
	// {
	// 	Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/service_test.go",
	// 	Content: "package {{ ToLower .Entity }}_test",
	// },
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/orchestrator.go",
		Content: template.ApplicationOrchestrator,
	},
	// {
	// 	Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/orchestrator_test.go",
	// 	Content: "package {{ ToLower .Entity }}_test",
	// },

	// {
	// 	Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/model.go",
	// 	Content: "package {{ ToLower .Entity }}",
	// },
	// {
	// 	Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/mapper.go",
	// 	Content: "package {{ ToLower .Entity }}",
	// },
}

var domainTemplate = []Template{
	{
		Path:    "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}/entity.go",
		Content: "package {{ ToLower .Entity }}",
	},
	{
		Path:    "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}/factory.go",
		Content: "package {{ ToLower .Entity }}",
	},
	{
		Path:    "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}/interface.go",
		Content: "package {{ ToLower .Entity }}",
	},
	{
		Path:    "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}/service.go",
		Content: "package {{ ToLower .Entity }}",
	},
	{
		Path:    "internal/domain/{{ ToLower .Folder }}/{{ ToLower .Entity }}/valueobject.go",
		Content: "package {{ ToLower .Entity }}",
	},
}

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
		Content:   template.ControllerTemplateClean,
		CanModify: true,
	},
	{
		Path:    "internal/transport/rest/{{ ToLower .Folder }}/{{ ToLower .Entity }}/controller_test.go",
		Content: template.ControllerTestTemplate,
	},
}

var applicationTemplate = []Template{
	{
		Path:    "internal/application/builder.go",
		Content: template.ApplicationBuilder,
	},
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/service.go",
		Content: template.ApplicationService,
	},
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/orchestrator.go",
		Content: template.ApplicationOrchestrator,
	},
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/model.go",
		Content: "package {{ ToLower .Entity }}",
	},
	{
		Path:    "internal/application/{{ ToLower .Folder }}/{{ ToLower .Entity }}/mapper.go",
		Content: "package {{ ToLower .Entity }}",
	},
}

// transportTemplateCRUD holds templates for generating transport layer files for crud
// var transportTemplateCRUD = []Template{
// {"internal/transport/handler.go", template.HandlerController, false},
// {"internal/transport/{{ .Layer }}/handler.go", template.HandlerLayerController, false},

// {"internal/transport/{{ .Layer }}/{{ .Folder }}/controller/{{ ToLower .Entity }}_controller.go", template.ControllerTemplate, false},
// {"internal/transport/{{ .Layer }}/{{ .Folder }}/request/{{ ToLower .Entity }}_request.go", template.TransportModel, false},

// {"internal/application/{{ .Folder }}/orchestrator/{{ ToLower .Entity }}_orchestrator.go", template.ApplicationOrchestrator, false},
// {"internal/application/{{ .Folder }}/command/{{ ToLower .Entity }}_command.go", template.ApplicationCommand, false},
// {"internal/application/{{ .Folder }}/query/{{ ToLower .Entity }}_query.go", template.ApplicationQuery, false},

// {"internal/transport/{{ .Layer }}/{{ .Folder }}/mapper/{{ ToLower .Entity }}_mapper.go", `package mapper`, false},
// {"internal/transport/{{ .Layer }}/{{ .Folder }}/response/{{ ToLower .Entity }}_response.go", `package response`, false},
// {"internal/transport/{{ .Layer }}/{{ .Folder }}/view/{{ ToLower .Entity }}_view.go", `package view`, false},
// }

// transportTemplateCreateMethod holds templates for generating transport layer files for creating a method
// var transportTemplateCreateMethod = []Template{
// {"internal/transport/handler.go", template.HandlerController, false},
// {"internal/transport/{{ .Layer }}/handler.go", template.HandlerLayerController, false},

// {"internal/transport/{{ .Layer }}/{{ .Folder }}/controller/{{ ToLower .Entity }}_controller.go", template.ControllerMethod, true},
// {"internal/transport/{{ .Layer }}/{{ .Folder }}/request/{{ ToLower .Entity }}_request.go", template.Request, false},
// }

package template

import newapp "github.com/isaqueveras/jangada/internal/template/new-app"

// GetTemplateForNewApp returns a map of templates for new app
func GetTemplateForNewApp() map[string]string { return newapp.Template }

// GetTemplateForNewWorker returns a map of templates for new worker
func GetTemplateForNewWorker() map[string]string { return nil }

// GetTemplateForNewConsole returns a map of templates for new console
func GetTemplateForNewConsole() map[string]string { return nil }

// GetTemplateForNewJob returns a map of templates for new job
func GetTemplateForNewJob() map[string]string { return nil }

// GetTemplateForNewApplicationLayer returns a map of templates for new application layer
func GetTemplateForNewApplicationLayer() map[string]string { return nil }

// GetTemplateForNewDomainLayer returns a map of templates for new domain layer
func GetTemplateForNewDomainLayer() map[string]string { return nil }

// GetTemplateForNewInfrastructureLayer returns a map of templates for new infrastructure layer
func GetTemplateForNewInfrastructureLayer() map[string]string { return nil }

// GetTemplateForNewInterfaceLayer returns a map of templates for new interface layer
func GetTemplateForNewInterfaceLayer() map[string]string { return nil }

// GetTemplateForNewTest returns a map of templates for new test
func GetTemplateForNewTest() map[string]string { return nil }

// GetTemplateForAddPackage returns a map of templates for add package
func GetTemplateForAddPackage() map[string]string { return nil }

// GetTemplateForNewPlugin returns a map of templates for new plugin
func GetTemplateForNewPlugin() map[string]string { return nil }

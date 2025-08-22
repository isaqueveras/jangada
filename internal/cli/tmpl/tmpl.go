package tmpl

import newapp "github.com/isaqueveras/jangada/internal/cli/tmpl/new-app"

// GetTemplateForNewProject returns a map of templates for new project
func GetTemplateForNewProject() map[string]newapp.TmplNewProject {
	return newapp.TemplatesNewProject
}

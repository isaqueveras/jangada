package newapp

type TmplNewProject string

func (tmpl TmplNewProject) String() string {
	return string(tmpl)
}

var TemplatesNewProject = map[string]TmplNewProject{
	"go.mod":     tmplGoMod,
	".gitignore": tmplGitIgnore,
	"Makefile":   tmplMakefile,

	"core/core.go":   tmplCoreCore,
	"core/plugin.go": tmplCorePlugin,

	// "bin/": "",
	// ".devcontainer/devcontainer.yaml": "",

	// "config/environment/development.go": "package environment",
	// "config/environment/production.go":  "package environment",
	// "config/environment/test.go":        "package environment",

	"config/config.go":     tmplConfigConfig,
	"config/database.yaml": tmplConfigDatabaseYAML,
	"config/storage.yaml":  tmplConfigStorageYaml,

	"cmd/app/main.go": tmplAppMain,

	"db/seeds.go": tmplDBSeeds,
	// "db/schema.go": "package db",

	// "docs/docs.go":     "package docs",
	// "helper/helper.go": "package helper",

	"log/access.log": "",
	"log/error.log":  "",

	// "internal/.keep": "",
	"public/.keep": "",

	"web/assets/css/input.css":  "",
	"web/assets/css/output.css": "",
	"web/assets/js/.keep":       "",
	"web/components/.keep":      "",

	"web/layouts/welcome.templ":   tmplWebLayoutsWelcome,
	"web/layouts/page.templ":      tmplWebLayoutsPage,
	"web/layouts/not_found.templ": tmplWebLayoutsNotFound,

	// "web/utils/templui.go": "",
}

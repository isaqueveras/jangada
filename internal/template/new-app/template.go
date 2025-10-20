// Package newapp contains templates for new app
package newapp

// Template is a map of templates for new app
var Template = map[string]string{
	"go.mod":     tmplGoMod,     // contains module of the app
	".gitignore": tmplGitIgnore, // contains gitignore of the app
	"Makefile":   tmplMakefile,  // contains makefile of the app
	"README.md":  tmplReadme,    // contains readme of the app

	"core/core.go":   tmplCoreCore,   // contains core of the app
	"core/plugin.go": tmplCorePlugin, // contains plugin of the app

	"log/access.log": "", // contains access log of the app
	"log/error.log":  "", // contains error log of the app

	"config/config.go":     tmplConfigConfig,       // contains config of the app
	"config/database.yaml": tmplConfigDatabaseYAML, // contains database of the app

	"services/database/interface.go":         serviceDatabaseInterfaceTemp,
	"services/database/pool.go":              serviceDatabasePoolTemp,
	"services/database/transaction.go":       serviceDatabaseTransactionTemp,
	"services/database/postgres/postgres.go": serviceDatabasePostgresTemp,
	"services/database/seeds/seeds.go":       serviceDatabaseSeedsTemp,
	"services/database/migrations/.keep":     "",

	"cmd/app/main.go": tmplAppMain, // contains main of the app

	"public/.keep": "", // contains public files of the app

	"web/assets/css/input.css":  "", // contains input css of the app
	"web/assets/css/output.css": "", // contains output css of the app
	"web/assets/js/.keep":       "", // contains js files of the app
	"web/components/.keep":      "", // contains components of the app

	"web/layouts/welcome.templ":   tmplWebLayoutsWelcome,  // contains welcome page of the app
	"web/layouts/page.templ":      tmplWebLayoutsPage,     // contains page of the app
	"web/layouts/not_found.templ": tmplWebLayoutsNotFound, // contains not found page of the app
}

// "bin/": "",
// ".devcontainer/devcontainer.yaml": "",

// "config/environment/development.go": "package environment",
// "config/environment/production.go":  "package environment",
// "config/environment/test.go":        "package environment",

// "cmd/app/app.go":  tmplAppApp,  // contains app of the app

// "db/schema.go": "package db",

// "docs/docs.go":     "package docs",
// "helper/helper.go": "package helper",

// "internal/.keep": "",

// "web/utils/templui.go": "", // contains templui of the app

// Package newapp contains templates for new app
package newapp

// Template is a map of templates for new app
var Template = map[string]string{
	"go.mod":     tmplGoMod,     // contains module of the app
	".gitignore": tmplGitIgnore, // contains gitignore of the app
	"Makefile":   tmplMakefile,  // contains makefile of the app
	"README.md":  tmplReadme,    // contains readme of the app

	"core/core.go": tmplCoreCore, // contains core of the app
	// "core/plugin.go": tmplCorePlugin, // contains plugin of the app

	"log/access.log": "", // contains access log of the app
	"log/error.log":  "", // contains error log of the app

	"config/config.go":     tmplConfigConfig,       // contains config of the app
	"config/database.yaml": tmplConfigDatabaseYAML, // contains database of the app
	"config/app.yaml":      tmplConfigAppYAML,      // contains app of the app

	"pkg/database/interface.go":         serviceDatabaseInterfaceTemp,
	"pkg/database/pool.go":              serviceDatabasePoolTemp,
	"pkg/database/transaction.go":       serviceDatabaseTransactionTemp,
	"pkg/database/postgres/postgres.go": serviceDatabasePostgresTemp,
	"pkg/database/seeds/seeds.go":       serviceDatabaseSeedsTemp,
	"pkg/database/migrations/.keep":     "",

	"cmd/app/main.go": tmplAppMain, // contains main of the app

	"public/.keep": "", // contains public files of the app

	"web/assets/css/input.css":  "", // contains input css of the app
	"web/assets/css/output.css": "", // contains output css of the app
	"web/assets/js/.keep":       "", // contains js files of the app
	"web/components/.keep":      "", // contains components of the app

	"web/layouts/welcome.templ":   tmplWebLayoutsWelcome,  // contains welcome page of the app
	"web/layouts/page.templ":      tmplWebLayoutsPage,     // contains page of the app
	"web/layouts/not_found.templ": tmplWebLayoutsNotFound, // contains not found page of the app

	"internal/application/builder.go":    tmplApplicationBuilder,    // contains builder of the app
	"internal/domain/builder.go":         tmplDomainBuilder,         // contains builder of the domain
	"internal/infrastructure/builder.go": tmplInfrastructureBuilder, // contains builder of the infrastructure
}

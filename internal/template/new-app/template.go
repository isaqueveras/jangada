// Package newapp contains templates for new app
package newapp

// Template is a map of templates for new app
var Template = map[string]string{
	"go.mod":       tmplGoMod,     // contains module of the app
	".air.toml":    tmplAirToml,   // contains air.toml of the app
	".gitignore":   tmplGitIgnore, // contains gitignore of the app
	".env":         tmplEnv,       // contains env of the app
	".env.example": tmplEnv,       // contains env of example of the app
	"Makefile":     tmplMakefile,  // contains makefile of the app
	"README.md":    tmplReadme,    // contains readme of the app

	"core/core.go": tmplCoreCore, // contains core of the app

	"config/config.go": tmplConfigConfig, // contains config of the app

	"pkg/metric/metric.go": tmplMetricMetric, // contains metric of the app

	"pkg/database/interface.go":         serviceDatabaseInterfaceTemp,
	"pkg/database/pool.go":              serviceDatabasePoolTemp,
	"pkg/database/transaction.go":       serviceDatabaseTransactionTemp,
	"pkg/database/postgres/postgres.go": serviceDatabasePostgresTemp,
	"pkg/database/seeds/seeds.go":       serviceDatabaseSeedsTemp,
	"pkg/database/migrations/.keep":     "",

	"docker/Dockerfile":         tmplDockerfile,
	"docker/prometheus.yml":     tmplPrometheus,
	"docker/docker-compose.yml": tmplDockerCompose,

	"cmd/app/main.go": tmplAppMain, // contains main of the app

	"public/.keep": "", // contains public files of the app

	"web/assets/css/input.css":  "", // contains input css of the app
	"web/assets/css/output.css": "", // contains output css of the app
	"web/assets/js/.keep":       "", // contains js files of the app
	"web/components/.keep":      "", // contains components of the app

	"web/layouts/welcome.templ":   tmplWebLayoutsWelcome,  // contains welcome page of the app
	"web/layouts/page.templ":      tmplWebLayoutsPage,     // contains page of the app
	"web/layouts/not_found.templ": tmplWebLayoutsNotFound, // contains not found page of the app

	"internal/application/service.go":       tmplApplicationService,       // contains service builder of the app
	"internal/domain/interface.go":          tmplDomainInterface,          // contains interface of the domain
	"internal/infrastructure/repository.go": tmplInfrastructureRepository, // contains repository builder of the infrastructure
}

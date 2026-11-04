package application

import (
	"log/slog"
	"os"
)

// Environment represents the environment in which the application is running.
type Environment string

// String returns the string representation of the Environment type.
func (e Environment) String() string {
	return string(e)
}

const (
	// Development represents the development environment.
	//
	// Example usage:
	// app := application.New("MyApp", func(s *application.Settings) {
	//     s.SetEnvironment(application.Development)
	// })
	// fmt.Println(app.Environment())
	Development Environment = "development"

	// Staging represents the staging environment.
	//
	// Example usage:
	// app := application.New("MyApp", func(s *application.Settings) {
	//     s.SetEnvironment(application.Staging)
	// })
	// fmt.Println(app.Environment())
	Staging Environment = "staging"

	// Production represents the production environment.
	//
	// Example usage:
	// app := application.New("MyApp", func(s *application.Settings) {
	//     s.SetEnvironment(application.Production)
	// })
	// fmt.Println(app.Environment())
	Production Environment = "production"
)

// Application represents the core structure of an application,
// holding its name, version, environment, and logger.
type Application struct {
	name        string
	version     string
	environment Environment
	logger      *slog.Logger
	debug       bool
}

// Name returns the name of the application.
func (a *Application) Name() string {
	return a.name
}

// Version returns the current version of the application.
func (a *Application) Version() string {
	return a.version
}

// Environment returns the current environment of the application (development, staging, or production).
func (a *Application) Environment() Environment {
	return a.environment
}

// IsDebug returns whether the application is running in debug mode.
func (a *Application) IsDebug() bool {
	return a.debug
}

// Logger returns the logger instance associated with the application.
func (a *Application) Logger() *slog.Logger {
	return a.logger
}

// New creates a new Application instance with the provided name and optional settings.
//
// The optional settings can be provided as functions that modify the Settings struct.
// If no settings are provided, default values will be used.
//
// Example usage:
//
//	app := application.New("MyApp", func(s *application.Settings) {
//		s.SetVersion("v1.0.0")
//		s.SetEnvironment(application.Production)
//		s.SetDebug(false)
//		s.SetLogger(func(level slog.Leveler) slog.Handler {
//			return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
//				AddSource: true,
//				Level:     level,
//			})
//		})
//	})
func New(name string, opts ...func(*Settings)) *Application {
	settings := &Settings{}
	for _, opt := range opts {
		opt(settings)
	}

	app := &Application{
		name:        name,
		version:     settings.version,
		environment: settings.environment,
		debug:       settings.debug,
	}

	if app.version == "" {
		app.version = "0.0.1"
	}

	if app.environment == "" {
		app.environment = Development
	}

	{ // Initialize the logger based on the environment
		levelLogger := slog.LevelVar{}
		levelLogger.Set(slog.LevelInfo)
		if app.debug {
			levelLogger.Set(slog.LevelDebug)
		}

		if settings.logger != nil {
			app.logger = slog.New(settings.logger(&levelLogger))
		} else {
			// If a custom logger is provided in the settings,
			// use it; otherwise, use the default logger.
			var handler slog.Handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
				Level: &levelLogger,
			})

			// If the application is running in the development environment,
			// use a text handler for better readability.
			if app.Environment() == Development {
				handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
					Level:     &levelLogger,
					AddSource: true,
				})
			}

			app.logger = slog.New(handler)
		}

		app.logger = app.logger.With("application", app.name)
		app.logger = app.logger.With("version", app.version)
		app.logger = app.logger.With("environment", app.environment)
		app.logger = app.logger.With("debug", app.debug)
	}

	return app
}

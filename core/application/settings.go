package application

import "log/slog"

// Settings holds the configuration for the Application instance.
// It allows setting the version and environment of the application.
type Settings struct {
	version     string
	environment Environment
	debug       bool
	logger      func(level slog.Leveler) slog.Handler
}

// SetVersion sets the service version.
//
//	app := application.New(func(s *application.Settings) {
//		s.SetVersion("v0.0.1")
//	})
func (s *Settings) SetVersion(version string) {
	s.version = version
}

// SetEnvironment sets the service environment.
//
//	app := application.New(func(s *application.Settings) {
//		s.SetEnvironment(application.Production)
//	})
func (s *Settings) SetEnvironment(environment Environment) {
	s.environment = environment
}

// SetDebug sets the debug mode for the application.
//
//	app := application.New(func(s *application.Settings) {
//		s.SetDebug(true)
//	})
func (s *Settings) SetDebug(debug bool) {
	s.debug = debug
}

// SetLogger sets the logger for the application.
//
//	app := application.New(func(s *application.Settings) {
//		s.SetLogger(func(level slog.Leveler) slog.Handler {
//			return slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
//				AddSource: true,
//				Level:     level,
//			})
//		})
//	})
func (s *Settings) SetLogger(logger func(level slog.Leveler) slog.Handler) {
	s.logger = logger
}

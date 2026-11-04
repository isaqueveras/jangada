package application_test

import (
	"log/slog"
	"testing"

	"github.com/isaqueveras/jangada/core/application"
)

func TestNewApplication(t *testing.T) {
	app := application.New("TestNewApplication", func(s *application.Settings) {
		s.SetEnvironment(application.Development)
		s.SetVersion("v0.0.1-dev")
		s.SetDebug(true)
		s.SetLogger(func(level slog.Leveler) slog.Handler {
			return slog.DiscardHandler
		})
	})

	if app.Name() != "TestNewApplication" {
		t.Errorf("Expected application name to be 'TestNewApplication', got '%s'", app.Name())
	}

	if app.Version() != "v0.0.1-dev" {
		t.Errorf("Expected application version to be 'v0.0.1-dev', got '%s'", app.Version())
	}

	if app.Environment() != application.Development {
		t.Errorf("Expected application environment to be 'development', got '%s'", app.Environment())
	}

	if !app.IsDebug() {
		t.Errorf("Expected application to be in debug mode, but it is not")
	}

	if app.Logger() == nil {
		t.Errorf("Expected application logger to be initialized, but it is nil")
	}
}

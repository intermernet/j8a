package j8a

import (
	"github.com/rs/zerolog"
	"os"
	"testing"
)

func TestServerID(t *testing.T) {
	os.Setenv("HOSTNAME", "localhost")
	Version = "v0.0.0"
	initServerID()
	want := "f47f7b28"
	if ID != want {
		t.Errorf("serverID did not properly compute, got %v, want %v", ID, want)
	}
}

func TestDefaultLogLevelInit(t *testing.T) {
	os.Setenv("LOGLEVEL", "not set")
	initLogger()
	got := zerolog.GlobalLevel().String()
	want := "info"
	if got != want {
		t.Errorf("default log level not properly initialised, got %v, want %v", got, want)
	}
}

func TestTraceLogLevelInit(t *testing.T) {
	os.Setenv("LOGLEVEL", "TRACE")
	initLogger()
	got := zerolog.GlobalLevel().String()
	want := "trace"
	if got != want {
		t.Errorf("default log level not properly initialised, got %v, want %v", got, want)
	}
}

func TestDebugLogLevelInit(t *testing.T) {
	os.Setenv("LOGLEVEL", "DEBUG")
	initLogger()
	got := zerolog.GlobalLevel().String()
	want := "debug"
	if got != want {
		t.Errorf("log level not properly initialised, got %v, want %v", got, want)
	}
}

func TestInfoLogLevelInit(t *testing.T) {
	os.Setenv("LOGLEVEL", "INFO")
	initLogger()
	got := zerolog.GlobalLevel().String()
	want := "info"
	if got != want {
		t.Errorf("log level not properly initialised, got %v, want %v", got, want)
	}
}

func TestWarnLogLevelInit(t *testing.T) {
	os.Setenv("LOGLEVEL", "WARN")
	initLogger()
	got := zerolog.GlobalLevel().String()
	want := "warn"
	if got != want {
		t.Errorf("log level not properly initialised, got %v, want %v", got, want)
	}
}

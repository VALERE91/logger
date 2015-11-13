package logger

import (
	"testing"
)

func TestConfigureFileLogger(t *testing.T) {
	var logger, err = ConfigureFileLogger("./test.log")

	if logger == nil || FileLogger == nil {
		t.Error("FileLogger should be not nil")
	}

	logger, err = ConfigureFileLogger("")

	if err == nil {
		t.Error("ConfigureFileLogger : Error with an empty string should not be nil")
	}
}

func TestConfigureMongoLogger(t *testing.T) {
	var logger, err = ConfigureMongoLogger("localhost", "test", "logs")

	if err != nil || logger == nil {
		t.Error("Error must be nil and logger not")
	}

	logger, err = ConfigureMongoLogger("stupid_cdn", "test", "logs")

	if err == nil || logger != nil {
		t.Error("Without connection logger must be nil and error set properly")
	}
}

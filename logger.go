package logger

import (
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	DebugF(layout string, args ...interface{})
	Log(args ...interface{})
	LogF(layout string, args ...interface{})
	Error(args ...interface{})
	ErrorF(layout string, args ...interface{})
	Fatal(args ...interface{})
	FatalF(layout string, args ...interface{})
}

var MongoDBLogger *mongoLogger = nil
var StdoutLogger *streamLogger = newStreamLogger(os.Stdout)
var StderrLogger *streamLogger = newStreamLogger(os.Stderr)
var FileLogger *streamLogger = nil

func ConfigureFileLogger(path string) (logger *streamLogger, err error) {
	var stream, er = os.Create(path)
	if er != nil {
		return nil, er
	}

	FileLogger = newStreamLogger(stream)
	return FileLogger, nil
}

func ConfigureMongoLogger(path string) (logger *mongoLogger, err error) {
	return nil, nil
}

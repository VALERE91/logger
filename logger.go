package logger
import "os"

type Logger interface{
	Debug(args ...interface{})
	DebugF(layout string, args ...interface{})
	Log(args ...interface{})
	LogF(layout string, args ...interface{})
	Error(args ...interface{})
	ErrorF(layout string, args ...interface{})
	Fatal(args ...interface{})
	FatalF(layout string, args ...interface{})
}

var MongoDBLogger mongoLogger
var StdoutLogger stdLogger = newStdLogger(os.Stdout)
var StderrLogger stdLogger = newStdLogger(os.Stderr)
var FileLogger mongoLogger
package logger

import (
	"os"

	"gopkg.in/mgo.v2"
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

func ConfigureMongoLogger(mongoURL string, database string, collection string) (logger *mongoLogger, err error) {

	if MongoDBLogger != nil {
		MongoDBLogger.session.Close()
	}

	session, er := mgo.Dial(mongoURL)

	if er != nil {
		logger = nil
		MongoDBLogger = nil
		err = er
		return
	}

	c := session.DB(database).C(collection)

	MongoDBLogger = newMongoLogger(session, c)

	return MongoDBLogger, nil
}

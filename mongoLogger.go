package logger

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoLogger struct {
	coll    *mgo.Collection
	session *mgo.Session
}

type logEntry struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Date    time.Time     `bson:"date"`
	Level   string        `bson:"level"`
	Message string        `bson:"message"`
}

func newMongoLogger(session *mgo.Session, c *mgo.Collection) (logger *mongoLogger) {
	return &mongoLogger{
		coll:    c,
		session: session,
	}
}

func (self *mongoLogger) defaultWrite(level string, logLine string) {
	self.coll.Insert(logEntry{
		Date:    time.Now(),
		Level:   level,
		Message: logLine,
	})
}

func (self *mongoLogger) Debug(args ...interface{}) {
	self.defaultWrite("DEBUG", fmt.Sprint(args...))
}

func (self *mongoLogger) DebugF(layout string, args ...interface{}) {
	self.defaultWrite("DEBUG", fmt.Sprintf(layout, args...))
}

func (self *mongoLogger) Log(args ...interface{}) {
	self.defaultWrite("LOG", fmt.Sprint(args...))
}

func (self *mongoLogger) LogF(layout string, args ...interface{}) {
	self.defaultWrite("LOG", fmt.Sprintf(layout, args...))
}

func (self *mongoLogger) Error(args ...interface{}) {
	self.defaultWrite("ERROR", fmt.Sprint(args...))
}

func (self *mongoLogger) ErrorF(layout string, args ...interface{}) {
	self.defaultWrite("ERROR", fmt.Sprintf(layout, args...))
}

func (self *mongoLogger) Fatal(args ...interface{}) {
	self.defaultWrite("FATAL", fmt.Sprint(args...))
}

func (self *mongoLogger) FatalF(layout string, args ...interface{}) {
	self.defaultWrite("FATAL", fmt.Sprintf(layout, args...))
}

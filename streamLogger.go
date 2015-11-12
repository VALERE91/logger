package logger

import (
	"fmt"
	"io"
	"time"
)

type streamLogger struct {
	stream io.Writer
}

func newStreamLogger(stream io.Writer) *streamLogger {
	return &streamLogger{
		stream: stream,
	}
}

func (self *streamLogger) defaultWrite(level string, logLine string) {
	self.stream.Write([]byte(fmt.Sprintf("[%s]%s : %s", time.Now().Format("Mon _2 Jan 2006 15:04"), level, logLine)))
}

func (self *streamLogger) Debug(args ...interface{}) {
	self.defaultWrite("DEBUG", fmt.Sprint(args...))
}

func (self *streamLogger) DebugF(layout string, args ...interface{}) {
	self.defaultWrite("DEBUG", fmt.Sprintf(layout, args...))
}

func (self *streamLogger) Log(args ...interface{}) {
	self.defaultWrite("LOG", fmt.Sprint(args...))
}

func (self *streamLogger) LogF(layout string, args ...interface{}) {
	self.defaultWrite("LOG", fmt.Sprintf(layout, args...))
}

func (self *streamLogger) Error(args ...interface{}) {
	self.defaultWrite("ERROR", fmt.Sprint(args...))
}

func (self *streamLogger) ErrorF(layout string, args ...interface{}) {
	self.defaultWrite("ERROR", fmt.Sprintf(layout, args...))
}

func (self *streamLogger) Fatal(args ...interface{}) {
	self.defaultWrite("FATAL", fmt.Sprint(args...))
}

func (self *streamLogger) FatalF(layout string, args ...interface{}) {
	self.defaultWrite("FATAL", fmt.Sprintf(layout, args...))
}

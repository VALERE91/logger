package logger
import (
	"fmt"
	"time"
	"io"
)

type stdLogger struct{
	stream io.Writer
}

func newStdLogger(stream io.Writer) stdLogger{
	return stdLogger{
		stream: stream,
	}
}

func (self *stdLogger) defaultWrite(level string,logLine string){
	self.stream.Write([]byte(fmt.Sprintf("[%s]%s : %s", time.Now().Format("Mon _2 Jan 2006 15:04"), level, logLine)))
}

func (self *stdLogger) Debug(args ...interface{}){
	self.defaultWrite("DEBUG", fmt.Sprint(args...))
}

func (self *stdLogger) DebugF(layout string, args ...interface{}){
	self.defaultWrite("DEBUG", fmt.Sprintf(layout, args...))
}

func (self *stdLogger) Log(args ...interface{}){
	self.defaultWrite("LOG", fmt.Sprint(args...))
}

func (self *stdLogger) LogF(layout string, args ...interface{}){
	self.defaultWrite("LOG", fmt.Sprintf(layout, args...))
}

func (self *stdLogger) Error(args ...interface{}){
	self.defaultWrite("ERROR", fmt.Sprint(args...))
}

func (self *stdLogger) ErrorF(layout string, args ...interface{}){
	self.defaultWrite("ERROR", fmt.Sprintf(layout, args...))
}

func (self *stdLogger) Fatal(args ...interface{}){
	self.defaultWrite("FATAL", fmt.Sprint(args...))
}

func (self *stdLogger) FatalF(layout string, args ...interface{}){
	self.defaultWrite("FATAL", fmt.Sprintf(layout, args...))
}
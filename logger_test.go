package logger
import "testing"

func TestConfigureFileLogger(t *testing.T) {
	var logger , _ = ConfigureFileLogger("./test.log")

	if logger == nil || FileLogger == nil{
		t.Error("FileLogger should be not nil")
	}
}

func TestConfigureMongoLogger(t *testing.T) {
	var logger , _ = ConfigureMongoLogger("./test.log")

	if logger != nil {
		t.Error("Mongo should be nil (implementation not finished)")
	}
}

func TestTestWriterWrite(t *testing.T) {
	wr := testWriter{}
	wr.Write([]byte("HELLO"))
	if wr.output != "HELLO"{
		t.Error("Output of testWriter should be HELLO")
	}
}
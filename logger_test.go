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
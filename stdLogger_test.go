package logger
import (
	"testing"
	"time"
	"fmt"
)

type testWriter struct{
	output string
}

func (self *testWriter) Write(b []byte) (n int, err error){
	self.output = string(b)
	return len(b), nil
}

var TestWriter = testWriter{}

func TestStdDebug(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.Debug("Something append")

	if TestWriter.output != fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"){
		t.Errorf("StdLogger.Debug(\"Something append\") should log %s instead of %s",
			fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"),
			TestWriter.output)
	}

	testLogger.Debug("Something append:", "Test")
	if TestWriter.output != fmt.Sprintf("[%s]DEBUG : %s%s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"){
		t.Errorf("StdLogger.Debug(\"Something append\", new(error)) should log %s instead of %s",
			fmt.Sprintf("[%s]DEBUG : %s %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"),
			TestWriter.output)
	}
}

func TestStdDebugF(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.DebugF("Something append:%s", "Test")

	if TestWriter.output != fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")){
		t.Errorf("StdLogger.Debugf should log %s instead of %s",
			fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")),
			TestWriter.output)
	}

	testLogger.DebugF("Something append:%s %d", "Test", 35)
	if TestWriter.output != fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)){
		t.Errorf("StdLogger.Debugf should log %s instead of %s",
			fmt.Sprintf("[%s]DEBUG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)),
			TestWriter.output)
	}
}

func TestStdLog(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.Log("Something append")

	if TestWriter.output != fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"){
		t.Errorf("StdLogger.Log(\"Something append\") should log %s instead of %s",
			fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"),
			TestWriter.output)
	}

	testLogger.Log("Something append:", "Test")
	if TestWriter.output != fmt.Sprintf("[%s]LOG : %s%s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"){
		t.Errorf("StdLogger.Log(\"Something append\", new(error)) should log %s instead of %s",
			fmt.Sprintf("[%s]LOG : %s %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"),
			TestWriter.output)
	}
}

func TestStdLogF(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.LogF("Something append:%s", "Test")

	if TestWriter.output != fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")){
		t.Errorf("StdLogger.Logf should log %s instead of %s",
			fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")),
			TestWriter.output)
	}

	testLogger.LogF("Something append:%s %d", "Test", 35)
	if TestWriter.output != fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)){
		t.Errorf("StdLogger.Logf should log %s instead of %s",
			fmt.Sprintf("[%s]LOG : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)),
			TestWriter.output)
	}
}

func TestStdError(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.Error("Something append")

	if TestWriter.output != fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"){
		t.Errorf("StdLogger.Error(\"Something append\") should log %s instead of %s",
			fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"),
			TestWriter.output)
	}

	testLogger.Error("Something append:", "Test")
	if TestWriter.output != fmt.Sprintf("[%s]ERROR : %s%s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"){
		t.Errorf("StdLogger.Error(\"Something append\", new(error)) should log %s instead of %s",
			fmt.Sprintf("[%s]ERROR : %s %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"),
			TestWriter.output)
	}
}

func TestStdErrorF(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.ErrorF("Something append:%s", "Test")

	if TestWriter.output != fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")){
		t.Errorf("StdLogger.Errorf should log %s instead of %s",
			fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")),
			TestWriter.output)
	}

	testLogger.ErrorF("Something append:%s %d", "Test", 35)
	if TestWriter.output != fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)){
		t.Errorf("StdLogger.Errorf should log %s instead of %s",
			fmt.Sprintf("[%s]ERROR : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)),
			TestWriter.output)
	}
}

func TestStdFatal(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.Fatal("Something append")

	if TestWriter.output != fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"){
		t.Errorf("StdLogger.Fatal(\"Something append\") should log %s instead of %s",
			fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append"),
			TestWriter.output)
	}

	testLogger.Fatal("Something append:", "Test")
	if TestWriter.output != fmt.Sprintf("[%s]FATAL : %s%s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"){
		t.Errorf("StdLogger.Fatal(\"Something append\", new(error)) should log %s instead of %s",
			fmt.Sprintf("[%s]FATAL : %s %s",time.Now().Format("Mon _2 Jan 2006 15:04"), "Something append:", "Test"),
			TestWriter.output)
	}
}

func TestStdFatalF(t *testing.T) {
	var testLogger = newStdLogger(&TestWriter)
	testLogger.FatalF("Something append:%s", "Test")

	if TestWriter.output != fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")){
		t.Errorf("StdLogger.Fatalf should log %s instead of %s",
			fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s", "Test")),
			TestWriter.output)
	}

	testLogger.FatalF("Something append:%s %d", "Test", 35)
	if TestWriter.output != fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)){
		t.Errorf("StdLogger.Fatalf should log %s instead of %s",
			fmt.Sprintf("[%s]FATAL : %s",time.Now().Format("Mon _2 Jan 2006 15:04"), fmt.Sprintf("Something append:%s %d", "Test", 35)),
			TestWriter.output)
	}
}

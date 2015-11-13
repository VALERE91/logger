package logger

import (
	"testing"

	"gopkg.in/mgo.v2"
)

func TestMongoDebug(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")
	collection.DropCollection()

	var l logEntry
	MongoDBLogger.Debug("Something append")

	if collection.Find(nil).One(&l); l.Message != "Something append" || l.Level != "DEBUG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Debug should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.Debug("Something append:", "Test")

	if collection.Find(nil).One(&l); l.Message != "Something append:Test" || l.Level != "DEBUG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Debug should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoDebugF(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")

	collection.DropCollection()

	var l logEntry
	MongoDBLogger.DebugF("Test log : %s", "Something append")

	if collection.Find(nil).One(&l); l.Message != "Test log : Something append" || l.Level != "DEBUG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.DebugF should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.DebugF("Test log : %s %d", "Test", 35)

	if collection.Find(nil).One(&l); l.Message != "Test log : Test 35" || l.Level != "DEBUG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.DebugF should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoLog(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")
	collection.DropCollection()

	var l logEntry
	MongoDBLogger.Log("Something append")

	if collection.Find(nil).One(&l); l.Message != "Something append" || l.Level != "LOG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Log should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.Log("Something append:", "Test")

	if collection.Find(nil).One(&l); l.Message != "Something append:Test" || l.Level != "LOG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Log should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoLogF(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")

	collection.DropCollection()

	var l logEntry
	MongoDBLogger.LogF("Test log : %s", "Something append")

	if collection.Find(nil).One(&l); l.Message != "Test log : Something append" || l.Level != "LOG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.LogF should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.LogF("Test log : %s %d", "Test", 35)

	if collection.Find(nil).One(&l); l.Message != "Test log : Test 35" || l.Level != "LOG" || l.Date.IsZero() {
		t.Error("MongoDBLogger.LogF should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoError(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")
	collection.DropCollection()

	var l logEntry
	MongoDBLogger.Error("Something append")

	if collection.Find(nil).One(&l); l.Message != "Something append" || l.Level != "ERROR" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Error should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.Error("Something append:", "Test")

	if collection.Find(nil).One(&l); l.Message != "Something append:Test" || l.Level != "ERROR" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Error should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoErrorF(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")

	collection.DropCollection()

	var l logEntry
	MongoDBLogger.ErrorF("Test log : %s", "Something append")

	if collection.Find(nil).One(&l); l.Message != "Test log : Something append" || l.Level != "ERROR" || l.Date.IsZero() {
		t.Error("MongoDBLogger.ErrorF should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.ErrorF("Test log : %s %d", "Test", 35)

	if collection.Find(nil).One(&l); l.Message != "Test log : Test 35" || l.Level != "ERROR" || l.Date.IsZero() {
		t.Error("MongoDBLogger.ErrorF should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoFatal(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")
	collection.DropCollection()

	var l logEntry
	MongoDBLogger.Fatal("Something append")

	if collection.Find(nil).One(&l); l.Message != "Something append" || l.Level != "FATAL" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Fatal should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.Fatal("Something append:", "Test")

	if collection.Find(nil).One(&l); l.Message != "Something append:Test" || l.Level != "FATAL" || l.Date.IsZero() {
		t.Error("MongoDBLogger.Fatal should not respond that. It's either a malformed message or a wrong log level")
	}
}

func TestMongoFatalF(t *testing.T) {
	ConfigureMongoLogger("localhost", "test", "logs")
	session, err := mgo.Dial("localhost")

	if err != nil {
		t.Error("Unable to connect to MongoDB")
	}

	defer session.Close()

	collection := session.DB("test").C("logs")

	collection.DropCollection()

	var l logEntry
	MongoDBLogger.FatalF("Test log : %s", "Something append")

	if collection.Find(nil).One(&l); l.Message != "Test log : Something append" || l.Level != "FATAL" || l.Date.IsZero() {
		t.Error("MongoDBLogger.FatalF should not respond that. It's either a malformed message or a wrong log level")
	}

	collection.DropCollection()
	MongoDBLogger.FatalF("Test log : %s %d", "Test", 35)

	if collection.Find(nil).One(&l); l.Message != "Test log : Test 35" || l.Level != "FATAL" || l.Date.IsZero() {
		t.Error("MongoDBLogger.FatalF should not respond that. It's either a malformed message or a wrong log level")
	}
}

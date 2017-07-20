package main

import (
	"fmt"

	"github.com/docktermj/go-logrus/islog"
	log "github.com/sirupsen/logrus"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

type MyFormatter struct {}



func (f *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
    
    data := entry.Data
    level := entry.Level
    message := entry.Message
    time := entry.Time
    
    logRecord := fmt.Sprintf("%s, [%s] %s DATA: %+v\n", time, level, message, data)

  return []byte(logRecord), nil
}


// Test IsXXX() guards.
func testGuards() {

	if islog.Panic() {
		fmt.Println("IsPanic()")
	}

	if islog.Fatal() {
		fmt.Println("IsFatal()")
	}

	if islog.Error() {
		fmt.Println("IsError()")
	}

	if islog.Warning() {
		fmt.Println("IsWarning()")
	}

	if islog.Info() {
		fmt.Println("IsInfo()")
	}

	if islog.Debug() {
		fmt.Println("IsDebug()")
	}
}

func printLogs() {

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Debug("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Error("The group's number increased tremendously!")

	//  This will cause the program to exit.
	//	log.WithFields(log.Fields{
	//		"omg":    true,
	//		"number": 100,
	//	}).Fatal("The ice breaks!")
}

// Reusing a logger.
func testReusableLogger() {

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func main() {

	fmt.Printf("\n------ Test 1 --------------------------------------\n\n")

	log.SetFormatter(&log.JSONFormatter{})
	fmt.Printf("Log Level: %s Formatter: JSON\n", log.GetLevel().String())

	testGuards()
	printLogs()

	fmt.Printf("\n------ Test 2 --------------------------------------\n\n")

	log.SetFormatter(&log.TextFormatter{})
	fmt.Printf("Log Level: %s Formatter: Text\n", log.GetLevel().String())

	testGuards()
	printLogs()

	fmt.Printf("\n------ Test 3 --------------------------------------\n\n")
	
    log.SetFormatter(new(MyFormatter))
	fmt.Printf("Log Level: %s Formatter: MyFormatter\n", log.GetLevel().String())

	testGuards()
	printLogs()
	

	fmt.Printf("\n------ Test 3 --------------------------------------\n\n")

	stdLog()

}

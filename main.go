package main

import "fmt"

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func main() {
	fmt.Printf("Hello, world! from %s version %s-%s\n", programName, buildVersion, buildIteration)
}

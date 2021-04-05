package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	// InfoLogger logging info
	InfoLogger *log.Logger

	// ErrorLogger logging errors
	ErrorLogger *log.Logger

	// FatalLogger logging fatal
	FatalLogger *log.Logger

	err error
)

func init() {
	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	FatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	err = nil
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		FatalLogger.Fatalln("wrong number of arguments, exactly 1 is needed")
	}

	// go to correct dir
	err = os.Chdir(`./` + argsWithoutProg[0])
	if err != nil {
		FatalLogger.Fatalf("failed to change directory to %s", argsWithoutProg[0])
	}

	// format code
	runFormat()

	// lint code
	runLinting()

	// run test
	runTests()

	// InfoLogger.Printf("[SUCCESS]")
}

func runFormat() {
	runCmd([]string{"/usr/local/go/bin/gofmt"}, "failed to format the code")
}

func runLinting() {
	runCmd([]string{"/usr/local/bin/golint"}, "failed to run golint")
}

func runTests() {
	runCmd([]string{"/usr/local/go/bin/go", "test"}, "tests failed")
}

func runCmd(toExecute []string, errorMsg string) {
	cmd := exec.Command(toExecute[0])

	if len(toExecute) > 1 {
		cmd = exec.Command(toExecute[0], toExecute[1:len(toExecute)]...)
	}

	cmd.Stdout = InfoLogger.Writer()
	cmd.Stderr = ErrorLogger.Writer()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		FatalLogger.Fatalf(errorMsg)
	}
}

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
	if runFormat() != nil {
		panic(-1)
	}

	// lint code
	if runLinting() != nil {
		panic(-1)
	}

	// run test
	if runTests() != nil {
		panic(-1)
	}

}

func runFormat() error {
	return runCmd([]string{"go", "fmt"}, "failed to format the code")
}

func runLinting() error {
	return runCmd([]string{"golint"}, "failed to run golint")
}

func runTests() error {
	return runCmd([]string{"go", "test"}, "tests failed")
}

func runCmd(toExecute []string, errorMsg string) error {
	cmd := exec.Command(toExecute[0])

	if len(toExecute) > 1 {
		cmd = exec.Command(toExecute[0], toExecute[1:]...)
	}

	cmd.Stdout = InfoLogger.Writer()
	cmd.Stderr = ErrorLogger.Writer()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		FatalLogger.Fatalf(errorMsg)
		return err
	}

	return nil
}

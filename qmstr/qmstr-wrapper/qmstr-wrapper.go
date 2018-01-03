package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	analyze "qmstr-prototype/qmstr/qmstr-analyze"
	util "qmstr-prototype/qmstr/qmstr-util"
	"syscall"
)

const debugEnv string = "QMSTR_DEBUG"

var (
	// Log is the default logger.
	logger *log.Logger
	debug  bool
)

func init() {
	// setup logging
	var infoWriter io.Writer
	debugMode := os.Getenv(debugEnv)
	switch debugMode {
	case "stdout":
		infoWriter = os.Stdout
	case "remote":
		infoWriter = util.NewHTTPRemoteLogger("localhost", 9000, "log")
	default:
		infoWriter = ioutil.Discard
	}
	logger = log.New(infoWriter, "", log.Ldate|log.Ltime)
	if debugMode == "stdout" {
		logger.Print("Debug output on stdout enabled. This might break your build!")
	}
}

func main() {
	commandLine := os.Args
	logger.Printf("QMSTR called via %v", commandLine)
	//extract the compiler
	prog := filepath.Base(commandLine[0])

	if prog == "qmstr-wrapper" {
		log.Fatal("This is not how you should invoke the qmstr-wrapper.\n\tSee https://github.com/QMSTR/qmstr-prototype for more information on how to use the QMSTR.")
	}

	//extract the rest of the arguments
	commandLineArgs := commandLine[1:]
	// run actual compiler
	actualProg, err := findProg(prog)
	checkErr(err)
	cmd := exec.Command(actualProg, commandLineArgs...)
	var stdoutbuf, stderrbuf bytes.Buffer
	cmd.Stdout = &stdoutbuf
	cmd.Stderr = &stderrbuf

	err = cmd.Run()
	// preserve non-zero return code
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				// preserve stderr
				if stderr := stderrbuf.String(); len(stderr) > 0 {
					fmt.Fprintf(os.Stderr, "%s", stderr)
				}
				os.Exit(status.ExitStatus())
			}
		} else {
			log.Fatalf("Compiler failed with %v", err)
		}
	}

	// preserve stdout
	if stdout := stdoutbuf.String(); len(stdout) > 0 {
		fmt.Fprintf(os.Stdout, "%s", stdout)
	}

	// detect analyzer and start analysis
	cA := getAnalyzer(prog, commandLineArgs, debug)
	cA.Analyze(false)
	cA.Print()
	cA.SendResults()
}

func findProg(prog string) (string, error) {
	path := os.Getenv("PATH")
	foundUs := false
	for _, dir := range filepath.SplitList(path) {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		path := filepath.Join(dir, prog)
		if err := findExecutable(path); err == nil {
			if foundUs {
				return path, nil
			}
			foundUs = true
		}
	}
	return "", errors.New(fmt.Sprintf("executable file %s not found in $PATH", prog))
}

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}

//return a more generic type
func getAnalyzer(program string, args []string, debug bool) *analyze.GNUCAnalyzer {
	switch program {
	case "g++", "gcc":
		return analyze.NewGNUCAnalyzer(args, debug)
	default:
		log.Fatal("Compiler not supported")
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

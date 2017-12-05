package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	analyze "qmstr-prototype/qmstr/qmstr-analyze"
	"strings"
	"syscall"
)

func main() {
	commandLine := os.Args
	//extract the compiler
	prog := commandLine[0]

	if strings.HasSuffix(prog, "qmstr-wrapper") {
		log.Fatal("This is not how you should invoke the qmstr-wrapper.\n\tSee https://github.com/endocode/qmstr-prototype for more information on how to use the QMSTR.")
	}

	if len(commandLine) < 2 {
		log.Fatal("Too few arguments")
	}
	//extract the rest of the arguments
	commandLineArgs := commandLine[1:]

	// run actual compiler
	actualProg, err := findProg(prog)
	checkErr(err)
	log.Printf("Found %s at %s\n", prog, actualProg)
	cmd := exec.Command(actualProg, commandLineArgs...)
	err = cmd.Start()
	checkErr(err)

	// detect analyzer and start analysis
	cA := getAnalyzer(prog, commandLineArgs)
	cA.Analyze(false)

	// join actual compiler and preserve non-zero return code
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
				os.Exit(status.ExitStatus())
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
	}

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
	return "", errors.New("executable file not found in $PATH")
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
func getAnalyzer(program string, args []string) *analyze.GNUCAnalyzer {
	switch program {
	case "g++", "gcc":
		return analyze.NewGNUCAnalyzer(args)
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

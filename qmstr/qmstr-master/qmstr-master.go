package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var (
	// Info is the logger for INFO level output.
	Info *log.Logger
	// Log is the default logger.
	Log *log.Logger
)

func main() {
	var verbose bool
	var printVersion bool
	pflag.BoolVarP(&verbose, "verbose", "v", false, "enable verbose log output")
	pflag.BoolVarP(&printVersion, "version", "V", false, "print version and exit")
	pflag.Parse()

	var infoWriter io.Writer
	if verbose {
		infoWriter = os.Stdout
	} else {
		infoWriter = ioutil.Discard
	}
	Info = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime)
	Log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Info.Printf("quartermaster master process starting")
	defer Info.Printf("quartermaster master process exiting")
	if printVersion {
		fmt.Println("Quartermaster master 0.0.1")
		return
	}
}

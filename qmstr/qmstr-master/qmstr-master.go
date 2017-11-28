package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	model "qmstr-prototype/qmstr/qmstr-model"

	"github.com/spf13/pflag"
)

var (
	// Info is the logger for INFO level output.
	Info *log.Logger
	// Log is the default logger.
	Log *log.Logger
	// Model is the data model managed by the master process.
	Model *model.DataModel
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
		// --version prints the version and then exits:
		fmt.Println("Quartermaster master 0.0.1")
		return
	}
	// default: run the master server until a quit requests comes in
	Model = model.NewModel()
	// TODO: also react to a SIGTERM/SIGKILL
	Log.Printf(<-startHTTPServer())
}

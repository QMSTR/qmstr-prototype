package compiler_analyzer

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	util "qmstr-prototype/qmstr/qmstr-util"
)

var (
	// Logger is the default logger.
	Logger *log.Logger
)

func initLogging(debugMode string) {
	// setup logging
	var infoWriter io.Writer
	switch debugMode {
	case "stdout":
		infoWriter = os.Stdout
	case "remote":
		infoWriter = util.NewHTTPRemoteLogger("localhost", 9000, "log")
	default:
		infoWriter = ioutil.Discard
	}
	Logger = log.New(infoWriter, "", log.Ldate|log.Ltime)
	if debugMode == "stdout" {
		Logger.Print("Debug output on stdout enabled. This might break your build!")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildCleanPath(base string, subpath string) string {
	if filepath.IsAbs(subpath) {
		return filepath.Clean(subpath)
	}

	if !filepath.IsAbs(base) {
		var err error
		base, err = filepath.Abs(base)
		checkErr(err)
	}
	tmpPath := filepath.Join(base, subpath)
	return filepath.Clean(tmpPath)
}

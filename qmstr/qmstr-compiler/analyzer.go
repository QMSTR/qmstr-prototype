package compiler_analyzer

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	// Logger is the default logger.
	Logger *log.Logger
)

func initLogging(debug bool) {
	var infoWriter io.Writer
	if debug {
		infoWriter = os.Stdout
	} else {
		infoWriter = ioutil.Discard
	}
	Logger = log.New(infoWriter, "", log.Ldate|log.Ltime)
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

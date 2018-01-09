package compiler_analyzer

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func analyzeSourceFile(sourcefile string) []string {
	licenses := []string{}
	cmd := exec.Command("ninka", "-i", sourcefile)
	err := cmd.Start()
	checkErr(err)
	if err := cmd.Wait(); err != nil {
		log.Fatalf("License analysis failed for %s", sourcefile)
	}

	licenseFile, err := os.Open(sourcefile + ".license")
	checkErr(err)
	r := csv.NewReader(licenseFile)
	r.Comma = ';'
	records, err := r.ReadAll()
	checkErr(err)

	for _, fields := range records {
		if len(fields) > 0 {
			licenses = append(licenses, fields[0])
		}
	}
	return licenses
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

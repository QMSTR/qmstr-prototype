package analyze

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
)

var (
	// Info is the logger for INFO level output.
	Info *log.Logger
	// Log is the default logger.
	Log *log.Logger
)

func init() {
	Log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
}

func AnalyzeSourceFile(sourcefile string) []string {
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

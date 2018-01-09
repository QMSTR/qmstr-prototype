package analysis

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
)

type NinkaAnalyzer struct {
	cmd     string
	cmdargs []string
	result  map[string]interface{}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewNinkjaAnalyzer() *NinkaAnalyzer {
	na := NinkaAnalyzer{"ninka", []string{"-i"}, map[string]interface{}{}}
	return &na
}

func (na *NinkaAnalyzer) analyzeSourceFile(sourcefile string) error {
	log.Printf("scanning %s", sourcefile)
	licenses := []string{}
	cmd := exec.Command(na.cmd, append(na.cmdargs, sourcefile)...)
	err := cmd.Start()
	checkErr(err)
	if err := cmd.Wait(); err != nil {
		log.Fatalf("License analysis failed for %s", sourcefile)
	}

	licenseFile, err := os.Open(sourcefile + ".license")
	if err != nil {
		return err
	}
	r := csv.NewReader(licenseFile)
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, fields := range records {
		if len(fields) > 0 {
			licenses = append(licenses, fields[0])
		}
	}

	na.result["licenses"] = licenses
	log.Printf("Found the following licenses: %v", licenses)
	return nil
}

func (na *NinkaAnalyzer) GetName() string {
	return "Ninka analyzer"
}

func (na *NinkaAnalyzer) Analyze(a Analyzable) error {
	err := na.analyzeSourceFile(a.GetFile())
	if err != nil {
		return err
	}

	a.StoreResult(na.result)

	return nil
}

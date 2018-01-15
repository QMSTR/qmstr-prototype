package analysis

import "fmt"

type PrescanScanCodeAnalyzer struct {
	ScanData map[string]interface{}
}

func (la *PrescanScanCodeAnalyzer) Configure(data map[string]interface{}) error {
	la.ScanData = data["scancode"].(map[string]interface{})
	return nil
}

func (la *PrescanScanCodeAnalyzer) Analyze(a Analyzable) error {
	filename := a.GetFile()
	licenses := []string{}
	for _, file := range la.ScanData["files"].([]interface{}) {
		fileData := file.(map[string]interface{})
		if fileData["path"] == filename {
			fmt.Printf("Found %s", filename)
			for _, license := range fileData["licenses"].([]interface{}) {
				licenses = append(licenses, license.(map[string]interface{})["spdx_license_key"].(string))
			}
		}
	}
	result := map[string]interface{}{}
	result["licenses"] = licenses
	a.StoreResult(result)
	return nil
}

func (la *PrescanScanCodeAnalyzer) GetName() string {
	return "PrescanScanCode Analyzer"
}

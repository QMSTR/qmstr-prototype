package main

import (
	"bytes"
	"log"
	model "qmstr-prototype/qmstr/qmstr-model"
	"strings"
	"text/template"
)

type report struct {
	SPDXVersion, DataLicense, Name string
	License                        string
}

// CreateReport renders an SPDX document for the given TargetEntity
func CreateReport(target model.TargetEntity) string {

	//Define the template
	const reportTemplate = `
SPDXVersion: {{.SPDXVersion}}
DataLicense: {{.DataLicense}}
PackageName:  {{.Name}}
PackageLicenseDeclared: {{.License}}
`
	//Create a new template and parse the data
	r := template.Must(template.New("report").Parse(reportTemplate))

	licenses := extractLicenses(target.Sources)
	report := report{"SPDX-2.0", "CCO-1.0", target.Name, strings.Join(licenses, " AND ")}

	//Execute the template
	b := bytes.Buffer{}
	err := r.Execute(&b, report)
	if err != nil {
		log.Println("Failed to render report template:", err)
	}
	return b.String()
}

func extractLicenses(sources []string) []string {
	licenseSet := map[string]struct{}{}
	for _, v := range sources {
		s, err := Model.GetSourceEntity(v)
		if err != nil {
			return []string{}
		}
		if s.Licenses == nil || len(s.Licenses) == 0 {
			t, err := Model.GetTargetEntity(v)
			if err != nil {
				return []string{}
			}
			for _, source := range t.Sources {
				ts, err := Model.GetSourceEntity(source)
				if err != nil {
					return []string{}
				}
				for _, license := range ts.Licenses {
					licenseSet[license] = struct{}{}
				}
			}
		} else {
			for _, license := range s.Licenses {
				licenseSet[license] = struct{}{}
			}
		}
	}
	license := []string{}
	for k := range licenseSet {
		license = append(license, k)
	}
	return license
}

package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
)

func ScanDir(workdir string, scanResult chan interface{}) {
	scanResult <- scanCode(workdir, runtime.NumCPU())
}

func scanCode(workdir string, jobs int) interface{} {
	cmdlineargs := []string{"--quiet", "--full-root"}
	if jobs > 1 {
		cmdlineargs = append(cmdlineargs, "--processes", fmt.Sprintf("%d", jobs))
	}
	cmd := exec.Command("scancode", append(cmdlineargs, workdir)...)
	fmt.Printf("Calling %s", cmd.Args)
	scanResult, err := cmd.Output()
	if err != nil {
		log.Printf("Scandir failed %s", err)
	}
	re := regexp.MustCompile("{.+")
	jsonScanResult := re.Find(scanResult)
	var scanData interface{}
	err = json.Unmarshal(jsonScanResult, &scanData)
	if err != nil {
		log.Printf("parsing scan data failed %s", err)
	}
	return scanData
}

package analyze

import (
	"log"
	"os"
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

type Report interface {
	Send()
}

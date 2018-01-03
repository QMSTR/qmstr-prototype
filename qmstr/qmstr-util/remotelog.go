package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LogMessage struct {
	Msg string
}

type HTTPRemoteLogWriter struct {
	host     string
	port     int
	endpoint string
}

func NewHTTPRemoteLogger(host string, port int, endpoint string) *HTTPRemoteLogWriter {
	return &HTTPRemoteLogWriter{host, port, endpoint}
}

func (rlw *HTTPRemoteLogWriter) Write(p []byte) (int, error) {
	url := fmt.Sprintf("http://%s:%d/%s", rlw.host, rlw.port, rlw.endpoint)
	b, err := json.Marshal(&LogMessage{string(p)})
	if err != nil {
		return 0, fmt.Errorf("Unabe to create log message payload for \"%s\"", string(p))
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return 0, fmt.Errorf("Unabe to create POST request for URL \"%s\"", url)
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, fmt.Errorf("Error performing POST request for URL \"%s\": %s", url, err.Error())
	}

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return 0, fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return 0, fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return 0, fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return len(p), nil
}

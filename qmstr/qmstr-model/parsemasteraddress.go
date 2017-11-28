package model

import (
	"fmt"
	"net"
	"net/url"
)

// ParseMasterAddress parses the server address passed as a string.
func ParseMasterAddress(address string) (string, string, string, error) {
	if len(address) == 0 {
		return "", "", "", fmt.Errorf("Quartermaster master address not specified")
	}

	url, err := url.Parse(address)
	if err != nil {
		return "", "", "", fmt.Errorf("error parsing master URL \"%s\"", address)
	}
	host, port, err := net.SplitHostPort(url.Host)
	if err != nil {
		return "", "", "", fmt.Errorf("error retrieving host and port from URL \"%s\"", address)
	}
	scheme := url.Scheme
	if len(scheme) == 0 {
		return "", "", "", fmt.Errorf("error parsing scheme in URL \"%s\"", address)
	}
	return scheme, host, port, nil
}

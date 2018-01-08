package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client implements the operations a QMSTR client program can initiate on the master server.
type Client struct {
	masterURL string
}

// NewClient creates a new Client object.
func NewClient(url string) *Client {
	c := new(Client)
	c.masterURL = url
	return c
}

func (c *Client) getMasterURL(path string) (*url.URL, error) {
	u, err := url.Parse(c.masterURL)
	if err != nil {
		return nil, fmt.Errorf("Malformed server URL \"%s\"", c.masterURL)
	}
	u.Path = path
	return u, nil
}

func (c *Client) getGetRequestWithID(u *url.URL, id string) (*http.Request, error) {
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Unabe to create GET request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", id)
	request.URL.RawQuery = q.Encode()
	return request, nil
}

func (c *Client) performGetRequestWithID(path string, id string, value interface{}) ([]byte, error) {
	u, err := c.getMasterURL(path)
	request, err := c.getGetRequestWithID(u, id)
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error performing GET request for path \"%s\" and id \"%s\" - server not available?", path, id)
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
	// fmt.Println("Done")

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return nil, fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return nil, fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return nil, fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading source entity response body")
	}
	// fmt.Printf("performGetRequestWithID: response: %v", string(body[:]))
	return body, nil
}

// GetSourceEntity retrieves a source entity from the master server.
// err indicates a communication error. If the entity did not exist, an empty object will be returned.
func (c *Client) GetSourceEntity(id string) (SourceEntity, error) {
	var result SourceEntity
	body, err := c.performGetRequestWithID("sources", id, result)
	if err != nil {
		return SourceEntity{"", "", []string{}}, fmt.Errorf("error retrieving source entity \"%s\"", id)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SourceEntity{"", "", []string{}}, fmt.Errorf("error parsing source entity response")
	}
	return result, nil
}

// AddSourceEntity adds a source entity to the master data model.
func (c *Client) AddSourceEntity(s SourceEntity) error {
	u, err := c.getMasterURL("sources")
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("unable to encode source entity \"%s\" into JSON format", s.ID())
	}
	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create POST request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", s.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing POST request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// ModifySourceEntity modifies, guess what, a source entity.
func (c *Client) ModifySourceEntity(s SourceEntity) error {
	u, err := c.getMasterURL("sources")
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("unable to encode source entity \"%s\" into JSON format", s.ID())
	}
	request, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create PUT request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", s.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing PUT request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// DeleteSourceEntity deletes a source entity from the master server.
func (c *Client) DeleteSourceEntity(s SourceEntity) error {
	u, err := c.getMasterURL("sources")

	request, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return fmt.Errorf("Unabe to create DELETE request for URL \"%s\"", u.String())
	}
	// request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", s.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error performing DELETE request for entity \"%s\" - server not available?", s.ID())
	}
	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading source entity response body")
	}
	return nil
}

// GetDependencyEntity retrieves a dependency entity from the master server.
// err indicates a communication error. If the entity did not exist, an empty object will be returned.
func (c *Client) GetDependencyEntity(id string) (DependencyEntity, error) {
	var result DependencyEntity
	body, err := c.performGetRequestWithID("dependencies", id, result)
	if err != nil {
		return DependencyEntity{"", ""}, fmt.Errorf("error retrieving dependency entity \"%s\"", id)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return DependencyEntity{"", ""}, fmt.Errorf("error parsing dependency entity response")
	}
	return result, nil
}

// AddDependencyEntity adds a dependency entity to the master data model.
func (c *Client) AddDependencyEntity(d DependencyEntity) error {
	u, err := c.getMasterURL("dependencies")
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to encode dependency entity \"%s\" into JSON format", d.ID())
	}
	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create POST request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", d.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing POST request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// ModifyDependencyEntity modifies, guess what, a dependency entity.
func (c *Client) ModifyDependencyEntity(d DependencyEntity) error {
	u, err := c.getMasterURL("dependencies")
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to encode dependency entity \"%s\" into JSON format", d.ID())
	}
	request, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create PUT request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", d.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing PUT request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// DeleteDependencyEntity deletes a dependency entity from the master server.
func (c *Client) DeleteDependencyEntity(d DependencyEntity) error {
	u, err := c.getMasterURL("dependencies")

	request, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return fmt.Errorf("Unabe to create DELETE request for URL \"%s\"", u.String())
	}
	// request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", d.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error performing DELETE request for entity \"%s\" - server not available?", d.ID())
	}
	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading source entity response body")
	}
	return nil
}

// GetTargetEntity retrieves a target entity from the master server.
// err indicates a communication error. If the entity did not exist, an empty object will be returned.
func (c *Client) GetTargetEntity(id string) (TargetEntity, error) {
	var result TargetEntity
	body, err := c.performGetRequestWithID("targets", id, result)
	if err != nil {
		return TargetEntity{Name: "", Hash: ""}, fmt.Errorf("error retrieving entity \"%s\"", id)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return TargetEntity{"", "", []string{}, []string{}, false, ""}, fmt.Errorf("error parsing entity response")
	}
	return result, nil
}

// AddTargetEntity adds a target entity to the master data model.
func (c *Client) AddTargetEntity(d TargetEntity) error {
	u, err := c.getMasterURL("targets")
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to encode entity \"%s\" into JSON format", d.ID())
	}
	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create POST request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", d.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing POST request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// ModifyTargetEntity modifies, guess what, a target entity.
func (c *Client) ModifyTargetEntity(e TargetEntity) error {
	u, err := c.getMasterURL("targets")
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("unable to encode entity \"%s\" into JSON format", e.ID())
	}
	request, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Unabe to create PUT request for URL \"%s\"", u.String())
	}
	request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", e.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error performing PUT request for URL \"%s\": %s", u.String(), err.Error())
	}

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	return nil
}

// DeleteTargetEntity deletes a dependency entity from the master server.
func (c *Client) DeleteTargetEntity(e TargetEntity) error {
	u, err := c.getMasterURL("targets")

	request, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return fmt.Errorf("Unabe to create DELETE request for URL \"%s\"", u.String())
	}
	// request.Header.Set("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("id", e.ID())
	request.URL.RawQuery = q.Encode()
	query := request.URL.String()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error performing DELETE request for entity \"%s\" - server not available?", e.ID())
	}
	switch {
	case response.StatusCode == http.StatusNotImplemented:
		return fmt.Errorf("operation not implemented: \"%s\"", query)
	case response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons:
		// all bad :-(
		return fmt.Errorf("request failed for \"%s\": %d", query, response.StatusCode)
	case response.StatusCode != http.StatusOK:
		return fmt.Errorf("request did not succeed for \"%s\": %d", query, response.StatusCode)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading source entity response body")
	}
	return nil
}

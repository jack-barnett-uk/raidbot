package ffxivtoolkit

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Endpoint represents a set of endpoints on FFXIV Toolkit
type Endpoint struct {
	client   *Client
	endpoint string
}

func (e *Endpoint) makeRequest(httpMethod string, method string, v interface{}) {

	url := e.client.BaseURL + e.endpoint + method
	requestBody := bytes.NewBuffer(nil)

	if httpMethod == http.MethodPatch {

		vAsJSON, err := getObjectAsJSON(v)

		if err != nil {
			panic(err.Error())
		}

		requestBody = bytes.NewBuffer(vAsJSON)
	}

	req, err := http.NewRequest(httpMethod, url, requestBody)

	e.addTokenToRequest(req)

	if err != nil {
		panic(err.Error())
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		panic(response.StatusCode)

	}

	defer response.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &v)

	if err != nil {
		panic(err.Error())
	}
}

func (e Endpoint) addTokenToRequest(request *http.Request) {
	q := request.URL.Query()
	q.Add("token", e.client.Token)

	request.URL.RawQuery = q.Encode()
}

func getObjectAsJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

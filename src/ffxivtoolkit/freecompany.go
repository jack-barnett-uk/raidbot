package ffxivtoolkit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BasicFreeCompanyDetails details a Free Company from FFXIV Toolkit
type BasicFreeCompanyDetails struct {
	LodestoneID string `json:"lodestone_id"`
	Name        string `json:"name"`
}

// FreeCompanyMembers provides details of an FC member.
type FreeCompanyMembers []struct {
	LodestoneID string `json:"lodestone_id"`
	Name        string `json:"name"`
}

// FreeCompany provides methods for accessing the Free Company API
type FreeCompany struct {
	fcEndpointURL string
	client        *Client
}

func (fc *FreeCompany) makeRequest(endpoint string, v interface{}) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	fc.addTokenToRequest(req)

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

// Basic retrieves information about a provided FC lodestone id.
func (fc *FreeCompany) Basic(lodestoneID string) BasicFreeCompanyDetails {
	basic := BasicFreeCompanyDetails{}
	fc.makeRequest(fc.fcEndpointURL+lodestoneID, &basic)

	return basic
}

// Members retrieves a list of members for a provided FC lodestone id
func (fc *FreeCompany) Members(lodestoneID string) FreeCompanyMembers {
	arrayOfMembers := FreeCompanyMembers{}

	fc.makeRequest(fc.fcEndpointURL+lodestoneID+"/Members", &arrayOfMembers)

	return arrayOfMembers
}

func (fc FreeCompany) addTokenToRequest(request *http.Request) {
	q := request.URL.Query()
	q.Add("token", fc.client.Token)

	request.URL.RawQuery = q.Encode()
}

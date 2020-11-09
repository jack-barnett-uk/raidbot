package ffxivtoolkit

import "net/http"

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
	Endpoint
}

// Basic retrieves information about a provided FC lodestone id.
func (fc *FreeCompany) Basic(lodestoneID string) BasicFreeCompanyDetails {
	basic := BasicFreeCompanyDetails{}

	fc.makeRequest(http.MethodGet, lodestoneID, &basic)

	return basic
}

// Members retrieves a list of members for a provided FC lodestone id
func (fc *FreeCompany) Members(lodestoneID string) FreeCompanyMembers {
	arrayOfMembers := FreeCompanyMembers{}

	fc.makeRequest(http.MethodGet, lodestoneID+"/Members", &arrayOfMembers)

	return arrayOfMembers
}

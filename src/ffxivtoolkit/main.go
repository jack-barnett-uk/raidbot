package ffxivtoolkit

// Client is used to access several API endpoints for FFXIV Toolkit
type Client struct {
	baseURL     string
	Token       string
	FreeCompany FreeCompany
}

// NewWithBaseURL builds a new FFXIVToolkitClient with provided base URL
func NewWithBaseURL(token string, baseURL string) Client {
	newClient := Client{Token: token, baseURL: baseURL}
	newClient.FreeCompany = FreeCompany{client: &newClient, fcEndpointURL: newClient.baseURL + "FreeCompany/"}

	return newClient

}

// New builds a new FFXIVToolkitClient
func New(token string) Client {
	newClient := NewWithBaseURL(token, "https://ffxiv.ramirezrtg.app/API/")

	return newClient
}

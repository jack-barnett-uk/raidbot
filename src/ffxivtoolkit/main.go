package ffxivtoolkit

// Client is used to access several API endpoints for FFXIV Toolkit
type Client struct {
	BaseURL     string
	Token       string
	FreeCompany FreeCompany
	Group       Group
}

// NewWithBaseURL builds a new FFXIVToolkitClient with provided base URL
func NewWithBaseURL(token string, baseURL string) Client {
	newClient := Client{Token: token, BaseURL: baseURL}

	newClient.FreeCompany = FreeCompany{Endpoint: Endpoint{client: &newClient, endpoint: "FreeCompany/"}}
	newClient.Group = Group{Endpoint: Endpoint{client: &newClient, endpoint: "Group/"}}

	return newClient

}

// New builds a new FFXIVToolkitClient
func New(token string) Client {
	newClient := NewWithBaseURL(token, "https://ffxiv.ramirezrtg.app/API/")

	return newClient
}

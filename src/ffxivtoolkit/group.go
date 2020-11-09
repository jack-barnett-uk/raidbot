package ffxivtoolkit

import (
	"fmt"
	"net/http"
)

// Group represents the endpoints used to managing groups in ffxiv toolkit.
type Group struct {
	Endpoint
	Members Members
}

// GroupResponse represents the server response when creating a new group
type GroupResponse struct {
	Code    int    `json:"code"`
	Message string `json:"name"`
}

// GroupDetails represents a single group
type GroupDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create will return details on whether a group has/hasn't been created.
func (g *Group) Create(groupName string) GroupResponse {
	groupResponse := GroupResponse{}

	g.makeRequest(http.MethodPost, groupName, &groupResponse)

	return groupResponse
}

// Delete will delete a group with the provided name
func (g *Group) Delete(groupName string) GroupResponse {
	groupResponse := GroupResponse{}

	g.makeRequest(http.MethodDelete, groupName, &groupResponse)

	return groupResponse
}

// Get will find and return a group with the provided name
func (g *Group) Get(groupName string) (group GroupDetails, err error) {
	groupResponse := []GroupDetails{}

	g.makeRequest(http.MethodGet, groupName, &groupResponse)

	if len(groupResponse) == 1 {
		group = groupResponse[0]
	} else {
		err = fmt.Errorf("No groups found for %v", groupName)
	}

	return
}

// GetAll will return all groups the provided token has access to.
func (g *Group) GetAll() []GroupDetails {
	groups := []GroupDetails{}

	g.makeRequest(http.MethodGet, "", &groups)

	return groups
}

// Update will update the provided group on FFXIV Toolkit
func (g *Group) Update(group GroupDetails) GroupDetails {
	g.makeRequest(http.MethodPatch, group.Name, group)

	return group
}

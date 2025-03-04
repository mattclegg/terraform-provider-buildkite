// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package buildkite

import (
	"github.com/Khan/genqlient/graphql"
)

// The roles a user can be within a team
type GTeamMemberRole string

const (
	// The user is a regular member of the team
	GTeamMemberRoleMember GTeamMemberRole = "MEMBER"
	// The user can manage pipelines and users within the team
	GTeamMemberRoleMaintainer GTeamMemberRole = "MAINTAINER"
)

// Whether a team is visible or secret within an organization
type GenqlientTeamPrivacy string

const (
	// Visible to all members of the organization
	GenqlientTeamPrivacyVisible GenqlientTeamPrivacy = "VISIBLE"
	// Visible to organization administrators and members
	GenqlientTeamPrivacySecret GenqlientTeamPrivacy = "SECRET"
)

// __getOrganizationInput is used internally by genqlient
type __getOrganizationInput struct {
	Slug string `json:"slug"`
}

// GetSlug returns __getOrganizationInput.Slug, and is useful for accessing the field via an interface.
func (v *__getOrganizationInput) GetSlug() string { return v.Slug }

// __getTeamInput is used internally by genqlient
type __getTeamInput struct {
	Slug string `json:"slug"`
}

// GetSlug returns __getTeamInput.Slug, and is useful for accessing the field via an interface.
func (v *__getTeamInput) GetSlug() string { return v.Slug }

// __setApiIpAddressesInput is used internally by genqlient
type __setApiIpAddressesInput struct {
	OrganizationID string `json:"organizationID"`
	IpAddresses    string `json:"ipAddresses"`
}

// GetOrganizationID returns __setApiIpAddressesInput.OrganizationID, and is useful for accessing the field via an interface.
func (v *__setApiIpAddressesInput) GetOrganizationID() string { return v.OrganizationID }

// GetIpAddresses returns __setApiIpAddressesInput.IpAddresses, and is useful for accessing the field via an interface.
func (v *__setApiIpAddressesInput) GetIpAddresses() string { return v.IpAddresses }

// getOrganizationOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An organization
type getOrganizationOrganization struct {
	// A space-separated allowlist of IP addresses that can access the organization via the GraphQL or REST API
	AllowedApiIpAddresses string `json:"allowedApiIpAddresses"`
	Id                    string `json:"id"`
	// The public UUID for this organization
	Uuid string `json:"uuid"`
}

// GetAllowedApiIpAddresses returns getOrganizationOrganization.AllowedApiIpAddresses, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetAllowedApiIpAddresses() string {
	return v.AllowedApiIpAddresses
}

// GetId returns getOrganizationOrganization.Id, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetId() string { return v.Id }

// GetUuid returns getOrganizationOrganization.Uuid, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetUuid() string { return v.Uuid }

// getOrganizationResponse is returned by getOrganization on success.
type getOrganizationResponse struct {
	// Find an organization
	Organization getOrganizationOrganization `json:"organization"`
}

// GetOrganization returns getOrganizationResponse.Organization, and is useful for accessing the field via an interface.
func (v *getOrganizationResponse) GetOrganization() getOrganizationOrganization {
	return v.Organization
}

// getTeamResponse is returned by getTeam on success.
type getTeamResponse struct {
	// Find a team
	Team getTeamTeam `json:"team"`
}

// GetTeam returns getTeamResponse.Team, and is useful for accessing the field via an interface.
func (v *getTeamResponse) GetTeam() getTeamTeam { return v.Team }

// getTeamTeam includes the requested fields of the GraphQL type Team.
// The GraphQL type's documentation follows.
//
// An organization team
type getTeamTeam struct {
	// New organization members will be granted this role on this team
	DefaultMemberRole GTeamMemberRole `json:"defaultMemberRole"`
	// A description of the team
	Description string `json:"description"`
	Id          string `json:"id"`
	// Add new organization members to this team by default
	IsDefaultTeam bool `json:"isDefaultTeam"`
	// Whether or not team members can create new pipelines in this team
	MembersCanCreatePipelines bool `json:"membersCanCreatePipelines"`
	// The name of the team
	Name string `json:"name"`
	// The privacy setting for this team
	Privacy GenqlientTeamPrivacy `json:"privacy"`
	// The slug of the team
	Slug string `json:"slug"`
	// The public UUID for this team
	Uuid string `json:"uuid"`
}

// GetDefaultMemberRole returns getTeamTeam.DefaultMemberRole, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetDefaultMemberRole() GTeamMemberRole { return v.DefaultMemberRole }

// GetDescription returns getTeamTeam.Description, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetDescription() string { return v.Description }

// GetId returns getTeamTeam.Id, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetId() string { return v.Id }

// GetIsDefaultTeam returns getTeamTeam.IsDefaultTeam, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetIsDefaultTeam() bool { return v.IsDefaultTeam }

// GetMembersCanCreatePipelines returns getTeamTeam.MembersCanCreatePipelines, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetMembersCanCreatePipelines() bool { return v.MembersCanCreatePipelines }

// GetName returns getTeamTeam.Name, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetName() string { return v.Name }

// GetPrivacy returns getTeamTeam.Privacy, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetPrivacy() GenqlientTeamPrivacy { return v.Privacy }

// GetSlug returns getTeamTeam.Slug, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetSlug() string { return v.Slug }

// GetUuid returns getTeamTeam.Uuid, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetUuid() string { return v.Uuid }

// setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload includes the requested fields of the GraphQL type OrganizationAPIIPAllowlistUpdateMutationPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of OrganizationAPIIPAllowlistUpdateMutation.
type setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload struct {
	Organization setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization `json:"organization"`
}

// GetOrganization returns setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload.Organization, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload) GetOrganization() setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization {
	return v.Organization
}

// setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An organization
type setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization struct {
	// A space-separated allowlist of IP addresses that can access the organization via the GraphQL or REST API
	AllowedApiIpAddresses string `json:"allowedApiIpAddresses"`
}

// GetAllowedApiIpAddresses returns setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization.AllowedApiIpAddresses, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization) GetAllowedApiIpAddresses() string {
	return v.AllowedApiIpAddresses
}

// setApiIpAddressesResponse is returned by setApiIpAddresses on success.
type setApiIpAddressesResponse struct {
	// Sets an allowlist of IP addresses for API access to an organization. Please note that this is a beta feature and is not yet available to all organizations.
	OrganizationApiIpAllowlistUpdate setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload `json:"organizationApiIpAllowlistUpdate"`
}

// GetOrganizationApiIpAllowlistUpdate returns setApiIpAddressesResponse.OrganizationApiIpAllowlistUpdate, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesResponse) GetOrganizationApiIpAllowlistUpdate() setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload {
	return v.OrganizationApiIpAllowlistUpdate
}

func getOrganization(
	client graphql.Client,
	slug string,
) (*getOrganizationResponse, error) {
	__input := __getOrganizationInput{
		Slug: slug,
	}
	var err error

	var retval getOrganizationResponse
	err = client.MakeRequest(
		nil,
		"getOrganization",
		`
query getOrganization ($slug: ID!) {
	organization(slug: $slug) {
		allowedApiIpAddresses
		id
		uuid
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func getTeam(
	client graphql.Client,
	slug string,
) (*getTeamResponse, error) {
	__input := __getTeamInput{
		Slug: slug,
	}
	var err error

	var retval getTeamResponse
	err = client.MakeRequest(
		nil,
		"getTeam",
		`
query getTeam ($slug: ID!) {
	team(slug: $slug) {
		defaultMemberRole
		description
		id
		isDefaultTeam
		membersCanCreatePipelines
		name
		privacy
		slug
		uuid
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func setApiIpAddresses(
	client graphql.Client,
	organizationID string,
	ipAddresses string,
) (*setApiIpAddressesResponse, error) {
	__input := __setApiIpAddressesInput{
		OrganizationID: organizationID,
		IpAddresses:    ipAddresses,
	}
	var err error

	var retval setApiIpAddressesResponse
	err = client.MakeRequest(
		nil,
		"setApiIpAddresses",
		`
mutation setApiIpAddresses ($organizationID: ID!, $ipAddresses: String!) {
	organizationApiIpAllowlistUpdate(input: {organizationID:$organizationID,ipAddresses:$ipAddresses}) {
		organization {
			allowedApiIpAddresses
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

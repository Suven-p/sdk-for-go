package teams

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Teams service
type Teams struct {
	client appwrite.Client
}

func NewTeams(clt appwrite.Client) *Teams {
	return &Teams{
		client: clt,
	}
}

// List get a list of all the teams in which the current user is a member. You
// can use the parameters to filter your results.
type ListOptions struct {
	Queries []interface{}
	Search string
}

type ListOption func(*ListOptions)

func WithListQueries(v []interface{}) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
	}
}
func WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
	}
}

	
func (srv *Teams) List(optionalSetters ...ListOption)  (*appwriteModel.TeamList, error) {
	path := "/teams"

	options := &ListOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.TeamList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.TeamList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Create create a new team. The user who creates the team will automatically
// be assigned as the owner of the team. Only the users with the owner role
// can invite new members, add new owners and delete or update the team.
type CreateOptions struct {
	Roles []interface{}
}

type CreateOption func(*CreateOptions)

func WithCreateRoles(v []interface{}) CreateOption {
	return func(o *CreateOptions) {
		o.Roles = v
	}
}

					
func (srv *Teams) Create(TeamId string, Name string, optionalSetters ...CreateOption)  (*appwriteModel.Team, error) {
	path := "/teams"

	options := &CreateOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["name"] = Name
	params["roles"] = options.Roles
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Team
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Get get a team by its ID. All team members have read access for this
// resource.
type GetOptions struct {
}

type GetOption func(*GetOptions)


	
func (srv *Teams) Get(TeamId string)  (*appwriteModel.Team, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Team
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateName update the team's name by its unique ID.
type UpdateNameOptions struct {
}

type UpdateNameOption func(*UpdateNameOptions)


			
func (srv *Teams) UpdateName(TeamId string, Name string)  (*appwriteModel.Team, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Team
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Delete delete a team using its ID. Only team members with the owner role
// can delete the team.
type DeleteOptions struct {
}

type DeleteOption func(*DeleteOptions)


	
func (srv *Teams) Delete(TeamId string)  (*interface{}, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListMemberships use this endpoint to list a team's members using the team's
// ID. All team members have read access to this endpoint.
type ListMembershipsOptions struct {
	Queries []interface{}
	Search string
}

type ListMembershipsOption func(*ListMembershipsOptions)

func WithListMembershipsQueries(v []interface{}) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Queries = v
	}
}
func WithListMembershipsSearch(v string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Search = v
	}
}

			
func (srv *Teams) ListMemberships(TeamId string, optionalSetters ...ListMembershipsOption)  (*appwriteModel.MembershipList, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")

	options := &ListMembershipsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.MembershipList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.MembershipList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateMembership invite a new member to join your team. Provide an ID for
// existing users, or invite unregistered users using an email or phone
// number. If initiated from a Client SDK, Appwrite will send an email or sms
// with a link to join the team to the invited user, and an account will be
// created for them if one doesn't exist. If initiated from a Server SDK, the
// new member will be added automatically to the team.
// 
// You only need to provide one of a user ID, email, or phone number. Appwrite
// will prioritize accepting the user ID > email > phone number if you provide
// more than one of these parameters.
// 
// Use the `url` parameter to redirect the user from the invitation email to
// your app. After the user is redirected, use the [Update Team Membership
// Status](/docs/client/teams#teamsUpdateMembershipStatus) endpoint to allow
// the user to accept the invitation to the team. 
// 
// Please note that to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// Appwrite will accept the only redirect URLs under the domains you have
// added as a platform on the Appwrite Console.
// 
type CreateMembershipOptions struct {
	Email string
	UserId string
	Phone string
	Name string
}

type CreateMembershipOption func(*CreateMembershipOptions)

func WithCreateMembershipEmail(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Email = v
	}
}
func WithCreateMembershipUserId(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.UserId = v
	}
}
func WithCreateMembershipPhone(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Phone = v
	}
}
func WithCreateMembershipName(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Name = v
	}
}

							
func (srv *Teams) CreateMembership(TeamId string, Roles []interface{}, Url string, optionalSetters ...CreateMembershipOption)  (*appwriteModel.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")

	options := &CreateMembershipOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["roles"] = Roles
	params["url"] = Url
	params["email"] = options.Email
	params["userId"] = options.UserId
	params["phone"] = options.Phone
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Membership
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetMembership get a team member by the membership unique id. All team
// members have read access for this resource.
type GetMembershipOptions struct {
}

type GetMembershipOption func(*GetMembershipOptions)


			
func (srv *Teams) GetMembership(TeamId string, MembershipId string)  (*appwriteModel.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Membership
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateMembershipRoles modify the roles of a team member. Only team members
// with the owner role have access to this endpoint. Learn more about [roles
// and permissions](/docs/permissions).
type UpdateMembershipRolesOptions struct {
}

type UpdateMembershipRolesOption func(*UpdateMembershipRolesOptions)


					
func (srv *Teams) UpdateMembershipRoles(TeamId string, MembershipId string, Roles []interface{})  (*appwriteModel.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	params["roles"] = Roles
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Membership
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteMembership this endpoint allows a user to leave a team or for a team
// owner to delete the membership of any other team member. You can also use
// this endpoint to delete a user membership even if it is not accepted.
type DeleteMembershipOptions struct {
}

type DeleteMembershipOption func(*DeleteMembershipOptions)


			
func (srv *Teams) DeleteMembership(TeamId string, MembershipId string)  (*interface{}, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateMembershipStatus use this endpoint to allow a user to accept an
// invitation to join a team after being redirected back to your app from the
// invitation email received by the user.
// 
// If the request is successful, a session for the user is automatically
// created.
// 
type UpdateMembershipStatusOptions struct {
}

type UpdateMembershipStatusOption func(*UpdateMembershipStatusOptions)


							
func (srv *Teams) UpdateMembershipStatus(TeamId string, MembershipId string, UserId string, Secret string)  (*appwriteModel.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}/status")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Membership
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetPrefs get the team's shared preferences by its unique ID. If a
// preference doesn't need to be shared by all team members, prefer storing
// them in [user preferences](/docs/client/account#accountGetPrefs).
type GetPrefsOptions struct {
}

type GetPrefsOption func(*GetPrefsOptions)


	
func (srv *Teams) GetPrefs(TeamId string)  (*appwriteModel.Preferences, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/prefs")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Preferences
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdatePrefs update the team's preferences by its unique ID. The object you
// pass is stored as is and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws an error if exceeded.
type UpdatePrefsOptions struct {
}

type UpdatePrefsOption func(*UpdatePrefsOptions)


			
func (srv *Teams) UpdatePrefs(TeamId string, Prefs interface{})  (*appwriteModel.Preferences, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/prefs")

	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Preferences
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

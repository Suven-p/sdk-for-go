package account

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Account service
type Account struct {
	client appwrite.Client
}

func NewAccount(clt appwrite.Client) *Account {
	return &Account{
		client: clt,
	}
}

// Get get currently logged in user data as JSON object.
type GetOptions struct {
}

type GetOption func(*GetOptions)



func (srv *Account) Get()  (*appwriteModel.User, error) {
	path := "/account"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateEmail update currently logged in user account email address. After
// changing user address, the user confirmation status will get reset. A new
// confirmation email is not sent automatically however you can use the send
// confirmation email endpoint again to send the confirmation email. For
// security measures, user password is required to complete this request.
// This endpoint can also be used to convert an anonymous account to a normal
// one, by passing an email address and a new password.
// 
type UpdateEmailOptions struct {
}

type UpdateEmailOption func(*UpdateEmailOptions)


			
func (srv *Account) UpdateEmail(Email string, Password string)  (*appwriteModel.User, error) {
	path := "/account/email"

	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListLogs get currently logged in user list of latest security activity
// logs. Each log returns user IP address, location and date and time of log.
type ListLogsOptions struct {
	Queries []interface{}
}

type ListLogsOption func(*ListLogsOptions)

func WithListLogsQueries(v []interface{}) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
	}
}

	
func (srv *Account) ListLogs(optionalSetters ...ListLogsOption)  (*appwriteModel.LogList, error) {
	path := "/account/logs"

	options := &ListLogsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["queries"] = options.Queries
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.LogList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateName update currently logged in user account name.
type UpdateNameOptions struct {
}

type UpdateNameOption func(*UpdateNameOptions)


	
func (srv *Account) UpdateName(Name string)  (*appwriteModel.User, error) {
	path := "/account/name"

	params := map[string]interface{}{}
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdatePassword update currently logged in user password. For validation,
// user is required to pass in the new password, and the old password. For
// users created with OAuth, Team Invites and Magic URL, oldPassword is
// optional.
type UpdatePasswordOptions struct {
	OldPassword string
}

type UpdatePasswordOption func(*UpdatePasswordOptions)

func WithUpdatePasswordOldPassword(v string) UpdatePasswordOption {
	return func(o *UpdatePasswordOptions) {
		o.OldPassword = v
	}
}

			
func (srv *Account) UpdatePassword(Password string, optionalSetters ...UpdatePasswordOption)  (*appwriteModel.User, error) {
	path := "/account/password"

	options := &UpdatePasswordOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["password"] = Password
	params["oldPassword"] = options.OldPassword
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdatePhone update the currently logged in user's phone number. After
// updating the phone number, the phone verification status will be reset. A
// confirmation SMS is not sent automatically, however you can use the [POST
// /account/verification/phone](/docs/client/account#accountCreatePhoneVerification)
// endpoint to send a confirmation SMS.
type UpdatePhoneOptions struct {
}

type UpdatePhoneOption func(*UpdatePhoneOptions)


			
func (srv *Account) UpdatePhone(Phone string, Password string)  (*appwriteModel.User, error) {
	path := "/account/phone"

	params := map[string]interface{}{}
	params["phone"] = Phone
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetPrefs get currently logged in user preferences as a key-value object.
type GetPrefsOptions struct {
}

type GetPrefsOption func(*GetPrefsOptions)



func (srv *Account) GetPrefs()  (*appwriteModel.Preferences, error) {
	path := "/account/prefs"

	params := map[string]interface{}{}
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
// UpdatePrefs update currently logged in user account preferences. The object
// you pass is stored as is, and replaces any previous value. The maximum
// allowed prefs size is 64kB and throws error if exceeded.
type UpdatePrefsOptions struct {
}

type UpdatePrefsOption func(*UpdatePrefsOptions)


	
func (srv *Account) UpdatePrefs(Prefs interface{})  (*appwriteModel.User, error) {
	path := "/account/prefs"

	params := map[string]interface{}{}
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateRecovery sends the user an email with a temporary secret key for
// password reset. When the user clicks the confirmation link he is redirected
// back to your app password reset URL with the secret key and email address
// values attached to the URL query string. Use the query string params to
// submit a request to the [PUT
// /account/recovery](/docs/client/account#accountUpdateRecovery) endpoint to
// complete the process. The verification link sent to the user's email
// address is valid for 1 hour.
type CreateRecoveryOptions struct {
}

type CreateRecoveryOption func(*CreateRecoveryOptions)


			
func (srv *Account) CreateRecovery(Email string, Url string)  (*appwriteModel.Token, error) {
	path := "/account/recovery"

	params := map[string]interface{}{}
	params["email"] = Email
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateRecovery use this endpoint to complete the user account password
// reset. Both the **userId** and **secret** arguments will be passed as query
// parameters to the redirect URL you have provided when sending your request
// to the [POST /account/recovery](/docs/client/account#accountCreateRecovery)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
type UpdateRecoveryOptions struct {
}

type UpdateRecoveryOption func(*UpdateRecoveryOptions)


							
func (srv *Account) UpdateRecovery(UserId string, Secret string, Password string, PasswordAgain string)  (*appwriteModel.Token, error) {
	path := "/account/recovery"

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	params["password"] = Password
	params["passwordAgain"] = PasswordAgain
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListSessions get currently logged in user list of active sessions across
// different devices.
type ListSessionsOptions struct {
}

type ListSessionsOption func(*ListSessionsOptions)



func (srv *Account) ListSessions()  (*appwriteModel.SessionList, error) {
	path := "/account/sessions"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.SessionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.SessionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteSessions delete all sessions from the user account and remove any
// sessions cookies from the end client.
type DeleteSessionsOptions struct {
}

type DeleteSessionsOption func(*DeleteSessionsOptions)



func (srv *Account) DeleteSessions()  (*interface{}, error) {
	path := "/account/sessions"

	params := map[string]interface{}{}
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
// GetSession use this endpoint to get a logged in user's session using a
// Session ID. Inputting 'current' will return the current session being used.
type GetSessionOptions struct {
}

type GetSessionOption func(*GetSessionOptions)


	
func (srv *Account) GetSession(SessionId string)  (*appwriteModel.Session, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{}
	params["sessionId"] = SessionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateSession access tokens have limited lifespan and expire to mitigate
// security risks. If session was created using an OAuth provider, this route
// can be used to "refresh" the access token.
type UpdateSessionOptions struct {
}

type UpdateSessionOption func(*UpdateSessionOptions)


	
func (srv *Account) UpdateSession(SessionId string)  (*appwriteModel.Session, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{}
	params["sessionId"] = SessionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteSession use this endpoint to log out the currently logged in user
// from all their account sessions across all of their different devices. When
// using the Session ID argument, only the unique session ID provided is
// deleted.
// 
type DeleteSessionOptions struct {
}

type DeleteSessionOption func(*DeleteSessionOptions)


	
func (srv *Account) DeleteSession(SessionId string)  (*interface{}, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{}
	params["sessionId"] = SessionId
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
// UpdateStatus block the currently logged in user account. Behind the scene,
// the user record is not deleted but permanently blocked from any access. To
// completely delete a user, use the Users API instead.
type UpdateStatusOptions struct {
}

type UpdateStatusOption func(*UpdateStatusOptions)



func (srv *Account) UpdateStatus()  (*appwriteModel.User, error) {
	path := "/account/status"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateVerification use this endpoint to send a verification message to your
// user email address to confirm they are the valid owners of that address.
// Both the **userId** and **secret** arguments will be passed as query
// parameters to the URL you have provided to be attached to the verification
// email. The provided URL should redirect the user back to your app and allow
// you to complete the verification process by verifying both the **userId**
// and **secret** parameters. Learn more about how to [complete the
// verification process](/docs/client/account#accountUpdateEmailVerification).
// The verification link sent to the user's email address is valid for 7 days.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md),
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
// 
type CreateVerificationOptions struct {
}

type CreateVerificationOption func(*CreateVerificationOptions)


	
func (srv *Account) CreateVerification(Url string)  (*appwriteModel.Token, error) {
	path := "/account/verification"

	params := map[string]interface{}{}
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateVerification use this endpoint to complete the user email
// verification process. Use both the **userId** and **secret** parameters
// that were attached to your app URL to verify the user email ownership. If
// confirmed this route will return a 200 status code.
type UpdateVerificationOptions struct {
}

type UpdateVerificationOption func(*UpdateVerificationOptions)


			
func (srv *Account) UpdateVerification(UserId string, Secret string)  (*appwriteModel.Token, error) {
	path := "/account/verification"

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreatePhoneVerification use this endpoint to send a verification SMS to the
// currently logged in user. This endpoint is meant for use after updating a
// user's phone number using the
// [accountUpdatePhone](/docs/client/account#accountUpdatePhone) endpoint.
// Learn more about how to [complete the verification
// process](/docs/client/account#accountUpdatePhoneVerification). The
// verification code sent to the user's phone number is valid for 15 minutes.
type CreatePhoneVerificationOptions struct {
}

type CreatePhoneVerificationOption func(*CreatePhoneVerificationOptions)



func (srv *Account) CreatePhoneVerification()  (*appwriteModel.Token, error) {
	path := "/account/verification/phone"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdatePhoneVerification use this endpoint to complete the user phone
// verification process. Use the **userId** and **secret** that were sent to
// your user's phone number to verify the user email ownership. If confirmed
// this route will return a 200 status code.
type UpdatePhoneVerificationOptions struct {
}

type UpdatePhoneVerificationOption func(*UpdatePhoneVerificationOptions)


			
func (srv *Account) UpdatePhoneVerification(UserId string, Secret string)  (*appwriteModel.Token, error) {
	path := "/account/verification/phone"

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

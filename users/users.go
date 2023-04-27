package users

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Users service
type Users struct {
	client appwrite.Client
}

func NewUsers(clt appwrite.Client) *Users {
	return &Users{
		client: clt,
	}
}

// List get a list of all the project's users. You can use the query params to
// filter your results.
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

	
func (srv *Users) List(optionalSetters ...ListOption)  (*appwriteModel.UserList, error) {
	path := "/users"

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
	var parsed appwriteModel.UserList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.UserList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Create create a new user.
type CreateOptions struct {
	Email string
	Phone string
	Password string
	Name string
}

type CreateOption func(*CreateOptions)

func WithCreateEmail(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Email = v
	}
}
func WithCreatePhone(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Phone = v
	}
}
func WithCreatePassword(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Password = v
	}
}
func WithCreateName(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Name = v
	}
}

			
func (srv *Users) Create(UserId string, optionalSetters ...CreateOption)  (*appwriteModel.User, error) {
	path := "/users"

	options := &CreateOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = options.Email
	params["phone"] = options.Phone
	params["password"] = options.Password
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateArgon2User create a new user. Password provided must be hashed with
// the [Argon2](https://en.wikipedia.org/wiki/Argon2) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
type CreateArgon2UserOptions struct {
	Name string
}

type CreateArgon2UserOption func(*CreateArgon2UserOptions)

func WithCreateArgon2UserName(v string) CreateArgon2UserOption {
	return func(o *CreateArgon2UserOptions) {
		o.Name = v
	}
}

							
func (srv *Users) CreateArgon2User(UserId string, Email string, Password string, optionalSetters ...CreateArgon2UserOption)  (*appwriteModel.User, error) {
	path := "/users/argon2"

	options := &CreateArgon2UserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateBcryptUser create a new user. Password provided must be hashed with
// the [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
type CreateBcryptUserOptions struct {
	Name string
}

type CreateBcryptUserOption func(*CreateBcryptUserOptions)

func WithCreateBcryptUserName(v string) CreateBcryptUserOption {
	return func(o *CreateBcryptUserOptions) {
		o.Name = v
	}
}

							
func (srv *Users) CreateBcryptUser(UserId string, Email string, Password string, optionalSetters ...CreateBcryptUserOption)  (*appwriteModel.User, error) {
	path := "/users/bcrypt"

	options := &CreateBcryptUserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateMD5User create a new user. Password provided must be hashed with the
// [MD5](https://en.wikipedia.org/wiki/MD5) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
type CreateMD5UserOptions struct {
	Name string
}

type CreateMD5UserOption func(*CreateMD5UserOptions)

func WithCreateMD5UserName(v string) CreateMD5UserOption {
	return func(o *CreateMD5UserOptions) {
		o.Name = v
	}
}

							
func (srv *Users) CreateMD5User(UserId string, Email string, Password string, optionalSetters ...CreateMD5UserOption)  (*appwriteModel.User, error) {
	path := "/users/md5"

	options := &CreateMD5UserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreatePHPassUser create a new user. Password provided must be hashed with
// the [PHPass](https://www.openwall.com/phpass/) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
type CreatePHPassUserOptions struct {
	Name string
}

type CreatePHPassUserOption func(*CreatePHPassUserOptions)

func WithCreatePHPassUserName(v string) CreatePHPassUserOption {
	return func(o *CreatePHPassUserOptions) {
		o.Name = v
	}
}

							
func (srv *Users) CreatePHPassUser(UserId string, Email string, Password string, optionalSetters ...CreatePHPassUserOption)  (*appwriteModel.User, error) {
	path := "/users/phpass"

	options := &CreatePHPassUserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateScryptUser create a new user. Password provided must be hashed with
// the [Scrypt](https://github.com/Tarsnap/scrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
type CreateScryptUserOptions struct {
	Name string
}

type CreateScryptUserOption func(*CreateScryptUserOptions)

func WithCreateScryptUserName(v string) CreateScryptUserOption {
	return func(o *CreateScryptUserOptions) {
		o.Name = v
	}
}

																	
func (srv *Users) CreateScryptUser(UserId string, Email string, Password string, PasswordSalt string, PasswordCpu int, PasswordMemory int, PasswordParallel int, PasswordLength int, optionalSetters ...CreateScryptUserOption)  (*appwriteModel.User, error) {
	path := "/users/scrypt"

	options := &CreateScryptUserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["passwordSalt"] = PasswordSalt
	params["passwordCpu"] = PasswordCpu
	params["passwordMemory"] = PasswordMemory
	params["passwordParallel"] = PasswordParallel
	params["passwordLength"] = PasswordLength
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateScryptModifiedUser create a new user. Password provided must be
// hashed with the [Scrypt
// Modified](https://gist.github.com/Meldiron/eecf84a0225eccb5a378d45bb27462cc)
// algorithm. Use the [POST /users](/docs/server/users#usersCreate) endpoint
// to create users with a plain text password.
type CreateScryptModifiedUserOptions struct {
	Name string
}

type CreateScryptModifiedUserOption func(*CreateScryptModifiedUserOptions)

func WithCreateScryptModifiedUserName(v string) CreateScryptModifiedUserOption {
	return func(o *CreateScryptModifiedUserOptions) {
		o.Name = v
	}
}

													
func (srv *Users) CreateScryptModifiedUser(UserId string, Email string, Password string, PasswordSalt string, PasswordSaltSeparator string, PasswordSignerKey string, optionalSetters ...CreateScryptModifiedUserOption)  (*appwriteModel.User, error) {
	path := "/users/scrypt-modified"

	options := &CreateScryptModifiedUserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["passwordSalt"] = PasswordSalt
	params["passwordSaltSeparator"] = PasswordSaltSeparator
	params["passwordSignerKey"] = PasswordSignerKey
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// CreateSHAUser create a new user. Password provided must be hashed with the
// [SHA](https://en.wikipedia.org/wiki/Secure_Hash_Algorithm) algorithm. Use
// the [POST /users](/docs/server/users#usersCreate) endpoint to create users
// with a plain text password.
type CreateSHAUserOptions struct {
	PasswordVersion string
	Name string
}

type CreateSHAUserOption func(*CreateSHAUserOptions)

func WithCreateSHAUserPasswordVersion(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.PasswordVersion = v
	}
}
func WithCreateSHAUserName(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.Name = v
	}
}

							
func (srv *Users) CreateSHAUser(UserId string, Email string, Password string, optionalSetters ...CreateSHAUserOption)  (*appwriteModel.User, error) {
	path := "/users/sha"

	options := &CreateSHAUserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["passwordVersion"] = options.PasswordVersion
	params["name"] = options.Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
// Get get a user by its unique ID.
type GetOptions struct {
}

type GetOption func(*GetOptions)


	
func (srv *Users) Get(UserId string)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// Delete delete a user by its unique ID, thereby releasing it's ID. Since ID
// is released and can be reused, all user-related resources like documents or
// storage files should be deleted before user deletion. If you want to keep
// ID reserved, use the [updateStatus](/docs/server/users#usersUpdateStatus)
// endpoint instead.
type DeleteOptions struct {
}

type DeleteOption func(*DeleteOptions)


	
func (srv *Users) Delete(UserId string)  (*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdateEmail update the user email by its unique ID.
type UpdateEmailOptions struct {
}

type UpdateEmailOption func(*UpdateEmailOptions)


			
func (srv *Users) UpdateEmail(UserId string, Email string)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/email")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
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
// ListLogs get the user activity logs list by its unique ID.
type ListLogsOptions struct {
	Queries []interface{}
}

type ListLogsOption func(*ListLogsOptions)

func WithListLogsQueries(v []interface{}) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
	}
}

			
func (srv *Users) ListLogs(UserId string, optionalSetters ...ListLogsOption)  (*appwriteModel.LogList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")

	options := &ListLogsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
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
// ListMemberships get the user membership list by its unique ID.
type ListMembershipsOptions struct {
}

type ListMembershipsOption func(*ListMembershipsOptions)


	
func (srv *Users) ListMemberships(UserId string)  (*appwriteModel.MembershipList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/memberships")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdateName update the user name by its unique ID.
type UpdateNameOptions struct {
}

type UpdateNameOption func(*UpdateNameOptions)


			
func (srv *Users) UpdateName(UserId string, Name string)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/name")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdatePassword update the user password by its unique ID.
type UpdatePasswordOptions struct {
}

type UpdatePasswordOption func(*UpdatePasswordOptions)


			
func (srv *Users) UpdatePassword(UserId string, Password string)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/password")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdatePhone update the user phone by its unique ID.
type UpdatePhoneOptions struct {
}

type UpdatePhoneOption func(*UpdatePhoneOptions)


			
func (srv *Users) UpdatePhone(UserId string, Number string)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/phone")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["number"] = Number
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
// GetPrefs get the user preferences by its unique ID.
type GetPrefsOptions struct {
}

type GetPrefsOption func(*GetPrefsOptions)


	
func (srv *Users) GetPrefs(UserId string)  (*appwriteModel.Preferences, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdatePrefs update the user preferences by its unique ID. The object you
// pass is stored as is, and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws error if exceeded.
type UpdatePrefsOptions struct {
}

type UpdatePrefsOption func(*UpdatePrefsOptions)


			
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{})  (*appwriteModel.Preferences, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
// ListSessions get the user sessions list by its unique ID.
type ListSessionsOptions struct {
}

type ListSessionsOption func(*ListSessionsOptions)


	
func (srv *Users) ListSessions(UserId string)  (*appwriteModel.SessionList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// DeleteSessions delete all user's sessions by using the user's unique ID.
type DeleteSessionsOptions struct {
}

type DeleteSessionsOption func(*DeleteSessionsOptions)


	
func (srv *Users) DeleteSessions(UserId string)  (*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// DeleteSession delete a user sessions by its unique ID.
type DeleteSessionOptions struct {
}

type DeleteSessionOption func(*DeleteSessionOptions)


			
func (srv *Users) DeleteSession(UserId string, SessionId string)  (*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{sessionId}", SessionId)
	path := r.Replace("/users/{userId}/sessions/{sessionId}")

	params := map[string]interface{}{}
	params["userId"] = UserId
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
// UpdateStatus update the user status by its unique ID. Use this endpoint as
// an alternative to deleting a user if you want to keep user's ID reserved.
type UpdateStatusOptions struct {
}

type UpdateStatusOption func(*UpdateStatusOptions)


			
func (srv *Users) UpdateStatus(UserId string, Status bool)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/status")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["status"] = Status
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
// UpdateEmailVerification update the user email verification status by its
// unique ID.
type UpdateEmailVerificationOptions struct {
}

type UpdateEmailVerificationOption func(*UpdateEmailVerificationOptions)


			
func (srv *Users) UpdateEmailVerification(UserId string, EmailVerification bool)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["emailVerification"] = EmailVerification
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
// UpdatePhoneVerification update the user phone verification status by its
// unique ID.
type UpdatePhoneVerificationOptions struct {
}

type UpdatePhoneVerificationOption func(*UpdatePhoneVerificationOptions)


			
func (srv *Users) UpdatePhoneVerification(UserId string, PhoneVerification bool)  (*appwriteModel.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification/phone")

	params := map[string]interface{}{}
	params["userId"] = UserId
	params["phoneVerification"] = PhoneVerification
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

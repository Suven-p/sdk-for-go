package users

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Users service
type Users struct {
	client client.Client
}

func NewUsers(clt client.Client) *Users {
	return &Users{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []interface{}
	Search         string
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListOption func(*ListOptions)

func WithListQueries(v []interface{}) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// List get a list of all the project's users. You can use the query params to
// filter your results.
func (srv *Users) List(optionalSetters ...ListOption) (*models.UserList, error) {
	path := "/users"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.UserList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.UserList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Email          string
	Phone          string
	Password       string
	Name           string
	enabledSetters map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Email":    false,
		"Phone":    false,
		"Password": false,
		"Name":     false,
	}
	return &options
}

type CreateOption func(*CreateOptions)

func WithCreateEmail(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func WithCreatePhone(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func WithCreatePassword(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func WithCreateName(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// Create create a new user.
func (srv *Users) Create(UserId string, optionalSetters ...CreateOption) (*models.User, error) {
	path := "/users"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateArgon2UserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateArgon2UserOptions) New() *CreateArgon2UserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateArgon2UserOption func(*CreateArgon2UserOptions)

func WithCreateArgon2UserName(v string) CreateArgon2UserOption {
	return func(o *CreateArgon2UserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateArgon2User create a new user. Password provided must be hashed with
// the [Argon2](https://en.wikipedia.org/wiki/Argon2) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateArgon2User(UserId string, Email string, Password string, optionalSetters ...CreateArgon2UserOption) (*models.User, error) {
	path := "/users/argon2"
	options := CreateArgon2UserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBcryptUserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateBcryptUserOptions) New() *CreateBcryptUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateBcryptUserOption func(*CreateBcryptUserOptions)

func WithCreateBcryptUserName(v string) CreateBcryptUserOption {
	return func(o *CreateBcryptUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateBcryptUser create a new user. Password provided must be hashed with
// the [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateBcryptUser(UserId string, Email string, Password string, optionalSetters ...CreateBcryptUserOption) (*models.User, error) {
	path := "/users/bcrypt"
	options := CreateBcryptUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateMD5UserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateMD5UserOptions) New() *CreateMD5UserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateMD5UserOption func(*CreateMD5UserOptions)

func WithCreateMD5UserName(v string) CreateMD5UserOption {
	return func(o *CreateMD5UserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateMD5User create a new user. Password provided must be hashed with the
// [MD5](https://en.wikipedia.org/wiki/MD5) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateMD5User(UserId string, Email string, Password string, optionalSetters ...CreateMD5UserOption) (*models.User, error) {
	path := "/users/md5"
	options := CreateMD5UserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreatePHPassUserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreatePHPassUserOptions) New() *CreatePHPassUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreatePHPassUserOption func(*CreatePHPassUserOptions)

func WithCreatePHPassUserName(v string) CreatePHPassUserOption {
	return func(o *CreatePHPassUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreatePHPassUser create a new user. Password provided must be hashed with
// the [PHPass](https://www.openwall.com/phpass/) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreatePHPassUser(UserId string, Email string, Password string, optionalSetters ...CreatePHPassUserOption) (*models.User, error) {
	path := "/users/phpass"
	options := CreatePHPassUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateScryptUserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateScryptUserOptions) New() *CreateScryptUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateScryptUserOption func(*CreateScryptUserOptions)

func WithCreateScryptUserName(v string) CreateScryptUserOption {
	return func(o *CreateScryptUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateScryptUser create a new user. Password provided must be hashed with
// the [Scrypt](https://github.com/Tarsnap/scrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateScryptUser(UserId string, Email string, Password string, PasswordSalt string, PasswordCpu int, PasswordMemory int, PasswordParallel int, PasswordLength int, optionalSetters ...CreateScryptUserOption) (*models.User, error) {
	path := "/users/scrypt"
	options := CreateScryptUserOptions{}.New()
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
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateScryptModifiedUserOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateScryptModifiedUserOptions) New() *CreateScryptModifiedUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateScryptModifiedUserOption func(*CreateScryptModifiedUserOptions)

func WithCreateScryptModifiedUserName(v string) CreateScryptModifiedUserOption {
	return func(o *CreateScryptModifiedUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateScryptModifiedUser create a new user. Password provided must be
// hashed with the [Scrypt
// Modified](https://gist.github.com/Meldiron/eecf84a0225eccb5a378d45bb27462cc)
// algorithm. Use the [POST /users](/docs/server/users#usersCreate) endpoint
// to create users with a plain text password.
func (srv *Users) CreateScryptModifiedUser(UserId string, Email string, Password string, PasswordSalt string, PasswordSaltSeparator string, PasswordSignerKey string, optionalSetters ...CreateScryptModifiedUserOption) (*models.User, error) {
	path := "/users/scrypt-modified"
	options := CreateScryptModifiedUserOptions{}.New()
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
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateSHAUserOptions struct {
	PasswordVersion string
	Name            string
	enabledSetters  map[string]bool
}

func (options CreateSHAUserOptions) New() *CreateSHAUserOptions {
	options.enabledSetters = map[string]bool{
		"PasswordVersion": false,
		"Name":            false,
	}
	return &options
}

type CreateSHAUserOption func(*CreateSHAUserOptions)

func WithCreateSHAUserPasswordVersion(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.PasswordVersion = v
		o.enabledSetters["PasswordVersion"] = true
	}
}
func WithCreateSHAUserName(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateSHAUser create a new user. Password provided must be hashed with the
// [SHA](https://en.wikipedia.org/wiki/Secure_Hash_Algorithm) algorithm. Use
// the [POST /users](/docs/server/users#usersCreate) endpoint to create users
// with a plain text password.
func (srv *Users) CreateSHAUser(UserId string, Email string, Password string, optionalSetters ...CreateSHAUserOption) (*models.User, error) {
	path := "/users/sha"
	options := CreateSHAUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["PasswordVersion"] {
		params["passwordVersion"] = options.PasswordVersion
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a user by its unique ID.
func (srv *Users) Get(UserId string) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
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
func (srv *Users) Delete(UserId string) (*interface{}, error) {
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
func (srv *Users) UpdateEmail(UserId string, Email string) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListLogsOptions struct {
	Queries        []interface{}
	enabledSetters map[string]bool
}

func (options ListLogsOptions) New() *ListLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListLogsOption func(*ListLogsOptions)

func WithListLogsQueries(v []interface{}) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListLogs get the user activity logs list by its unique ID.
func (srv *Users) ListLogs(UserId string, optionalSetters ...ListLogsOption) (*models.LogList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")
	options := ListLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.LogList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListMemberships get the user membership list by its unique ID.
func (srv *Users) ListMemberships(UserId string) (*models.MembershipList, error) {
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
	var parsed models.MembershipList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MembershipList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateName update the user name by its unique ID.
func (srv *Users) UpdateName(UserId string, Name string) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdatePassword update the user password by its unique ID.
func (srv *Users) UpdatePassword(UserId string, Password string) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdatePhone update the user phone by its unique ID.
func (srv *Users) UpdatePhone(UserId string, Number string) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPrefs get the user preferences by its unique ID.
func (srv *Users) GetPrefs(UserId string) (*models.Preferences, error) {
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
	var parsed models.Preferences
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdatePrefs update the user preferences by its unique ID. The object you
// pass is stored as is, and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws error if exceeded.
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{}) (*models.Preferences, error) {
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
	var parsed models.Preferences
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSessions get the user sessions list by its unique ID.
func (srv *Users) ListSessions(UserId string) (*models.SessionList, error) {
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
	var parsed models.SessionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.SessionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteSessions delete all user's sessions by using the user's unique ID.
func (srv *Users) DeleteSessions(UserId string) (*interface{}, error) {
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
func (srv *Users) DeleteSession(UserId string, SessionId string) (*interface{}, error) {
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
func (srv *Users) UpdateStatus(UserId string, Status bool) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateEmailVerification update the user email verification status by its
// unique ID.
func (srv *Users) UpdateEmailVerification(UserId string, EmailVerification bool) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdatePhoneVerification update the user phone verification status by its
// unique ID.
func (srv *Users) UpdatePhoneVerification(UserId string, PhoneVerification bool) (*models.User, error) {
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

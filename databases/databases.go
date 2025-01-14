package databases

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Databases service
type Databases struct {
	client client.Client
}

func NewDatabases(clt client.Client) *Databases {
	return &Databases{
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

// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
func (srv *Databases) List(optionalSetters ...ListOption) (*models.DatabaseList, error) {
	path := "/databases"
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
	var parsed models.DatabaseList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.DatabaseList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Create create a new Database.
func (srv *Databases) Create(DatabaseId string, Name string) (*models.Database, error) {
	path := "/databases"
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a database by its unique ID. This endpoint response returns a JSON
// object with the database metadata.
func (srv *Databases) Get(DatabaseId string) (*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Update update a database by its unique ID.
func (srv *Databases) Update(DatabaseId string, Name string) (*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a database by its unique ID. Only API keys with with
// databases.write scope can delete a database.
func (srv *Databases) Delete(DatabaseId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
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

type ListCollectionsOptions struct {
	Queries        []interface{}
	Search         string
	enabledSetters map[string]bool
}

func (options ListCollectionsOptions) New() *ListCollectionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListCollectionsOption func(*ListCollectionsOptions)

func WithListCollectionsQueries(v []interface{}) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListCollectionsSearch(v string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// ListCollections get a list of all collections that belong to the provided
// databaseId. You can use the search parameter to filter your results.
func (srv *Databases) ListCollections(DatabaseId string, optionalSetters ...ListCollectionsOption) (*models.CollectionList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")
	options := ListCollectionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
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
	var parsed models.CollectionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.CollectionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateCollectionOptions struct {
	Permissions      []interface{}
	DocumentSecurity bool
	enabledSetters   map[string]bool
}

func (options CreateCollectionOptions) New() *CreateCollectionOptions {
	options.enabledSetters = map[string]bool{
		"Permissions":      false,
		"DocumentSecurity": false,
	}
	return &options
}

type CreateCollectionOption func(*CreateCollectionOptions)

func WithCreateCollectionPermissions(v []interface{}) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func WithCreateCollectionDocumentSecurity(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}

// CreateCollection create a new Collection. Before using this route, you
// should create a new database resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
func (srv *Databases) CreateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...CreateCollectionOption) (*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")
	options := CreateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *Databases) GetCollection(DatabaseId string, CollectionId string) (*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateCollectionOptions struct {
	Permissions      []interface{}
	DocumentSecurity bool
	Enabled          bool
	enabledSetters   map[string]bool
}

func (options UpdateCollectionOptions) New() *UpdateCollectionOptions {
	options.enabledSetters = map[string]bool{
		"Permissions":      false,
		"DocumentSecurity": false,
		"Enabled":          false,
	}
	return &options
}

type UpdateCollectionOption func(*UpdateCollectionOptions)

func WithUpdateCollectionPermissions(v []interface{}) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func WithUpdateCollectionDocumentSecurity(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}
func WithUpdateCollectionEnabled(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// UpdateCollection update a collection by its unique ID.
func (srv *Databases) UpdateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...UpdateCollectionOption) (*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	options := UpdateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *Databases) DeleteCollection(DatabaseId string, CollectionId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
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

// ListAttributes
func (srv *Databases) ListAttributes(DatabaseId string, CollectionId string) (*models.AttributeList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBooleanAttributeOptions struct {
	Default        bool
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateBooleanAttributeOptions) New() *CreateBooleanAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateBooleanAttributeOption func(*CreateBooleanAttributeOptions)

func WithCreateBooleanAttributeDefault(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateBooleanAttributeArray(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateBooleanAttribute create a boolean attribute.
func (srv *Databases) CreateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateBooleanAttributeOption) (*models.AttributeBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean")
	options := CreateBooleanAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeBoolean
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateBooleanAttribute
func (srv *Databases) UpdateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default bool) (*models.AttributeBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeBoolean
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDatetimeAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateDatetimeAttributeOptions) New() *CreateDatetimeAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateDatetimeAttributeOption func(*CreateDatetimeAttributeOptions)

func WithCreateDatetimeAttributeDefault(v string) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateDatetimeAttributeArray(v bool) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateDatetimeAttribute
func (srv *Databases) CreateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateDatetimeAttributeOption) (*models.AttributeDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime")
	options := CreateDatetimeAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeDatetime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateDatetimeAttribute
func (srv *Databases) UpdateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string) (*models.AttributeDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeDatetime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateEmailAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateEmailAttributeOptions) New() *CreateEmailAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateEmailAttributeOption func(*CreateEmailAttributeOptions)

func WithCreateEmailAttributeDefault(v string) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateEmailAttributeArray(v bool) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateEmailAttribute create an email attribute.
func (srv *Databases) CreateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateEmailAttributeOption) (*models.AttributeEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email")
	options := CreateEmailAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeEmail
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateEmailAttribute update an email attribute. Changing the `default`
// value will not update already existing documents.
func (srv *Databases) UpdateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string) (*models.AttributeEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeEmail
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateEnumAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateEnumAttributeOptions) New() *CreateEnumAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateEnumAttributeOption func(*CreateEnumAttributeOptions)

func WithCreateEnumAttributeDefault(v string) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateEnumAttributeArray(v bool) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateEnumAttribute
func (srv *Databases) CreateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []interface{}, Required bool, optionalSetters ...CreateEnumAttributeOption) (*models.AttributeEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum")
	options := CreateEnumAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["elements"] = Elements
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeEnum
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateEnumAttribute update an enum attribute. Changing the `default` value
// will not update already existing documents.
func (srv *Databases) UpdateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []interface{}, Required bool, Default string) (*models.AttributeEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["elements"] = Elements
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeEnum
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateFloatAttributeOptions struct {
	Min            float64
	Max            float64
	Default        float64
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateFloatAttributeOptions) New() *CreateFloatAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min":     false,
		"Max":     false,
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateFloatAttributeOption func(*CreateFloatAttributeOptions)

func WithCreateFloatAttributeMin(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func WithCreateFloatAttributeMax(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func WithCreateFloatAttributeDefault(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateFloatAttributeArray(v bool) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateFloatAttribute create a float attribute. Optionally, minimum and
// maximum values can be provided.
func (srv *Databases) CreateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateFloatAttributeOption) (*models.AttributeFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float")
	options := CreateFloatAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeFloat
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateFloatAttribute update a float attribute. Changing the `default` value
// will not update already existing documents.
func (srv *Databases) UpdateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min float64, Max float64, Default float64) (*models.AttributeFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["min"] = Min
	params["max"] = Max
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeFloat
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateIntegerAttributeOptions struct {
	Min            int
	Max            int
	Default        int
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateIntegerAttributeOptions) New() *CreateIntegerAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min":     false,
		"Max":     false,
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateIntegerAttributeOption func(*CreateIntegerAttributeOptions)

func WithCreateIntegerAttributeMin(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func WithCreateIntegerAttributeMax(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func WithCreateIntegerAttributeDefault(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateIntegerAttributeArray(v bool) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateIntegerAttribute create an integer attribute. Optionally, minimum and
// maximum values can be provided.
func (srv *Databases) CreateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIntegerAttributeOption) (*models.AttributeInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer")
	options := CreateIntegerAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeInteger
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateIntegerAttribute update an integer attribute. Changing the `default`
// value will not update already existing documents.
func (srv *Databases) UpdateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min int, Max int, Default int) (*models.AttributeInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["min"] = Min
	params["max"] = Max
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeInteger
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateIpAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateIpAttributeOptions) New() *CreateIpAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateIpAttributeOption func(*CreateIpAttributeOptions)

func WithCreateIpAttributeDefault(v string) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateIpAttributeArray(v bool) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateIpAttribute create IP address attribute.
func (srv *Databases) CreateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIpAttributeOption) (*models.AttributeIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip")
	options := CreateIpAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeIp
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateIpAttribute update an ip attribute. Changing the `default` value will
// not update already existing documents.
func (srv *Databases) UpdateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string) (*models.AttributeIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeIp
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRelationshipAttributeOptions struct {
	TwoWay         bool
	Key            string
	TwoWayKey      string
	OnDelete       string
	enabledSetters map[string]bool
}

func (options CreateRelationshipAttributeOptions) New() *CreateRelationshipAttributeOptions {
	options.enabledSetters = map[string]bool{
		"TwoWay":    false,
		"Key":       false,
		"TwoWayKey": false,
		"OnDelete":  false,
	}
	return &options
}

type CreateRelationshipAttributeOption func(*CreateRelationshipAttributeOptions)

func WithCreateRelationshipAttributeTwoWay(v bool) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWay = v
		o.enabledSetters["TwoWay"] = true
	}
}
func WithCreateRelationshipAttributeKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func WithCreateRelationshipAttributeTwoWayKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWayKey = v
		o.enabledSetters["TwoWayKey"] = true
	}
}
func WithCreateRelationshipAttributeOnDelete(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}

// CreateRelationshipAttribute create relationship attribute. [Learn more
// about relationship
// attributes](/docs/databases-relationships#relationship-attributes).
func (srv *Databases) CreateRelationshipAttribute(DatabaseId string, CollectionId string, RelatedCollectionId string, Type string, optionalSetters ...CreateRelationshipAttributeOption) (*models.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/relationship")
	options := CreateRelationshipAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["relatedCollectionId"] = RelatedCollectionId
	params["type"] = Type
	if options.enabledSetters["TwoWay"] {
		params["twoWay"] = options.TwoWay
	}
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["TwoWayKey"] {
		params["twoWayKey"] = options.TwoWayKey
	}
	if options.enabledSetters["OnDelete"] {
		params["onDelete"] = options.OnDelete
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeRelationship
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateStringAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateStringAttributeOptions) New() *CreateStringAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateStringAttributeOption func(*CreateStringAttributeOptions)

func WithCreateStringAttributeDefault(v string) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateStringAttributeArray(v bool) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateStringAttribute create a string attribute.
func (srv *Databases) CreateStringAttribute(DatabaseId string, CollectionId string, Key string, Size int, Required bool, optionalSetters ...CreateStringAttributeOption) (*models.AttributeString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string")
	options := CreateStringAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["size"] = Size
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeString
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateStringAttribute update a string attribute. Changing the `default`
// value will not update already existing documents.
func (srv *Databases) UpdateStringAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string) (*models.AttributeString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeString
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateUrlAttributeOptions struct {
	Default        string
	Array          bool
	enabledSetters map[string]bool
}

func (options CreateUrlAttributeOptions) New() *CreateUrlAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array":   false,
	}
	return &options
}

type CreateUrlAttributeOption func(*CreateUrlAttributeOptions)

func WithCreateUrlAttributeDefault(v string) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func WithCreateUrlAttributeArray(v bool) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}

// CreateUrlAttribute create a URL attribute.
func (srv *Databases) CreateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateUrlAttributeOption) (*models.AttributeUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url")
	options := CreateUrlAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeUrl
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateUrlAttribute update an url attribute. Changing the `default` value
// will not update already existing documents.
func (srv *Databases) UpdateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string) (*models.AttributeUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeUrl
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetAttribute
func (srv *Databases) GetAttribute(DatabaseId string, CollectionId string, Key string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
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

// DeleteAttribute
func (srv *Databases) DeleteAttribute(DatabaseId string, CollectionId string, Key string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
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

type UpdateRelationshipAttributeOptions struct {
	OnDelete       string
	enabledSetters map[string]bool
}

func (options UpdateRelationshipAttributeOptions) New() *UpdateRelationshipAttributeOptions {
	options.enabledSetters = map[string]bool{
		"OnDelete": false,
	}
	return &options
}

type UpdateRelationshipAttributeOption func(*UpdateRelationshipAttributeOptions)

func WithUpdateRelationshipAttributeOnDelete(v string) UpdateRelationshipAttributeOption {
	return func(o *UpdateRelationshipAttributeOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}

// UpdateRelationshipAttribute update relationship attribute. [Learn more
// about relationship
// attributes](/docs/databases-relationships#relationship-attributes).
func (srv *Databases) UpdateRelationshipAttribute(DatabaseId string, CollectionId string, Key string, optionalSetters ...UpdateRelationshipAttributeOption) (*models.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}/relationship")
	options := UpdateRelationshipAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	if options.enabledSetters["OnDelete"] {
		params["onDelete"] = options.OnDelete
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.AttributeRelationship
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListDocumentsOptions struct {
	Queries        []interface{}
	enabledSetters map[string]bool
}

func (options ListDocumentsOptions) New() *ListDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListDocumentsOption func(*ListDocumentsOptions)

func WithListDocumentsQueries(v []interface{}) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
func (srv *Databases) ListDocuments(DatabaseId string, CollectionId string, optionalSetters ...ListDocumentsOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := ListDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
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
	var parsed models.DocumentList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDocumentOptions struct {
	Permissions    []interface{}
	enabledSetters map[string]bool
}

func (options CreateDocumentOptions) New() *CreateDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
	}
	return &options
}

type CreateDocumentOption func(*CreateDocumentOptions)

func WithCreateDocumentPermissions(v []interface{}) CreateDocumentOption {
	return func(o *CreateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}

// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
func (srv *Databases) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, optionalSetters ...CreateDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := CreateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["data"] = Data
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetDocumentOptions struct {
	Queries        []interface{}
	enabledSetters map[string]bool
}

func (options GetDocumentOptions) New() *GetDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type GetDocumentOption func(*GetDocumentOptions)

func WithGetDocumentQueries(v []interface{}) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
func (srv *Databases) GetDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...GetDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := GetDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
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
	var parsed models.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateDocumentOptions struct {
	Data           interface{}
	Permissions    []interface{}
	enabledSetters map[string]bool
}

func (options UpdateDocumentOptions) New() *UpdateDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Data":        false,
		"Permissions": false,
	}
	return &options
}

type UpdateDocumentOption func(*UpdateDocumentOptions)

func WithUpdateDocumentData(v interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func WithUpdateDocumentPermissions(v []interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}

// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
func (srv *Databases) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...UpdateDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := UpdateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteDocument delete a document by its unique ID.
func (srv *Databases) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
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

// ListIndexes
func (srv *Databases) ListIndexes(DatabaseId string, CollectionId string) (*models.IndexList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.IndexList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.IndexList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateIndexOptions struct {
	Orders         []interface{}
	enabledSetters map[string]bool
}

func (options CreateIndexOptions) New() *CreateIndexOptions {
	options.enabledSetters = map[string]bool{
		"Orders": false,
	}
	return &options
}

type CreateIndexOption func(*CreateIndexOptions)

func WithCreateIndexOrders(v []interface{}) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Orders = v
		o.enabledSetters["Orders"] = true
	}
}

// CreateIndex
func (srv *Databases) CreateIndex(DatabaseId string, CollectionId string, Key string, Type string, Attributes []interface{}, optionalSetters ...CreateIndexOption) (*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")
	options := CreateIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["type"] = Type
	params["attributes"] = Attributes
	if options.enabledSetters["Orders"] {
		params["orders"] = options.Orders
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Index
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetIndex
func (srv *Databases) GetIndex(DatabaseId string, CollectionId string, Key string) (*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Index
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteIndex
func (srv *Databases) DeleteIndex(DatabaseId string, CollectionId string, Key string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
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

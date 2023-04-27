package databases

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Databases service
type Databases struct {
	client appwrite.Client
}

func NewDatabases(clt appwrite.Client) *Databases {
	return &Databases{
		client: clt,
	}
}

// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
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

	
func (srv *Databases) List(optionalSetters ...ListOption)  (*appwriteModel.DatabaseList, error) {
	path := "/databases"

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
	var parsed appwriteModel.DatabaseList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.DatabaseList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Create create a new Database.
// 
type CreateOptions struct {
}

type CreateOption func(*CreateOptions)


			
func (srv *Databases) Create(DatabaseId string, Name string)  (*appwriteModel.Database, error) {
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
	var parsed appwriteModel.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Get get a database by its unique ID. This endpoint response returns a JSON
// object with the database metadata.
type GetOptions struct {
}

type GetOption func(*GetOptions)


	
func (srv *Databases) Get(DatabaseId string)  (*appwriteModel.Database, error) {
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
	var parsed appwriteModel.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Update update a database by its unique ID.
type UpdateOptions struct {
}

type UpdateOption func(*UpdateOptions)


			
func (srv *Databases) Update(DatabaseId string, Name string)  (*appwriteModel.Database, error) {
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
	var parsed appwriteModel.Database
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Delete delete a database by its unique ID. Only API keys with with
// databases.write scope can delete a database.
type DeleteOptions struct {
}

type DeleteOption func(*DeleteOptions)


	
func (srv *Databases) Delete(DatabaseId string)  (*interface{}, error) {
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
// ListCollections get a list of all collections that belong to the provided
// databaseId. You can use the search parameter to filter your results.
type ListCollectionsOptions struct {
	Queries []interface{}
	Search string
}

type ListCollectionsOption func(*ListCollectionsOptions)

func WithListCollectionsQueries(v []interface{}) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Queries = v
	}
}
func WithListCollectionsSearch(v string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Search = v
	}
}

			
func (srv *Databases) ListCollections(DatabaseId string, optionalSetters ...ListCollectionsOption)  (*appwriteModel.CollectionList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")

	options := &ListCollectionsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.CollectionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.CollectionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateCollection create a new Collection. Before using this route, you
// should create a new database resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
type CreateCollectionOptions struct {
	Permissions []interface{}
	DocumentSecurity bool
}

type CreateCollectionOption func(*CreateCollectionOptions)

func WithCreateCollectionPermissions(v []interface{}) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Permissions = v
	}
}
func WithCreateCollectionDocumentSecurity(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.DocumentSecurity = v
	}
}

							
func (srv *Databases) CreateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...CreateCollectionOption)  (*appwriteModel.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")

	options := &CreateCollectionOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	params["permissions"] = options.Permissions
	params["documentSecurity"] = options.DocumentSecurity
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
type GetCollectionOptions struct {
}

type GetCollectionOption func(*GetCollectionOptions)


			
func (srv *Databases) GetCollection(DatabaseId string, CollectionId string)  (*appwriteModel.Collection, error) {
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
	var parsed appwriteModel.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateCollection update a collection by its unique ID.
type UpdateCollectionOptions struct {
	Permissions []interface{}
	DocumentSecurity bool
	Enabled bool
}

type UpdateCollectionOption func(*UpdateCollectionOptions)

func WithUpdateCollectionPermissions(v []interface{}) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Permissions = v
	}
}
func WithUpdateCollectionDocumentSecurity(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.DocumentSecurity = v
	}
}
func WithUpdateCollectionEnabled(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Enabled = v
	}
}

							
func (srv *Databases) UpdateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...UpdateCollectionOption)  (*appwriteModel.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")

	options := &UpdateCollectionOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	params["permissions"] = options.Permissions
	params["documentSecurity"] = options.DocumentSecurity
	params["enabled"] = options.Enabled
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Collection
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
type DeleteCollectionOptions struct {
}

type DeleteCollectionOption func(*DeleteCollectionOptions)


			
func (srv *Databases) DeleteCollection(DatabaseId string, CollectionId string)  (*interface{}, error) {
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
type ListAttributesOptions struct {
}

type ListAttributesOption func(*ListAttributesOptions)


			
func (srv *Databases) ListAttributes(DatabaseId string, CollectionId string)  (*appwriteModel.AttributeList, error) {
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
	var parsed appwriteModel.AttributeList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateBooleanAttribute create a boolean attribute.
// 
type CreateBooleanAttributeOptions struct {
	Default bool
	Array bool
}

type CreateBooleanAttributeOption func(*CreateBooleanAttributeOptions)

func WithCreateBooleanAttributeDefault(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Default = v
	}
}
func WithCreateBooleanAttributeArray(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateBooleanAttributeOption)  (*appwriteModel.AttributeBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean")

	options := &CreateBooleanAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeBoolean
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateBooleanAttribute
type UpdateBooleanAttributeOptions struct {
}

type UpdateBooleanAttributeOption func(*UpdateBooleanAttributeOptions)


									
func (srv *Databases) UpdateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default bool)  (*appwriteModel.AttributeBoolean, error) {
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
	var parsed appwriteModel.AttributeBoolean
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateDatetimeAttribute
type CreateDatetimeAttributeOptions struct {
	Default string
	Array bool
}

type CreateDatetimeAttributeOption func(*CreateDatetimeAttributeOptions)

func WithCreateDatetimeAttributeDefault(v string) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Default = v
	}
}
func WithCreateDatetimeAttributeArray(v bool) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateDatetimeAttributeOption)  (*appwriteModel.AttributeDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime")

	options := &CreateDatetimeAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeDatetime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateDatetimeAttribute
type UpdateDatetimeAttributeOptions struct {
}

type UpdateDatetimeAttributeOption func(*UpdateDatetimeAttributeOptions)


									
func (srv *Databases) UpdateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string)  (*appwriteModel.AttributeDatetime, error) {
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
	var parsed appwriteModel.AttributeDatetime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateEmailAttribute create an email attribute.
// 
type CreateEmailAttributeOptions struct {
	Default string
	Array bool
}

type CreateEmailAttributeOption func(*CreateEmailAttributeOptions)

func WithCreateEmailAttributeDefault(v string) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Default = v
	}
}
func WithCreateEmailAttributeArray(v bool) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateEmailAttributeOption)  (*appwriteModel.AttributeEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email")

	options := &CreateEmailAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeEmail
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateEmailAttribute update an email attribute. Changing the `default`
// value will not update already existing documents.
// 
type UpdateEmailAttributeOptions struct {
}

type UpdateEmailAttributeOption func(*UpdateEmailAttributeOptions)


									
func (srv *Databases) UpdateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string)  (*appwriteModel.AttributeEmail, error) {
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
	var parsed appwriteModel.AttributeEmail
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateEnumAttribute
type CreateEnumAttributeOptions struct {
	Default string
	Array bool
}

type CreateEnumAttributeOption func(*CreateEnumAttributeOptions)

func WithCreateEnumAttributeDefault(v string) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Default = v
	}
}
func WithCreateEnumAttributeArray(v bool) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Array = v
	}
}

											
func (srv *Databases) CreateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []interface{}, Required bool, optionalSetters ...CreateEnumAttributeOption)  (*appwriteModel.AttributeEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum")

	options := &CreateEnumAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["elements"] = Elements
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeEnum
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateEnumAttribute update an enum attribute. Changing the `default` value
// will not update already existing documents.
// 
type UpdateEnumAttributeOptions struct {
}

type UpdateEnumAttributeOption func(*UpdateEnumAttributeOptions)


											
func (srv *Databases) UpdateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []interface{}, Required bool, Default string)  (*appwriteModel.AttributeEnum, error) {
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
	var parsed appwriteModel.AttributeEnum
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateFloatAttribute create a float attribute. Optionally, minimum and
// maximum values can be provided.
// 
type CreateFloatAttributeOptions struct {
	Min float64
	Max float64
	Default float64
	Array bool
}

type CreateFloatAttributeOption func(*CreateFloatAttributeOptions)

func WithCreateFloatAttributeMin(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Min = v
	}
}
func WithCreateFloatAttributeMax(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Max = v
	}
}
func WithCreateFloatAttributeDefault(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Default = v
	}
}
func WithCreateFloatAttributeArray(v bool) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateFloatAttributeOption)  (*appwriteModel.AttributeFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float")

	options := &CreateFloatAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["min"] = options.Min
	params["max"] = options.Max
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeFloat
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateFloatAttribute update a float attribute. Changing the `default` value
// will not update already existing documents.
// 
type UpdateFloatAttributeOptions struct {
}

type UpdateFloatAttributeOption func(*UpdateFloatAttributeOptions)


													
func (srv *Databases) UpdateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min float64, Max float64, Default float64)  (*appwriteModel.AttributeFloat, error) {
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
	var parsed appwriteModel.AttributeFloat
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateIntegerAttribute create an integer attribute. Optionally, minimum and
// maximum values can be provided.
// 
type CreateIntegerAttributeOptions struct {
	Min int
	Max int
	Default int
	Array bool
}

type CreateIntegerAttributeOption func(*CreateIntegerAttributeOptions)

func WithCreateIntegerAttributeMin(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Min = v
	}
}
func WithCreateIntegerAttributeMax(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Max = v
	}
}
func WithCreateIntegerAttributeDefault(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Default = v
	}
}
func WithCreateIntegerAttributeArray(v bool) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIntegerAttributeOption)  (*appwriteModel.AttributeInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer")

	options := &CreateIntegerAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["min"] = options.Min
	params["max"] = options.Max
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeInteger
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateIntegerAttribute update an integer attribute. Changing the `default`
// value will not update already existing documents.
// 
type UpdateIntegerAttributeOptions struct {
}

type UpdateIntegerAttributeOption func(*UpdateIntegerAttributeOptions)


													
func (srv *Databases) UpdateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min int, Max int, Default int)  (*appwriteModel.AttributeInteger, error) {
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
	var parsed appwriteModel.AttributeInteger
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateIpAttribute create IP address attribute.
// 
type CreateIpAttributeOptions struct {
	Default string
	Array bool
}

type CreateIpAttributeOption func(*CreateIpAttributeOptions)

func WithCreateIpAttributeDefault(v string) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Default = v
	}
}
func WithCreateIpAttributeArray(v bool) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIpAttributeOption)  (*appwriteModel.AttributeIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip")

	options := &CreateIpAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeIp
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateIpAttribute update an ip attribute. Changing the `default` value will
// not update already existing documents.
// 
type UpdateIpAttributeOptions struct {
}

type UpdateIpAttributeOption func(*UpdateIpAttributeOptions)


									
func (srv *Databases) UpdateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string)  (*appwriteModel.AttributeIp, error) {
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
	var parsed appwriteModel.AttributeIp
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateRelationshipAttribute create relationship attribute. [Learn more
// about relationship
// attributes](docs/databases-relationships#relationship-attributes).
// 
type CreateRelationshipAttributeOptions struct {
	TwoWay bool
	Key string
	TwoWayKey string
	OnDelete string
}

type CreateRelationshipAttributeOption func(*CreateRelationshipAttributeOptions)

func WithCreateRelationshipAttributeTwoWay(v bool) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWay = v
	}
}
func WithCreateRelationshipAttributeKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.Key = v
	}
}
func WithCreateRelationshipAttributeTwoWayKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWayKey = v
	}
}
func WithCreateRelationshipAttributeOnDelete(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.OnDelete = v
	}
}

									
func (srv *Databases) CreateRelationshipAttribute(DatabaseId string, CollectionId string, RelatedCollectionId string, Type string, optionalSetters ...CreateRelationshipAttributeOption)  (*appwriteModel.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/relationship")

	options := &CreateRelationshipAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["relatedCollectionId"] = RelatedCollectionId
	params["type"] = Type
	params["twoWay"] = options.TwoWay
	params["key"] = options.Key
	params["twoWayKey"] = options.TwoWayKey
	params["onDelete"] = options.OnDelete
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeRelationship
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateStringAttribute create a string attribute.
// 
type CreateStringAttributeOptions struct {
	Default string
	Array bool
}

type CreateStringAttributeOption func(*CreateStringAttributeOptions)

func WithCreateStringAttributeDefault(v string) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Default = v
	}
}
func WithCreateStringAttributeArray(v bool) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Array = v
	}
}

											
func (srv *Databases) CreateStringAttribute(DatabaseId string, CollectionId string, Key string, Size int, Required bool, optionalSetters ...CreateStringAttributeOption)  (*appwriteModel.AttributeString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string")

	options := &CreateStringAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["size"] = Size
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeString
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateStringAttribute update a string attribute. Changing the `default`
// value will not update already existing documents.
// 
type UpdateStringAttributeOptions struct {
}

type UpdateStringAttributeOption func(*UpdateStringAttributeOptions)


									
func (srv *Databases) UpdateStringAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string)  (*appwriteModel.AttributeString, error) {
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
	var parsed appwriteModel.AttributeString
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateUrlAttribute create a URL attribute.
// 
type CreateUrlAttributeOptions struct {
	Default string
	Array bool
}

type CreateUrlAttributeOption func(*CreateUrlAttributeOptions)

func WithCreateUrlAttributeDefault(v string) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Default = v
	}
}
func WithCreateUrlAttributeArray(v bool) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Array = v
	}
}

									
func (srv *Databases) CreateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateUrlAttributeOption)  (*appwriteModel.AttributeUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url")

	options := &CreateUrlAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = options.Default
	params["array"] = options.Array
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeUrl
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateUrlAttribute update an url attribute. Changing the `default` value
// will not update already existing documents.
// 
type UpdateUrlAttributeOptions struct {
}

type UpdateUrlAttributeOption func(*UpdateUrlAttributeOptions)


									
func (srv *Databases) UpdateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string)  (*appwriteModel.AttributeUrl, error) {
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
	var parsed appwriteModel.AttributeUrl
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetAttribute
type GetAttributeOptions struct {
}

type GetAttributeOption func(*GetAttributeOptions)


					
func (srv *Databases) GetAttribute(DatabaseId string, CollectionId string, Key string)  (*interface{}, error) {
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
type DeleteAttributeOptions struct {
}

type DeleteAttributeOption func(*DeleteAttributeOptions)


					
func (srv *Databases) DeleteAttribute(DatabaseId string, CollectionId string, Key string)  (*interface{}, error) {
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
// UpdateRelationshipAttribute update relationship attribute. [Learn more
// about relationship
// attributes](docs/databases-relationships#relationship-attributes).
// 
type UpdateRelationshipAttributeOptions struct {
	OnDelete string
}

type UpdateRelationshipAttributeOption func(*UpdateRelationshipAttributeOptions)

func WithUpdateRelationshipAttributeOnDelete(v string) UpdateRelationshipAttributeOption {
	return func(o *UpdateRelationshipAttributeOptions) {
		o.OnDelete = v
	}
}

							
func (srv *Databases) UpdateRelationshipAttribute(DatabaseId string, CollectionId string, Key string, optionalSetters ...UpdateRelationshipAttributeOption)  (*appwriteModel.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}/relationship")

	options := &UpdateRelationshipAttributeOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["onDelete"] = options.OnDelete
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.AttributeRelationship
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
type ListDocumentsOptions struct {
	Queries []interface{}
}

type ListDocumentsOption func(*ListDocumentsOptions)

func WithListDocumentsQueries(v []interface{}) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Queries = v
	}
}

					
func (srv *Databases) ListDocuments(DatabaseId string, CollectionId string, optionalSetters ...ListDocumentsOption)  (*appwriteModel.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	options := &ListDocumentsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["queries"] = options.Queries
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.DocumentList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
type CreateDocumentOptions struct {
	Permissions []interface{}
}

type CreateDocumentOption func(*CreateDocumentOptions)

func WithCreateDocumentPermissions(v []interface{}) CreateDocumentOption {
	return func(o *CreateDocumentOptions) {
		o.Permissions = v
	}
}

									
func (srv *Databases) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, optionalSetters ...CreateDocumentOption)  (*appwriteModel.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	options := &CreateDocumentOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["data"] = Data
	params["permissions"] = options.Permissions
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
type GetDocumentOptions struct {
	Queries []interface{}
}

type GetDocumentOption func(*GetDocumentOptions)

func WithGetDocumentQueries(v []interface{}) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.Queries = v
	}
}

							
func (srv *Databases) GetDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...GetDocumentOption)  (*appwriteModel.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	options := &GetDocumentOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["queries"] = options.Queries
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
type UpdateDocumentOptions struct {
	Data interface{}
	Permissions []interface{}
}

type UpdateDocumentOption func(*UpdateDocumentOptions)

func WithUpdateDocumentData(v interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Data = v
	}
}
func WithUpdateDocumentPermissions(v []interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Permissions = v
	}
}

							
func (srv *Databases) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...UpdateDocumentOption)  (*appwriteModel.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	options := &UpdateDocumentOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["data"] = options.Data
	params["permissions"] = options.Permissions
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Document
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteDocument delete a document by its unique ID.
type DeleteDocumentOptions struct {
}

type DeleteDocumentOption func(*DeleteDocumentOptions)


					
func (srv *Databases) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string)  (*interface{}, error) {
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
type ListIndexesOptions struct {
}

type ListIndexesOption func(*ListIndexesOptions)


			
func (srv *Databases) ListIndexes(DatabaseId string, CollectionId string)  (*appwriteModel.IndexList, error) {
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
	var parsed appwriteModel.IndexList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.IndexList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateIndex
type CreateIndexOptions struct {
	Orders []interface{}
}

type CreateIndexOption func(*CreateIndexOptions)

func WithCreateIndexOrders(v []interface{}) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Orders = v
	}
}

											
func (srv *Databases) CreateIndex(DatabaseId string, CollectionId string, Key string, Type string, Attributes []interface{}, optionalSetters ...CreateIndexOption)  (*appwriteModel.Index, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")

	options := &CreateIndexOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["type"] = Type
	params["attributes"] = Attributes
	params["orders"] = options.Orders
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Index
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetIndex
type GetIndexOptions struct {
}

type GetIndexOption func(*GetIndexOptions)


					
func (srv *Databases) GetIndex(DatabaseId string, CollectionId string, Key string)  (*appwriteModel.Index, error) {
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
	var parsed appwriteModel.Index
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteIndex
type DeleteIndexOptions struct {
}

type DeleteIndexOption func(*DeleteIndexOptions)


					
func (srv *Databases) DeleteIndex(DatabaseId string, CollectionId string, Key string)  (*interface{}, error) {
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

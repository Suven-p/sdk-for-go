package appwriteModel


// Collection Model
type Collection struct {
    // Collection ID.
    Id string `json:"$id"`
    // Collection creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Collection update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Collection permissions. [Learn more about permissions](/docs/permissions).
    Permissions []interface{} `json:"$permissions"`
    // Database ID.
    DatabaseId string `json:"databaseId"`
    // Collection name.
    Name string `json:"name"`
    // Collection enabled.
    Enabled bool `json:"enabled"`
    // Whether document-level permissions are enabled. [Learn more about
    // permissions](/docs/permissions).
    DocumentSecurity bool `json:"documentSecurity"`
    // Collection attributes.
    Attributes []interface{} `json:"attributes"`
    // Collection indexes.
    Indexes []interface{} `json:"indexes"`

}

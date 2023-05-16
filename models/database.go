package models


// Database Model
type Database struct {
    // Database ID.
    Id string `json:"$id"`
    // Database name.
    Name string `json:"name"`
    // Database creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`

}

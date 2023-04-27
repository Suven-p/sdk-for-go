package appwriteModel


// Token Model
type Token struct {
    // Token ID.
    Id string `json:"$id"`
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User ID.
    UserId string `json:"userId"`
    // Token secret key. This will return an empty string unless the response is
    // returned using an API key or as part of a webhook payload.
    Secret string `json:"secret"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`

}

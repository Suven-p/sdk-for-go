package appwriteModel


// Variable Model
type Variable struct {
    // Variable ID.
    Id string `json:"$id"`
    // Variable creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Variable creation date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Variable key.
    Key string `json:"key"`
    // Variable value.
    Value string `json:"value"`
    // Function ID.
    FunctionId string `json:"functionId"`

}

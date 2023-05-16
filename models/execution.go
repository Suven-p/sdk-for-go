package models


// Execution Model
type Execution struct {
    // Execution ID.
    Id string `json:"$id"`
    // Execution creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Execution upate date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Execution roles.
    Permissions []interface{} `json:"$permissions"`
    // Function ID.
    FunctionId string `json:"functionId"`
    // The trigger that caused the function to execute. Possible values can be:
    // `http`, `schedule`, or `event`.
    Trigger string `json:"trigger"`
    // The status of the function execution. Possible values can be: `waiting`,
    // `processing`, `completed`, or `failed`.
    Status string `json:"status"`
    // The script status code.
    StatusCode int `json:"statusCode"`
    // The script response output string. Logs the last 4,000 characters of the
    // execution response output.
    Response string `json:"response"`
    // The script stdout output string. Logs the last 4,000 characters of the
    // execution stdout output. This will return an empty string unless the
    // response is returned using an API key or as part of a webhook payload.
    Stdout string `json:"stdout"`
    // The script stderr output string. Logs the last 4,000 characters of the
    // execution stderr output. This will return an empty string unless the
    // response is returned using an API key or as part of a webhook payload.
    Stderr string `json:"stderr"`
    // The script execution duration in seconds.
    Duration float64 `json:"duration"`

}

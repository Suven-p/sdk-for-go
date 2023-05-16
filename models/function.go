package models


// Function Model
type Function struct {
    // Function ID.
    Id string `json:"$id"`
    // Function creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Function update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Execution permissions.
    Execute []interface{} `json:"execute"`
    // Function name.
    Name string `json:"name"`
    // Function enabled.
    Enabled bool `json:"enabled"`
    // Function execution runtime.
    Runtime string `json:"runtime"`
    // Function's active deployment ID.
    Deployment string `json:"deployment"`
    // Function variables.
    Vars []interface{} `json:"vars"`
    // Function trigger events.
    Events []interface{} `json:"events"`
    // Function execution schedult in CRON format.
    Schedule string `json:"schedule"`
    // Function's next scheduled execution time in ISO 8601 format.
    ScheduleNext string `json:"scheduleNext"`
    // Function's previous scheduled execution time in ISO 8601 format.
    SchedulePrevious string `json:"schedulePrevious"`
    // Function execution timeout in seconds.
    Timeout int `json:"timeout"`

}

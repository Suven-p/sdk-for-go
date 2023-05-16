package models


// User Model
type User struct {
    // User ID.
    Id string `json:"$id"`
    // User creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User name.
    Name string `json:"name"`
    // Hashed user password.
    Password string `json:"password"`
    // Password hashing algorithm.
    Hash string `json:"hash"`
    // Password hashing algorithm configuration.
    HashOptions interface{} `json:"hashOptions"`
    // User registration date in ISO 8601 format.
    Registration string `json:"registration"`
    // User status. Pass `true` for enabled and `false` for disabled.
    Status bool `json:"status"`
    // Password update time in ISO 8601 format.
    PasswordUpdate string `json:"passwordUpdate"`
    // User email address.
    Email string `json:"email"`
    // User phone number in E.164 format.
    Phone string `json:"phone"`
    // Email verification status.
    EmailVerification bool `json:"emailVerification"`
    // Phone verification status.
    PhoneVerification bool `json:"phoneVerification"`
    // User preferences as a key-value object
    Prefs interface{} `json:"prefs"`

}

package appwriteModel


// AttributeBoolean Model
type AttributeBoolean struct {
    // Attribute Key.
    Key string `json:"key"`
    // Attribute type.
    Type string `json:"xtype"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Is attribute required?
    Required bool `json:"required"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Default value for attribute when not provided. Cannot be set when attribute
    // is required.
    Default bool `json:"xdefault"`

}

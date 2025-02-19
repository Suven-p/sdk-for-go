package models

// AttributeFloat Model
type AttributeFloat struct {
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
	// Minimum value to enforce for new documents.
	Min float64 `json:"min"`
	// Maximum value to enforce for new documents.
	Max float64 `json:"max"`
	// Default value for attribute when not provided. Cannot be set when attribute
	// is required.
	Default float64 `json:"xdefault"`
}

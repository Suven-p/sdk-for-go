package models

// Index Model
type Index struct {
	// Index Key.
	Key string `json:"key"`
	// Index type.
	Type string `json:"xtype"`
	// Index status. Possible values: `available`, `processing`, `deleting`,
	// `stuck`, or `failed`
	Status string `json:"status"`
	// Index attributes.
	Attributes []interface{} `json:"attributes"`
	// Index orders.
	Orders []interface{} `json:"orders"`
}

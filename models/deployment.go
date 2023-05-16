package models

// Deployment Model
type Deployment struct {
	// Deployment ID.
	Id string `json:"$id"`
	// Deployment creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Deployment update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Resource ID.
	ResourceId string `json:"resourceId"`
	// Resource type.
	ResourceType string `json:"resourceType"`
	// The entrypoint file to use to execute the deployment code.
	Entrypoint string `json:"entrypoint"`
	// The code size in bytes.
	Size int `json:"size"`
	// The current build ID.
	BuildId string `json:"buildId"`
	// Whether the deployment should be automatically activated.
	Activate bool `json:"activate"`
	// The deployment status. Possible values are "processing", "building",
	// "pending", "ready", and "failed".
	Status string `json:"status"`
	// The build stdout.
	BuildStdout string `json:"buildStdout"`
	// The build stderr.
	BuildStderr string `json:"buildStderr"`
	// The current build time in seconds.
	BuildTime int `json:"buildTime"`
}

package functions

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Functions service
type Functions struct {
	client client.Client
}

func NewFunctions(clt client.Client) *Functions {
	return &Functions{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []interface{}
	Search         string
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListOption func(*ListOptions)

func WithListQueries(v []interface{}) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// List get a list of all the project's functions. You can use the query
// params to filter your results.
func (srv *Functions) List(optionalSetters ...ListOption) (*models.FunctionList, error) {
	path := "/functions"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.FunctionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.FunctionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Execute        []interface{}
	Events         []interface{}
	Schedule       string
	Timeout        int
	Enabled        bool
	enabledSetters map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Execute":  false,
		"Events":   false,
		"Schedule": false,
		"Timeout":  false,
		"Enabled":  false,
	}
	return &options
}

type CreateOption func(*CreateOptions)

func WithCreateExecute(v []interface{}) CreateOption {
	return func(o *CreateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func WithCreateEvents(v []interface{}) CreateOption {
	return func(o *CreateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func WithCreateSchedule(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func WithCreateTimeout(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// Create create a new function. You can pass a list of
// [permissions](/docs/permissions) to allow different project users or team
// with access to execute the function using the client API.
func (srv *Functions) Create(FunctionId string, Name string, Runtime string, optionalSetters ...CreateOption) (*models.Function, error) {
	path := "/functions"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	params["runtime"] = Runtime
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListRuntimes get a list of all runtimes that are currently active on your
// instance.
func (srv *Functions) ListRuntimes() (*models.RuntimeList, error) {
	path := "/functions/runtimes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.RuntimeList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.RuntimeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a function by its unique ID.
func (srv *Functions) Get(FunctionId string) (*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateOptions struct {
	Execute        []interface{}
	Events         []interface{}
	Schedule       string
	Timeout        int
	Enabled        bool
	enabledSetters map[string]bool
}

func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Execute":  false,
		"Events":   false,
		"Schedule": false,
		"Timeout":  false,
		"Enabled":  false,
	}
	return &options
}

type UpdateOption func(*UpdateOptions)

func WithUpdateExecute(v []interface{}) UpdateOption {
	return func(o *UpdateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func WithUpdateEvents(v []interface{}) UpdateOption {
	return func(o *UpdateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func WithUpdateSchedule(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func WithUpdateTimeout(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// Update update function by its unique ID.
func (srv *Functions) Update(FunctionId string, Name string, optionalSetters ...UpdateOption) (*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a function by its unique ID.
func (srv *Functions) Delete(FunctionId string) (*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListDeploymentsOptions struct {
	Queries        []interface{}
	Search         string
	enabledSetters map[string]bool
}

func (options ListDeploymentsOptions) New() *ListDeploymentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListDeploymentsOption func(*ListDeploymentsOptions)

func WithListDeploymentsQueries(v []interface{}) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListDeploymentsSearch(v string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// ListDeployments get a list of all the project's code deployments. You can
// use the query params to filter your results.
func (srv *Functions) ListDeployments(FunctionId string, optionalSetters ...ListDeploymentsOption) (*models.DeploymentList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")
	options := ListDeploymentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.DeploymentList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.DeploymentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateDeployment create a new function code deployment. Use this endpoint
// to upload a new version of your code function. To execute your newly
// uploaded code, you'll need to update the function's deployment to use your
// new deployment UID.
//
// This endpoint accepts a tar.gz file compressed with your code. Make sure to
// include any dependencies your code has within the compressed file. You can
// learn more about code packaging in the [Appwrite Cloud Functions
// tutorial](/docs/functions).
//
// Use the "command" param to set the entry point used to execute your code.
func (srv *Functions) CreateDeployment(FunctionId string, Entrypoint string, Code file.InputFile, Activate bool) (*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["entrypoint"] = Entrypoint
	params["code"] = Code
	params["activate"] = Activate
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}

	paramName := "code"

	uploadId := ""

	resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
	if err != nil {
		return nil, err
	}
	var parsed models.Deployment
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil
}

// GetDeployment get a code deployment by its unique ID.
func (srv *Functions) GetDeployment(FunctionId string, DeploymentId string) (*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Deployment
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateDeployment update the function code deployment ID using the unique
// function ID. Use this endpoint to switch the code deployment that should be
// executed by the execution endpoint.
func (srv *Functions) UpdateDeployment(FunctionId string, DeploymentId string) (*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteDeployment delete a code deployment by its unique ID.
func (srv *Functions) DeleteDeployment(FunctionId string, DeploymentId string) (*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateBuild
func (srv *Functions) CreateBuild(FunctionId string, DeploymentId string, BuildId string) (*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId, "{buildId}", BuildId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}/builds/{buildId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	params["buildId"] = BuildId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListExecutionsOptions struct {
	Queries        []interface{}
	Search         string
	enabledSetters map[string]bool
}

func (options ListExecutionsOptions) New() *ListExecutionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListExecutionsOption func(*ListExecutionsOptions)

func WithListExecutionsQueries(v []interface{}) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListExecutionsSearch(v string) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results.
func (srv *Functions) ListExecutions(FunctionId string, optionalSetters ...ListExecutionsOption) (*models.ExecutionList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")
	options := ListExecutionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.ExecutionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.ExecutionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateExecutionOptions struct {
	Data           string
	Async          bool
	enabledSetters map[string]bool
}

func (options CreateExecutionOptions) New() *CreateExecutionOptions {
	options.enabledSetters = map[string]bool{
		"Data":  false,
		"Async": false,
	}
	return &options
}

type CreateExecutionOption func(*CreateExecutionOptions)

func WithCreateExecutionData(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func WithCreateExecutionAsync(v bool) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Async = v
		o.enabledSetters["Async"] = true
	}
}

// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
func (srv *Functions) CreateExecution(FunctionId string, optionalSetters ...CreateExecutionOption) (*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")
	options := CreateExecutionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Async"] {
		params["async"] = options.Async
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Execution
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetExecution get a function execution log by its unique ID.
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string) (*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["executionId"] = ExecutionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Execution
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListVariables get a list of all variables of a specific function.
func (srv *Functions) ListVariables(FunctionId string) (*models.VariableList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.VariableList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.VariableList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateVariable create a new function variable. These variables can be
// accessed within function in the `env` object under the request variable.
func (srv *Functions) CreateVariable(FunctionId string, Key string, Value string) (*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["key"] = Key
	params["value"] = Value
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetVariable get a variable by its unique ID.
func (srv *Functions) GetVariable(FunctionId string, VariableId string) (*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateVariableOptions struct {
	Value          string
	enabledSetters map[string]bool
}

func (options UpdateVariableOptions) New() *UpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
	}
	return &options
}

type UpdateVariableOption func(*UpdateVariableOptions)

func WithUpdateVariableValue(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}

// UpdateVariable update variable by its unique ID.
func (srv *Functions) UpdateVariable(FunctionId string, VariableId string, Key string, optionalSetters ...UpdateVariableOption) (*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	options := UpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	params["key"] = Key
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteVariable delete a variable by its unique ID.
func (srv *Functions) DeleteVariable(FunctionId string, VariableId string) (*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

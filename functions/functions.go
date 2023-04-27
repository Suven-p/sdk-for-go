package functions

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Functions service
type Functions struct {
	client appwrite.Client
}

func NewFunctions(clt appwrite.Client) *Functions {
	return &Functions{
		client: clt,
	}
}

// List get a list of all the project's functions. You can use the query
// params to filter your results.
type ListOptions struct {
	Queries []interface{}
	Search string
}

type ListOption func(*ListOptions)

func WithListQueries(v []interface{}) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
	}
}
func WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
	}
}

	
func (srv *Functions) List(optionalSetters ...ListOption)  (*appwriteModel.FunctionList, error) {
	path := "/functions"

	options := &ListOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.FunctionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.FunctionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Create create a new function. You can pass a list of
// [permissions](/docs/permissions) to allow different project users or team
// with access to execute the function using the client API.
type CreateOptions struct {
	Execute []interface{}
	Events []interface{}
	Schedule string
	Timeout int
	Enabled bool
}

type CreateOption func(*CreateOptions)

func WithCreateExecute(v []interface{}) CreateOption {
	return func(o *CreateOptions) {
		o.Execute = v
	}
}
func WithCreateEvents(v []interface{}) CreateOption {
	return func(o *CreateOptions) {
		o.Events = v
	}
}
func WithCreateSchedule(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Schedule = v
	}
}
func WithCreateTimeout(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Timeout = v
	}
}
func WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
	}
}

							
func (srv *Functions) Create(FunctionId string, Name string, Runtime string, optionalSetters ...CreateOption)  (*appwriteModel.Function, error) {
	path := "/functions"

	options := &CreateOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	params["runtime"] = Runtime
	params["execute"] = options.Execute
	params["events"] = options.Events
	params["schedule"] = options.Schedule
	params["timeout"] = options.Timeout
	params["enabled"] = options.Enabled
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListRuntimes get a list of all runtimes that are currently active on your
// instance.
type ListRuntimesOptions struct {
}

type ListRuntimesOption func(*ListRuntimesOptions)



func (srv *Functions) ListRuntimes()  (*appwriteModel.RuntimeList, error) {
	path := "/functions/runtimes"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.RuntimeList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.RuntimeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Get get a function by its unique ID.
type GetOptions struct {
}

type GetOption func(*GetOptions)


	
func (srv *Functions) Get(FunctionId string)  (*appwriteModel.Function, error) {
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
	var parsed appwriteModel.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Update update function by its unique ID.
type UpdateOptions struct {
	Execute []interface{}
	Events []interface{}
	Schedule string
	Timeout int
	Enabled bool
}

type UpdateOption func(*UpdateOptions)

func WithUpdateExecute(v []interface{}) UpdateOption {
	return func(o *UpdateOptions) {
		o.Execute = v
	}
}
func WithUpdateEvents(v []interface{}) UpdateOption {
	return func(o *UpdateOptions) {
		o.Events = v
	}
}
func WithUpdateSchedule(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Schedule = v
	}
}
func WithUpdateTimeout(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Timeout = v
	}
}
func WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
	}
}

					
func (srv *Functions) Update(FunctionId string, Name string, optionalSetters ...UpdateOption)  (*appwriteModel.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	options := &UpdateOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	params["execute"] = options.Execute
	params["events"] = options.Events
	params["schedule"] = options.Schedule
	params["timeout"] = options.Timeout
	params["enabled"] = options.Enabled
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// Delete delete a function by its unique ID.
type DeleteOptions struct {
}

type DeleteOption func(*DeleteOptions)


	
func (srv *Functions) Delete(FunctionId string)  (*interface{}, error) {
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
// ListDeployments get a list of all the project's code deployments. You can
// use the query params to filter your results.
type ListDeploymentsOptions struct {
	Queries []interface{}
	Search string
}

type ListDeploymentsOption func(*ListDeploymentsOptions)

func WithListDeploymentsQueries(v []interface{}) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Queries = v
	}
}
func WithListDeploymentsSearch(v string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Search = v
	}
}

			
func (srv *Functions) ListDeployments(FunctionId string, optionalSetters ...ListDeploymentsOption)  (*appwriteModel.DeploymentList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")

	options := &ListDeploymentsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.DeploymentList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.DeploymentList)
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
type CreateDeploymentOptions struct {
}

type CreateDeploymentOption func(*CreateDeploymentOptions)


							
func (srv *Functions) CreateDeployment(FunctionId string, Entrypoint string, Code appwrite.InputFile, Activate bool)  (*appwriteModel.Deployment, error) {
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
	parsed, ok := resp.Result.(appwriteModel.Deployment)
	if !ok {
		return nil, errors.New("Unexpected response type")
	}
	return &parsed, nil

}
// GetDeployment get a code deployment by its unique ID.
type GetDeploymentOptions struct {
}

type GetDeploymentOption func(*GetDeploymentOptions)


			
func (srv *Functions) GetDeployment(FunctionId string, DeploymentId string)  (*appwriteModel.Deployment, error) {
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
	var parsed appwriteModel.Deployment
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateDeployment update the function code deployment ID using the unique
// function ID. Use this endpoint to switch the code deployment that should be
// executed by the execution endpoint.
type UpdateDeploymentOptions struct {
}

type UpdateDeploymentOption func(*UpdateDeploymentOptions)


			
func (srv *Functions) UpdateDeployment(FunctionId string, DeploymentId string)  (*appwriteModel.Function, error) {
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
	var parsed appwriteModel.Function
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteDeployment delete a code deployment by its unique ID.
type DeleteDeploymentOptions struct {
}

type DeleteDeploymentOption func(*DeleteDeploymentOptions)


			
func (srv *Functions) DeleteDeployment(FunctionId string, DeploymentId string)  (*interface{}, error) {
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
type CreateBuildOptions struct {
}

type CreateBuildOption func(*CreateBuildOptions)


					
func (srv *Functions) CreateBuild(FunctionId string, DeploymentId string, BuildId string)  (*interface{}, error) {
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
// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results.
type ListExecutionsOptions struct {
	Queries []interface{}
	Search string
}

type ListExecutionsOption func(*ListExecutionsOptions)

func WithListExecutionsQueries(v []interface{}) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Queries = v
	}
}
func WithListExecutionsSearch(v string) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Search = v
	}
}

			
func (srv *Functions) ListExecutions(FunctionId string, optionalSetters ...ListExecutionsOption)  (*appwriteModel.ExecutionList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

	options := &ListExecutionsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.ExecutionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.ExecutionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
type CreateExecutionOptions struct {
	Data string
	Async bool
}

type CreateExecutionOption func(*CreateExecutionOptions)

func WithCreateExecutionData(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Data = v
	}
}
func WithCreateExecutionAsync(v bool) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Async = v
	}
}

			
func (srv *Functions) CreateExecution(FunctionId string, optionalSetters ...CreateExecutionOption)  (*appwriteModel.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

	options := &CreateExecutionOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["data"] = options.Data
	params["async"] = options.Async
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Execution
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetExecution get a function execution log by its unique ID.
type GetExecutionOptions struct {
}

type GetExecutionOption func(*GetExecutionOptions)


			
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string)  (*appwriteModel.Execution, error) {
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
	var parsed appwriteModel.Execution
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListVariables get a list of all variables of a specific function.
type ListVariablesOptions struct {
}

type ListVariablesOption func(*ListVariablesOptions)


	
func (srv *Functions) ListVariables(FunctionId string)  (*appwriteModel.VariableList, error) {
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
	var parsed appwriteModel.VariableList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.VariableList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateVariable create a new function variable. These variables can be
// accessed within function in the `env` object under the request variable.
type CreateVariableOptions struct {
}

type CreateVariableOption func(*CreateVariableOptions)


					
func (srv *Functions) CreateVariable(FunctionId string, Key string, Value string)  (*appwriteModel.Variable, error) {
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
	var parsed appwriteModel.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetVariable get a variable by its unique ID.
type GetVariableOptions struct {
}

type GetVariableOption func(*GetVariableOptions)


			
func (srv *Functions) GetVariable(FunctionId string, VariableId string)  (*appwriteModel.Variable, error) {
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
	var parsed appwriteModel.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateVariable update variable by its unique ID.
type UpdateVariableOptions struct {
	Value string
}

type UpdateVariableOption func(*UpdateVariableOptions)

func WithUpdateVariableValue(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Value = v
	}
}

							
func (srv *Functions) UpdateVariable(FunctionId string, VariableId string, Key string, optionalSetters ...UpdateVariableOption)  (*appwriteModel.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")

	options := &UpdateVariableOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	params["key"] = Key
	params["value"] = options.Value
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Variable
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteVariable delete a variable by its unique ID.
type DeleteVariableOptions struct {
}

type DeleteVariableOption func(*DeleteVariableOptions)


			
func (srv *Functions) DeleteVariable(FunctionId string, VariableId string)  (*interface{}, error) {
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

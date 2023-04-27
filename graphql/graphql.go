package graphql

import (
	"encoding/json"
	"errors"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Graphql service
type Graphql struct {
	client appwrite.Client
}

func NewGraphql(clt appwrite.Client) *Graphql {
	return &Graphql{
		client: clt,
	}
}

// Query execute a GraphQL mutation.
type QueryOptions struct {
}

type QueryOption func(*QueryOptions)


	
func (srv *Graphql) Query(Query interface{})  (*interface{}, error) {
	path := "/graphql"

	params := map[string]interface{}{}
	params["query"] = Query
	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
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
// Mutation execute a GraphQL mutation.
type MutationOptions struct {
}

type MutationOption func(*MutationOptions)


	
func (srv *Graphql) Mutation(Query interface{})  (*interface{}, error) {
	path := "/graphql/mutation"

	params := map[string]interface{}{}
	params["query"] = Query
	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
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

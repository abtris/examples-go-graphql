package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestGraphQLReq(t *testing.T) {
	schema := GetSchema()
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	expectedJSON := `{"data":{"hello":"world"}}`
	if string(rJSON) != expectedJSON {
		t.Errorf("Expected %v and real %v)", expectedJSON, rJSON)
	}
}

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dominik-zeglen/ecoknow/mock"
	test "github.com/dominik-zeglen/ecoknow/testing"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource = mock.Adapter{}
var resolver = NewResolver(dataSource)
var schema = gql.MustParseSchema(Schema, &resolver)

var ErrNoError = fmt.Errorf("Did not return error")

func execQuery(query string, variables string) (string, error) {
	var vs map[string]interface{}

	err := json.Unmarshal([]byte(variables), &vs)
	if err != nil {
		return "", err
	}

	result := schema.Exec(context.TODO(), query, "", vs)
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	json.Indent(&out, jsonResult, "", "    ")
	return out.String(), nil
}

func resetDatabase() {
	dataSource.ResetMockDatabase(test.Directories, test.Templates, test.Pages)
}

package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"github.com/Khan/genql/graphql"
)

type QueryWithAliasResponse struct {
	User QueryWithAliasUser
}

type QueryWithAliasUser struct {
	ID string
}

func QueryWithAlias(
	client graphql.Client,
) (*QueryWithAliasResponse, error) {
	var retval QueryWithAliasResponse
	err := client.MakeRequest(
		nil,
		"QueryWithAlias",
		`
query QueryWithAlias {
	User: user {
		ID: id
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

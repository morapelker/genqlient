package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNoFragmentsQueryRandomItemArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRandomItemContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent()
}

func (v *InterfaceNoFragmentsQueryRandomItemArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}

func __unmarshalInterfaceNoFragmentsQueryRandomItemContent(v *InterfaceNoFragmentsQueryRandomItemContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNoFragmentsQueryRandomItemArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryRandomItemVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryRandomItemTopic)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNoFragmentsQueryRandomItemContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNoFragmentsQueryRandomItemTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryResponse is returned by InterfaceNoFragmentsQuery on success.
type InterfaceNoFragmentsQueryResponse struct {
	Root        InterfaceNoFragmentsQueryRootTopic           `json:"root"`
	RandomItem  InterfaceNoFragmentsQueryRandomItemContent   `json:"-"`
	WithPointer *InterfaceNoFragmentsQueryWithPointerContent `json:"-"`
}

func (v *InterfaceNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {

	type InterfaceNoFragmentsQueryResponseWrapper InterfaceNoFragmentsQueryResponse

	var firstPass struct {
		*InterfaceNoFragmentsQueryResponseWrapper
		RandomItem  json.RawMessage `json:"randomItem"`
		WithPointer json.RawMessage `json:"withPointer"`
	}
	firstPass.InterfaceNoFragmentsQueryResponseWrapper = (*InterfaceNoFragmentsQueryResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.RandomItem
		raw := firstPass.RandomItem
		err = __unmarshalInterfaceNoFragmentsQueryRandomItemContent(
			target, raw)
		if err != nil {
			return err
		}
	}
	{
		target := &v.WithPointer
		raw := firstPass.WithPointer
		*target = new(InterfaceNoFragmentsQueryWithPointerContent)
		err = __unmarshalInterfaceNoFragmentsQueryWithPointerContent(
			*target, raw)
		if err != nil {
			return err
		}
	}
	return nil
}

// InterfaceNoFragmentsQueryRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRootTopic struct {
	// ID is documented in the Content interface.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryWithPointerArticle struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent()
}

func (v *InterfaceNoFragmentsQueryWithPointerArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}
func (v *InterfaceNoFragmentsQueryWithPointerVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}
func (v *InterfaceNoFragmentsQueryWithPointerTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}

func __unmarshalInterfaceNoFragmentsQueryWithPointerContent(v *InterfaceNoFragmentsQueryWithPointerContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNoFragmentsQueryWithPointerArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryWithPointerVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryWithPointerTopic)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNoFragmentsQueryWithPointerContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNoFragmentsQueryWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryWithPointerTopic struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryWithPointerVideo struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

func InterfaceNoFragmentsQuery(
	client graphql.Client,
) (*InterfaceNoFragmentsQueryResponse, error) {
	var retval InterfaceNoFragmentsQueryResponse
	err := client.MakeRequest(
		nil,
		"InterfaceNoFragmentsQuery",
		`
query InterfaceNoFragmentsQuery {
	root {
		id
		name
	}
	randomItem {
		__typename
		id
		name
	}
	withPointer: randomItem {
		__typename
		id
		name
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

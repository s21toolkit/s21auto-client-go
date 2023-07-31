package gql

import (
	"errors"
	"regexp"
	"s21client/util"

	"github.com/go-resty/resty/v2"
)

type Request[TVariables any] struct {
	OperationName string     `json:"operationName"`
	Query         string     `json:"query"`
	Variables     TVariables `json:"variables"`
}

type Response[TData any] struct {
	Data TData `json:"data"`
}

// Empty struct to use as variables/data
type None struct{}

var operationNamePattern = regexp.MustCompile(`query\s+(?P<OperationName>\w+)`)

// Creates gql.Request from a graphql query
func NewQueryRequest[TVariables any](query string, optionalVariables ...TVariables) Request[TVariables] {
	match := operationNamePattern.FindStringSubmatch(query)

	if match == nil {
		panic(errors.New("unable to find operation name for a query"))
	}

	operationName := match[operationNamePattern.SubexpIndex("OperationName")]

	var variables TVariables

	if len(optionalVariables) > 0 {
		variables = optionalVariables[0]
	}

	return Request[TVariables]{
		OperationName: operationName,
		Query:         query,
		Variables:     variables,
	}
}

func ExtractResponseData[TData any](response *resty.Response) (data TData, err error) {
	body := response.Body()

	gqlData, err := util.UnmarshalJson[Response[TData]](body)

	if err != nil {
		return
	}

	data = gqlData.Data

	return
}

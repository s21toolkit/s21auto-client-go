package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetSearchHistory struct {
}


type Data_GetSearchHistory struct {
	Data_School21_GetSearchHistory Data_School21_GetSearchHistory `json:"school21"`
}

type Data_School21_GetSearchHistory struct {
	GetSearchHistoryTooltips []Data_GetSearchHistoryTooltip_GetSearchHistory `json:"getSearchHistoryTooltips"`
	Typename                 string                    `json:"__typename"`
}

type Data_GetSearchHistoryTooltip_GetSearchHistory struct {
	TooltipText     string `json:"tooltipText"`
	TooltipCategory string `json:"tooltipCategory"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetSearchHistory(variables Variables_GetSearchHistory) (Data_GetSearchHistory, error) {
	request := gql.NewQueryRequest[Variables_GetSearchHistory](
		"query getSearchHistory {\n  school21 {\n    getSearchHistoryTooltips {\n      tooltipText\n      tooltipCategory\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetSearchHistory](ctx, request)
}
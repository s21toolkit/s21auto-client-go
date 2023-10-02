package requests

import "github.com/s21toolkit/s21client/gql"

type GetSearchHistory_Variables struct {
}


type GetSearchHistory_Data struct {
	School21 GetSearchHistory_Data_School21 `json:"school21"`
}

type GetSearchHistory_Data_School21 struct {
	GetSearchHistoryTooltips []GetSearchHistory_Data_GetSearchHistoryTooltip `json:"getSearchHistoryTooltips"`
	Typename                 string                    `json:"__typename"`
}

type GetSearchHistory_Data_GetSearchHistoryTooltip struct {
	TooltipText     string `json:"tooltipText"`
	TooltipCategory string `json:"tooltipCategory"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetSearchHistory(variables GetSearchHistory_Variables) (GetSearchHistory_Data, error) {
	request := gql.NewQueryRequest[GetSearchHistory_Variables](
		"query getSearchHistory {\n  school21 {\n    getSearchHistoryTooltips {\n      tooltipText\n      tooltipCategory\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetSearchHistory_Data](ctx, request)
}
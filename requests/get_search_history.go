package requests

import "s21client/gql"

type Variables_GetSearchHistory struct {
}


type Data_GetSearchHistory struct {
	School21_GetSearchHistory School21_GetSearchHistory `json:"school21"`
}

type School21_GetSearchHistory struct {
	GetSearchHistoryTooltips []GetSearchHistoryTooltip_GetSearchHistory `json:"getSearchHistoryTooltips"`
	Typename                 string                    `json:"__typename"`
}

type GetSearchHistoryTooltip_GetSearchHistory struct {
	TooltipText     string `json:"tooltipText"`
	TooltipCategory string `json:"tooltipCategory"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetSearchHistory(variables Variables_GetSearchHistory) (Data_GetSearchHistory, error) {
	request := gql.NewQueryRequest[Variables_GetSearchHistory](
		"query getSearchHistory {   school21 {     getSearchHistoryTooltips {       tooltipText       tooltipCategory       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetSearchHistory](ctx, request)
}
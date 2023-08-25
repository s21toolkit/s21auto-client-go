package requests

import "s21client/gql"

type Request_Variables_GetSearchHistory struct {
}


type Response_Data_GetSearchHistory struct {
	Response_School21_GetSearchHistory Response_School21_GetSearchHistory `json:"school21"`
}

type Response_School21_GetSearchHistory struct {
	GetSearchHistoryTooltips []Response_GetSearchHistoryTooltip_GetSearchHistory `json:"getSearchHistoryTooltips"`
	Typename                 string                    `json:"__typename"`
}

type Response_GetSearchHistoryTooltip_GetSearchHistory struct {
	TooltipText     string `json:"tooltipText"`
	TooltipCategory string `json:"tooltipCategory"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetSearchHistory(variables Request_Variables_GetSearchHistory) (Response_Data_GetSearchHistory, error) {
	request := gql.NewQueryRequest[Request_Variables_GetSearchHistory](
		"query getSearchHistory {   school21 {     getSearchHistoryTooltips {       tooltipText       tooltipCategory       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetSearchHistory](ctx, request)
}
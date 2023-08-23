package requests

import "s21client/gql"

type Variables_GetSearchHistory struct {
}


type Data_GetSearchHistory struct {
	School21_GetSearchHistory School21_GetSearchHistory `json:"school21"`
}

type School21_GetSearchHistory struct {
	GetSearchHistoryTooltips []GetSearchHistoryTooltip_GetSearchHistory `json:"getSearchHistoryTooltips"`
	Typename_GetSearchHistory                 string                    `json:"__typename"`
}

type GetSearchHistoryTooltip_GetSearchHistory struct {
	TooltipText     string          `json:"tooltipText"`
	TooltipCategory_GetSearchHistory TooltipCategory_GetSearchHistory `json:"tooltipCategory"`
	Typename_GetSearchHistory        Typename_GetSearchHistory        `json:"__typename"`
}

type TooltipCategory_GetSearchHistory string

const (
	Profiles TooltipCategory_GetSearchHistory = "PROFILES"
)

type Typename_GetSearchHistory string

const (
	SearchTooltip Typename_GetSearchHistory = "SearchTooltip"
)

func (ctx *RequestContext) GetSearchHistory(variables Variables_GetSearchHistory) (Data_GetSearchHistory, error) {
	request := gql.NewQueryRequest[Variables_GetSearchHistory](
		"query getSearchHistory {   school21 {     getSearchHistoryTooltips {       tooltipText       tooltipCategory       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetSearchHistory](ctx, request)
}
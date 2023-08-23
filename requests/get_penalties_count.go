package requests

import "s21client/gql"

type Variables_GetPenaltiesCount struct {
	Statuses []string `json:"statuses"`
}


type Data_GetPenaltiesCount struct {
	School21_GetPenaltiesCount School21_GetPenaltiesCount `json:"school21"`
}

type School21_GetPenaltiesCount struct {
	CountMyPenalties int64  `json:"countMyPenalties"`
	Typename         string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltiesCount(variables Variables_GetPenaltiesCount) (Data_GetPenaltiesCount, error) {
	request := gql.NewQueryRequest[Variables_GetPenaltiesCount](
		"query getPenaltiesCount($statuses: [String]!) {   school21 {     countMyPenalties(statuses: $statuses)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetPenaltiesCount](ctx, request)
}
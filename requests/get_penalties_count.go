package requests

import "s21client/gql"

type Request_Variables_GetPenaltiesCount struct {
	Statuses []string `json:"statuses"`
}


type Response_Data_GetPenaltiesCount struct {
	Response_School21_GetPenaltiesCount Response_School21_GetPenaltiesCount `json:"school21"`
}

type Response_School21_GetPenaltiesCount struct {
	CountMyPenalties int64  `json:"countMyPenalties"`
	Typename         string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltiesCount(variables Request_Variables_GetPenaltiesCount) (Response_Data_GetPenaltiesCount, error) {
	request := gql.NewQueryRequest[Request_Variables_GetPenaltiesCount](
		"query getPenaltiesCount($statuses: [String]!) {   school21 {     countMyPenalties(statuses: $statuses)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetPenaltiesCount](ctx, request)
}
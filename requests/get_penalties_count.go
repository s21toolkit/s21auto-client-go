package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetPenaltiesCount struct {
	Statuses []string `json:"statuses"`
}


type Data_GetPenaltiesCount struct {
	Data_School21_GetPenaltiesCount Data_School21_GetPenaltiesCount `json:"school21"`
}

type Data_School21_GetPenaltiesCount struct {
	CountMyPenalties int64  `json:"countMyPenalties"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) GetPenaltiesCount(variables Variables_GetPenaltiesCount) (Data_GetPenaltiesCount, error) {
	request := gql.NewQueryRequest[Variables_GetPenaltiesCount](
		"query getPenaltiesCount($statuses: [String]!) {\n  school21 {\n    countMyPenalties(statuses: $statuses)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetPenaltiesCount](ctx, request)
}
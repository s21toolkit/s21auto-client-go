package requests

import "github.com/s21toolkit/s21client/gql"

type GetPenaltiesCount_Variables struct {
	Statuses []string `json:"statuses"`
}


type GetPenaltiesCount_Data struct {
	School21 GetPenaltiesCount_Data_School21 `json:"school21"`
}

type GetPenaltiesCount_Data_School21 struct {
	CountMyPenalties int64  `json:"countMyPenalties"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) GetPenaltiesCount(variables GetPenaltiesCount_Variables) (GetPenaltiesCount_Data, error) {
	request := gql.NewQueryRequest[GetPenaltiesCount_Variables](
		"query getPenaltiesCount($statuses: [String]!) {\n  school21 {\n    countMyPenalties(statuses: $statuses)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetPenaltiesCount_Data](ctx, request)
}
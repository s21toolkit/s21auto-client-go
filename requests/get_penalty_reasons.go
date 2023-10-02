package requests

import "github.com/s21toolkit/s21client/gql"

type GetPenaltyReasons_Variables struct {
}


type GetPenaltyReasons_Data struct {
	School21 GetPenaltyReasons_Data_School21 `json:"school21"`
}

type GetPenaltyReasons_Data_School21 struct {
	GetPenaltyReasons []GetPenaltyReasons_Data_GetPenaltyReason `json:"getPenaltyReasons"`
	Typename          string             `json:"__typename"`
}

type GetPenaltyReasons_Data_GetPenaltyReason struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetPenaltyReasons(variables GetPenaltyReasons_Variables) (GetPenaltyReasons_Data, error) {
	request := gql.NewQueryRequest[GetPenaltyReasons_Variables](
		"query getPenaltyReasons {\n  school21 {\n    getPenaltyReasons {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetPenaltyReasons_Data](ctx, request)
}
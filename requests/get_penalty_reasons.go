package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetPenaltyReasons struct {
}


type Data_GetPenaltyReasons struct {
	Data_School21_GetPenaltyReasons Data_School21_GetPenaltyReasons `json:"school21"`
}

type Data_School21_GetPenaltyReasons struct {
	GetPenaltyReasons []Data_GetPenaltyReason_GetPenaltyReasons `json:"getPenaltyReasons"`
	Typename          string             `json:"__typename"`
}

type Data_GetPenaltyReason_GetPenaltyReasons struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetPenaltyReasons(variables Variables_GetPenaltyReasons) (Data_GetPenaltyReasons, error) {
	request := gql.NewQueryRequest[Variables_GetPenaltyReasons](
		"query getPenaltyReasons {\n  school21 {\n    getPenaltyReasons {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetPenaltyReasons](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetActivityTypes struct {
}


type Data_GetActivityTypes struct {
	Data_School21_GetActivityTypes Data_School21_GetActivityTypes `json:"school21"`
}

type Data_School21_GetActivityTypes struct {
	GetActivityTypes []Data_GetActivityType_GetActivityTypes `json:"getActivityTypes"`
	Typename         string            `json:"__typename"`
}

type Data_GetActivityType_GetActivityTypes struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) GetActivityTypes(variables Variables_GetActivityTypes) (Data_GetActivityTypes, error) {
	request := gql.NewQueryRequest[Variables_GetActivityTypes](
		"query getActivityTypes {\n  school21 {\n    getActivityTypes {\n      id\n      description\n      category\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetActivityTypes](ctx, request)
}
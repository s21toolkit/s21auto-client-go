package requests

import "github.com/s21toolkit/s21client/gql"

type GetActivityTypes_Variables struct {
}


type GetActivityTypes_Data struct {
	School21 GetActivityTypes_Data_School21 `json:"school21"`
}

type GetActivityTypes_Data_School21 struct {
	GetActivityTypes []GetActivityTypes_Data_GetActivityType `json:"getActivityTypes"`
	Typename         string            `json:"__typename"`
}

type GetActivityTypes_Data_GetActivityType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) GetActivityTypes(variables GetActivityTypes_Variables) (GetActivityTypes_Data, error) {
	request := gql.NewQueryRequest[GetActivityTypes_Variables](
		"query getActivityTypes {\n  school21 {\n    getActivityTypes {\n      id\n      description\n      category\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetActivityTypes_Data](ctx, request)
}
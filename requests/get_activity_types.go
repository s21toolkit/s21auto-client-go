package requests

import "s21client/gql"

type Variables_GetActivityTypes struct {
}


type Data_GetActivityTypes struct {
	School21_GetActivityTypes School21_GetActivityTypes `json:"school21"`
}

type School21_GetActivityTypes struct {
	GetActivityTypes []GetActivityType_GetActivityTypes `json:"getActivityTypes"`
	Typename         string            `json:"__typename"`
}

type GetActivityType_GetActivityTypes struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) GetActivityTypes(variables Variables_GetActivityTypes) (Data_GetActivityTypes, error) {
	request := gql.NewQueryRequest[Variables_GetActivityTypes](
		"query getActivityTypes {   school21 {     getActivityTypes {       id       description       category       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetActivityTypes](ctx, request)
}
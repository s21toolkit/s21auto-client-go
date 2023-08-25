package requests

import "s21client/gql"

type Request_Variables_GetActivityTypes struct {
}


type Response_Data_GetActivityTypes struct {
	Response_School21_GetActivityTypes Response_School21_GetActivityTypes `json:"school21"`
}

type Response_School21_GetActivityTypes struct {
	GetActivityTypes []Response_GetActivityType_GetActivityTypes `json:"getActivityTypes"`
	Typename         string            `json:"__typename"`
}

type Response_GetActivityType_GetActivityTypes struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) GetActivityTypes(variables Request_Variables_GetActivityTypes) (Response_Data_GetActivityTypes, error) {
	request := gql.NewQueryRequest[Request_Variables_GetActivityTypes](
		"query getActivityTypes {   school21 {     getActivityTypes {       id       description       category       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetActivityTypes](ctx, request)
}
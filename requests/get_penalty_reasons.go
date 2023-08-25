package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetPenaltyReasons struct {
}


type Response_Data_GetPenaltyReasons struct {
	Response_School21_GetPenaltyReasons Response_School21_GetPenaltyReasons `json:"school21"`
}

type Response_School21_GetPenaltyReasons struct {
	GetPenaltyReasons []Response_GetPenaltyReason_GetPenaltyReasons `json:"getPenaltyReasons"`
	Typename          string             `json:"__typename"`
}

type Response_GetPenaltyReason_GetPenaltyReasons struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltyReasons(variables Request_Variables_GetPenaltyReasons) (Response_Data_GetPenaltyReasons, error) {
	request := gql.NewQueryRequest[Request_Variables_GetPenaltyReasons](
		"query getPenaltyReasons {   school21 {     getPenaltyReasons {       id       name       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetPenaltyReasons](ctx, request)
}
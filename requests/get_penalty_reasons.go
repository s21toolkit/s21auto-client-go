package requests

import "s21client/gql"

type Variables_GetPenaltyReasons struct {
}


type Data_GetPenaltyReasons struct {
	School21_GetPenaltyReasons School21_GetPenaltyReasons `json:"school21"`
}

type School21_GetPenaltyReasons struct {
	GetPenaltyReasons []GetPenaltyReason_GetPenaltyReasons `json:"getPenaltyReasons"`
	Typename          string             `json:"__typename"`
}

type GetPenaltyReason_GetPenaltyReasons struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltyReasons(variables Variables_GetPenaltyReasons) (Data_GetPenaltyReasons, error) {
	request := gql.NewQueryRequest[Variables_GetPenaltyReasons](
		"query getPenaltyReasons {   school21 {     getPenaltyReasons {       id       name       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetPenaltyReasons](ctx, request)
}
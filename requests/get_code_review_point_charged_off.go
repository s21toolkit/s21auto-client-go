package requests

import "s21client/gql"

type Request_Variables_GetCodeReviewPointChargedOff struct {
	GoalID string `json:"goalId"`
}


type Response_Data_GetCodeReviewPointChargedOff struct {
	Response_Student_GetCodeReviewPointChargedOff Response_Student_GetCodeReviewPointChargedOff `json:"student"`
}

type Response_Student_GetCodeReviewPointChargedOff struct {
	GetCodeReviewPointChargedOff bool   `json:"getCodeReviewPointChargedOff"`
	Typename                     string `json:"__typename"`
}

func (ctx *RequestContext) GetCodeReviewPointChargedOff(variables Request_Variables_GetCodeReviewPointChargedOff) (Response_Data_GetCodeReviewPointChargedOff, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCodeReviewPointChargedOff](
		"query getCodeReviewPointChargedOff($goalId: ID!) {   student {     getCodeReviewPointChargedOff(goalId: $goalId)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetCodeReviewPointChargedOff](ctx, request)
}
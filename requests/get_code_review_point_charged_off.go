package requests

import "s21client/gql"

type Variables_GetCodeReviewPointChargedOff struct {
	GoalID string `json:"goalId"`
}


type Data_GetCodeReviewPointChargedOff struct {
	Student_GetCodeReviewPointChargedOff Student_GetCodeReviewPointChargedOff `json:"student"`
}

type Student_GetCodeReviewPointChargedOff struct {
	GetCodeReviewPointChargedOff bool   `json:"getCodeReviewPointChargedOff"`
	Typename                     string `json:"__typename"`
}

func (ctx *RequestContext) GetCodeReviewPointChargedOff(variables Variables_GetCodeReviewPointChargedOff) (Data_GetCodeReviewPointChargedOff, error) {
	request := gql.NewQueryRequest[Variables_GetCodeReviewPointChargedOff](
		"query getCodeReviewPointChargedOff($goalId: ID!) {   student {     getCodeReviewPointChargedOff(goalId: $goalId)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewPointChargedOff](ctx, request)
}
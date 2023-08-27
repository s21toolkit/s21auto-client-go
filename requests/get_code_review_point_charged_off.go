package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCodeReviewPointChargedOff struct {
	GoalID string `json:"goalId"`
}


type Data_GetCodeReviewPointChargedOff struct {
	Data_Student_GetCodeReviewPointChargedOff Data_Student_GetCodeReviewPointChargedOff `json:"student"`
}

type Data_Student_GetCodeReviewPointChargedOff struct {
	GetCodeReviewPointChargedOff bool   `json:"getCodeReviewPointChargedOff"`
	Typename                     string `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewPointChargedOff(variables Variables_GetCodeReviewPointChargedOff) (Data_GetCodeReviewPointChargedOff, error) {
	request := gql.NewQueryRequest[Variables_GetCodeReviewPointChargedOff](
		"query getCodeReviewPointChargedOff($goalId: ID!) {\n  student {\n    getCodeReviewPointChargedOff(goalId: $goalId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewPointChargedOff](ctx, request)
}
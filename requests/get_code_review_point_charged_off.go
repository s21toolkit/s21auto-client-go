package requests

import "github.com/s21toolkit/s21client/gql"

type GetCodeReviewPointChargedOff_Variables struct {
	GoalID string `json:"goalId"`
}


type GetCodeReviewPointChargedOff_Data struct {
	Student GetCodeReviewPointChargedOff_Data_Student `json:"student"`
}

type GetCodeReviewPointChargedOff_Data_Student struct {
	GetCodeReviewPointChargedOff bool   `json:"getCodeReviewPointChargedOff"`
	Typename                     string `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewPointChargedOff(variables GetCodeReviewPointChargedOff_Variables) (GetCodeReviewPointChargedOff_Data, error) {
	request := gql.NewQueryRequest[GetCodeReviewPointChargedOff_Variables](
		"query getCodeReviewPointChargedOff($goalId: ID!) {\n  student {\n    getCodeReviewPointChargedOff(goalId: $goalId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCodeReviewPointChargedOff_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type GetCodeReviewMyStudent_Variables struct {
	StudentGoalID string `json:"studentGoalId"`
}


type GetCodeReviewMyStudent_Data struct {
	Student GetCodeReviewMyStudent_Data_Student `json:"student"`
}

type GetCodeReviewMyStudent_Data_Student struct {
	GetMyStudentCodeReview *GetCodeReviewMyStudent_Data_GetMyStudentCodeReview `json:"getMyStudentCodeReview"`
	Typename               string                  `json:"__typename"`
}

type GetCodeReviewMyStudent_Data_GetMyStudentCodeReview struct {
	ReviewerCommentsCount int64             `json:"reviewerCommentsCount"`
	CodeReviewRounds      []GetCodeReviewMyStudent_Data_CodeReviewRound `json:"codeReviewRounds"`
	Typename              string            `json:"__typename"`
}

type GetCodeReviewMyStudent_Data_CodeReviewRound struct {
	EventID             *string `json:"eventId"`
	CodeReviewRoundType string  `json:"codeReviewRoundType"`
	CodeReviewStatus    string  `json:"codeReviewStatus"`
	StartTime           string  `json:"startTime"`
	EndTime             string  `json:"endTime"`
	MergeRequestURL     string  `json:"mergeRequestURL"`
	CreateTime          string  `json:"createTime"`
	Typename            string  `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewMyStudent(variables GetCodeReviewMyStudent_Variables) (GetCodeReviewMyStudent_Data, error) {
	request := gql.NewQueryRequest[GetCodeReviewMyStudent_Variables](
		"query getCodeReviewMyStudent($studentGoalId: ID!) {\n  student {\n    getMyStudentCodeReview(studentGoalId: $studentGoalId) {\n      reviewerCommentsCount\n      codeReviewRounds {\n        ...CodeReviewRound\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewRound on CodeReviewRound {\n  eventId\n  codeReviewRoundType\n  codeReviewStatus\n  startTime\n  endTime\n  mergeRequestURL\n  createTime\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetCodeReviewMyStudent_Data](ctx, request)
}
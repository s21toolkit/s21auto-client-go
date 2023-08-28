package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCodeReviewMyStudent struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_GetCodeReviewMyStudent struct {
	Data_Student_GetCodeReviewMyStudent Data_Student_GetCodeReviewMyStudent `json:"student"`
}

type Data_Student_GetCodeReviewMyStudent struct {
	Data_GetMyStudentCodeReview_GetCodeReviewMyStudent *Data_GetMyStudentCodeReview_GetCodeReviewMyStudent `json:"getMyStudentCodeReview"`
	Typename               string                  `json:"__typename"`
}

type Data_GetMyStudentCodeReview_GetCodeReviewMyStudent struct {
	ReviewerCommentsCount int64             `json:"reviewerCommentsCount"`
	CodeReviewRounds      []Data_CodeReviewRound_GetCodeReviewMyStudent `json:"codeReviewRounds"`
	Typename              string            `json:"__typename"`
}

type Data_CodeReviewRound_GetCodeReviewMyStudent struct {
	EventID             *string `json:"eventId"`
	CodeReviewRoundType string  `json:"codeReviewRoundType"`
	CodeReviewStatus    string  `json:"codeReviewStatus"`
	StartTime           string  `json:"startTime"`
	EndTime             string  `json:"endTime"`
	MergeRequestURL     string  `json:"mergeRequestURL"`
	CreateTime          string  `json:"createTime"`
	Typename            string  `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewMyStudent(variables Variables_GetCodeReviewMyStudent) (Data_GetCodeReviewMyStudent, error) {
	request := gql.NewQueryRequest[Variables_GetCodeReviewMyStudent](
		"query getCodeReviewMyStudent($studentGoalId: ID!) {\n  student {\n    getMyStudentCodeReview(studentGoalId: $studentGoalId) {\n      reviewerCommentsCount\n      codeReviewRounds {\n        ...CodeReviewRound\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewRound on CodeReviewRound {\n  eventId\n  codeReviewRoundType\n  codeReviewStatus\n  startTime\n  endTime\n  mergeRequestURL\n  createTime\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewMyStudent](ctx, request)
}
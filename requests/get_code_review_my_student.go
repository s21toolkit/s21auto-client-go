package requests

import "s21client/gql"

type Variables_GetCodeReviewMyStudent struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_GetCodeReviewMyStudent struct {
	Student_GetCodeReviewMyStudent Student_GetCodeReviewMyStudent `json:"student"`
}

type Student_GetCodeReviewMyStudent struct {
	GetMyStudentCodeReview_GetCodeReviewMyStudent *GetMyStudentCodeReview_GetCodeReviewMyStudent `json:"getMyStudentCodeReview"`
	Typename               string                  `json:"__typename"`
}

type GetMyStudentCodeReview_GetCodeReviewMyStudent struct {
	ReviewerCommentsCount int64             `json:"reviewerCommentsCount"`
	CodeReviewRounds      []CodeReviewRound_GetCodeReviewMyStudent `json:"codeReviewRounds"`
	Typename              string            `json:"__typename"`
}

type CodeReviewRound_GetCodeReviewMyStudent struct {
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
		"query getCodeReviewMyStudent($studentGoalId: ID!) {   student {     getMyStudentCodeReview(studentGoalId: $studentGoalId) {       reviewerCommentsCount       codeReviewRounds {         ...CodeReviewRound         __typename       }       __typename     }     __typename   } }  fragment CodeReviewRound on CodeReviewRound {   eventId   codeReviewRoundType   codeReviewStatus   startTime   endTime   mergeRequestURL   createTime   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewMyStudent](ctx, request)
}
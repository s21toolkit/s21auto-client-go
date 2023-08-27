package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CreateFeedbackOnEvaluation struct {
	Variables_ReviewFeedbackInput_CreateFeedbackOnEvaluation Variables_ReviewFeedbackInput_CreateFeedbackOnEvaluation `json:"reviewFeedbackInput"`
}

type Variables_ReviewFeedbackInput_CreateFeedbackOnEvaluation struct {
	FilledChecklistID            string                        `json:"filledChecklistId"`
	Comment                      string                        `json:"comment"`
	ReviewFeedbackCategoryValues []Variables_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation `json:"reviewFeedbackCategoryValues"`
}

type Variables_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation struct {
	FeedbackCategory      string `json:"feedbackCategory"`
	FeedbackCategoryValue string `json:"feedbackCategoryValue"`
}


type Data_CreateFeedbackOnEvaluation struct {
	Data_Student_CreateFeedbackOnEvaluation Data_Student_CreateFeedbackOnEvaluation `json:"student"`
}

type Data_Student_CreateFeedbackOnEvaluation struct {
	Data_CreateReviewFeedback_CreateFeedbackOnEvaluation Data_CreateReviewFeedback_CreateFeedbackOnEvaluation `json:"createReviewFeedback"`
	Typename             string               `json:"__typename"`
}

type Data_CreateReviewFeedback_CreateFeedbackOnEvaluation struct {
	ID                           string                        `json:"id"`
	Comment                      string                        `json:"comment"`
	Data_FilledChecklist_CreateFeedbackOnEvaluation              Data_FilledChecklist_CreateFeedbackOnEvaluation               `json:"filledChecklist"`
	ReviewFeedbackCategoryValues []Data_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation `json:"reviewFeedbackCategoryValues"`
	Typename                     string                        `json:"__typename"`
}

type Data_FilledChecklist_CreateFeedbackOnEvaluation struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Data_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation struct {
	FeedbackCategory string `json:"feedbackCategory"`
	FeedbackValue    string `json:"feedbackValue"`
	ID               string `json:"id"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) CreateFeedbackOnEvaluation(variables Variables_CreateFeedbackOnEvaluation) (Data_CreateFeedbackOnEvaluation, error) {
	request := gql.NewQueryRequest[Variables_CreateFeedbackOnEvaluation](
		"mutation createFeedbackOnEvaluation($reviewFeedbackInput: ReviewFeedbackInput!) {\n  student {\n    createReviewFeedback(reviewFeedbackInput: $reviewFeedbackInput) {\n      ...EvaluationFeedback\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment EvaluationFeedback on ReviewFeedback {\n  id\n  comment\n  filledChecklist {\n    id\n    __typename\n  }\n  reviewFeedbackCategoryValues {\n    feedbackCategory\n    feedbackValue\n    id\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CreateFeedbackOnEvaluation](ctx, request)
}
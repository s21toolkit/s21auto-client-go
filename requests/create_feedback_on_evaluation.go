package requests

import "github.com/s21toolkit/s21client/gql"

type CreateFeedbackOnEvaluation_Variables struct {
	ReviewFeedbackInput CreateFeedbackOnEvaluation_Variables_ReviewFeedbackInput `json:"reviewFeedbackInput"`
}

type CreateFeedbackOnEvaluation_Variables_ReviewFeedbackInput struct {
	FilledChecklistID            string                        `json:"filledChecklistId"`
	Comment                      string                        `json:"comment"`
	ReviewFeedbackCategoryValues []CreateFeedbackOnEvaluation_Variables_ReviewFeedbackCategoryValue `json:"reviewFeedbackCategoryValues"`
}

type CreateFeedbackOnEvaluation_Variables_ReviewFeedbackCategoryValue struct {
	FeedbackCategory      string `json:"feedbackCategory"`
	FeedbackCategoryValue string `json:"feedbackCategoryValue"`
}


type CreateFeedbackOnEvaluation_Data struct {
	Student CreateFeedbackOnEvaluation_Data_Student `json:"student"`
}

type CreateFeedbackOnEvaluation_Data_Student struct {
	CreateReviewFeedback CreateFeedbackOnEvaluation_Data_CreateReviewFeedback `json:"createReviewFeedback"`
	Typename             string               `json:"__typename"`
}

type CreateFeedbackOnEvaluation_Data_CreateReviewFeedback struct {
	ID                           string                        `json:"id"`
	Comment                      string                        `json:"comment"`
	FilledChecklist              CreateFeedbackOnEvaluation_Data_FilledChecklist               `json:"filledChecklist"`
	ReviewFeedbackCategoryValues []CreateFeedbackOnEvaluation_Data_ReviewFeedbackCategoryValue `json:"reviewFeedbackCategoryValues"`
	Typename                     string                        `json:"__typename"`
}

type CreateFeedbackOnEvaluation_Data_FilledChecklist struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type CreateFeedbackOnEvaluation_Data_ReviewFeedbackCategoryValue struct {
	FeedbackCategory string `json:"feedbackCategory"`
	FeedbackValue    string `json:"feedbackValue"`
	ID               string `json:"id"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) CreateFeedbackOnEvaluation(variables CreateFeedbackOnEvaluation_Variables) (CreateFeedbackOnEvaluation_Data, error) {
	request := gql.NewQueryRequest[CreateFeedbackOnEvaluation_Variables](
		"mutation createFeedbackOnEvaluation($reviewFeedbackInput: ReviewFeedbackInput!) {\n  student {\n    createReviewFeedback(reviewFeedbackInput: $reviewFeedbackInput) {\n      ...EvaluationFeedback\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment EvaluationFeedback on ReviewFeedback {\n  id\n  comment\n  filledChecklist {\n    id\n    __typename\n  }\n  reviewFeedbackCategoryValues {\n    feedbackCategory\n    feedbackValue\n    id\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CreateFeedbackOnEvaluation_Data](ctx, request)
}
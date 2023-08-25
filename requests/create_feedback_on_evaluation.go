package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CreateFeedbackOnEvaluation struct {
	Request_ReviewFeedbackInput_CreateFeedbackOnEvaluation Request_ReviewFeedbackInput_CreateFeedbackOnEvaluation `json:"reviewFeedbackInput"`
}

type Request_ReviewFeedbackInput_CreateFeedbackOnEvaluation struct {
	FilledChecklistID            string                        `json:"filledChecklistId"`
	Comment                      string                        `json:"comment"`
	ReviewFeedbackCategoryValues []Request_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation `json:"reviewFeedbackCategoryValues"`
}

type Request_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation struct {
	FeedbackCategory      string `json:"feedbackCategory"`
	FeedbackCategoryValue string `json:"feedbackCategoryValue"`
}


type Response_Data_CreateFeedbackOnEvaluation struct {
	Response_Student_CreateFeedbackOnEvaluation Response_Student_CreateFeedbackOnEvaluation `json:"student"`
}

type Response_Student_CreateFeedbackOnEvaluation struct {
	Response_CreateReviewFeedback_CreateFeedbackOnEvaluation Response_CreateReviewFeedback_CreateFeedbackOnEvaluation `json:"createReviewFeedback"`
	Typename             string               `json:"__typename"`
}

type Response_CreateReviewFeedback_CreateFeedbackOnEvaluation struct {
	ID                           string                        `json:"id"`
	Comment                      string                        `json:"comment"`
	Response_FilledChecklist_CreateFeedbackOnEvaluation              Response_FilledChecklist_CreateFeedbackOnEvaluation               `json:"filledChecklist"`
	ReviewFeedbackCategoryValues []Response_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation `json:"reviewFeedbackCategoryValues"`
	Typename                     string                        `json:"__typename"`
}

type Response_FilledChecklist_CreateFeedbackOnEvaluation struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Response_ReviewFeedbackCategoryValue_CreateFeedbackOnEvaluation struct {
	FeedbackCategory string `json:"feedbackCategory"`
	FeedbackValue    string `json:"feedbackValue"`
	ID               string `json:"id"`
	Typename         string `json:"__typename"`
}

func (ctx *RequestContext) CreateFeedbackOnEvaluation(variables Request_Variables_CreateFeedbackOnEvaluation) (Response_Data_CreateFeedbackOnEvaluation, error) {
	request := gql.NewQueryRequest[Request_Variables_CreateFeedbackOnEvaluation](
		"mutation createFeedbackOnEvaluation($reviewFeedbackInput: ReviewFeedbackInput!) {   student {     createReviewFeedback(reviewFeedbackInput: $reviewFeedbackInput) {       ...EvaluationFeedback       __typename     }     __typename   } }  fragment EvaluationFeedback on ReviewFeedback {   id   comment   filledChecklist {     id     __typename   }   reviewFeedbackCategoryValues {     feedbackCategory     feedbackValue     id     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_CreateFeedbackOnEvaluation](ctx, request)
}
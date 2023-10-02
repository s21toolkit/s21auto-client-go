package requests

import "github.com/s21toolkit/s21client/gql"

type SaveFilledChecklist_Variables struct {
	FilledChecklistInput SaveFilledChecklist_Variables_FilledChecklistInput `json:"filledChecklistInput"`
}

type SaveFilledChecklist_Variables_FilledChecklistInput struct {
	QuickAction       interface{}     `json:"quickAction"`
	Comment           string          `json:"comment"`
	ScoreQuestions    []SaveFilledChecklist_Variables_ScoreQuestion `json:"scoreQuestions"`
	FilledChecklistID string          `json:"filledChecklistId"`
}

type SaveFilledChecklist_Variables_ScoreQuestion struct {
	SectionQuestionID string `json:"sectionQuestionId"`
	RatingWeightID    string `json:"ratingWeightId"`
}


type SaveFilledChecklist_Data struct {
	Student SaveFilledChecklist_Data_Student `json:"student"`
}

type SaveFilledChecklist_Data_Student struct {
	CompleteP2PCheck int64  `json:"completeP2pCheck"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) SaveFilledChecklist(variables SaveFilledChecklist_Variables) (SaveFilledChecklist_Data, error) {
	request := gql.NewQueryRequest[SaveFilledChecklist_Variables](
		"mutation saveFilledChecklist($filledChecklistInput: ChecklistFilledInput!) {\n  student {\n    completeP2pCheck(checklistFilledInput: $filledChecklistInput)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[SaveFilledChecklist_Data](ctx, request)
}
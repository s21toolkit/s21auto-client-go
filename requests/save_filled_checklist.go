package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_SaveFilledChecklist struct {
	Variables_FilledChecklistInput_SaveFilledChecklist Variables_FilledChecklistInput_SaveFilledChecklist `json:"filledChecklistInput"`
}

type Variables_FilledChecklistInput_SaveFilledChecklist struct {
	QuickAction       interface{}     `json:"quickAction"`
	Comment           string          `json:"comment"`
	ScoreQuestions    []Variables_ScoreQuestion_SaveFilledChecklist `json:"scoreQuestions"`
	FilledChecklistID string          `json:"filledChecklistId"`
}

type Variables_ScoreQuestion_SaveFilledChecklist struct {
	SectionQuestionID string `json:"sectionQuestionId"`
	RatingWeightID    string `json:"ratingWeightId"`
}


type Data_SaveFilledChecklist struct {
	Data_Student_SaveFilledChecklist Data_Student_SaveFilledChecklist `json:"student"`
}

type Data_Student_SaveFilledChecklist struct {
	CompleteP2PCheck int64  `json:"completeP2pCheck"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) SaveFilledChecklist(variables Variables_SaveFilledChecklist) (Data_SaveFilledChecklist, error) {
	request := gql.NewQueryRequest[Variables_SaveFilledChecklist](
		"mutation saveFilledChecklist($filledChecklistInput: ChecklistFilledInput!) {\n  student {\n    completeP2pCheck(checklistFilledInput: $filledChecklistInput)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_SaveFilledChecklist](ctx, request)
}
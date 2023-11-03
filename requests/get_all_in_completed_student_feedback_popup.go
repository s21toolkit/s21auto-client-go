package requests

import "github.com/s21toolkit/s21client/gql"

type GetAllInCompletedStudentFeedbackPopup_Variables struct {
}


type GetAllInCompletedStudentFeedbackPopup_Data struct {
	Sc21StudentTaskCheck GetAllInCompletedStudentFeedbackPopup_Data_Sc21StudentTaskCheck `json:"sc21StudentTaskCheck"`
}

type GetAllInCompletedStudentFeedbackPopup_Data_Sc21StudentTaskCheck struct {
	GetAllInCompletedStudentFeedbackPopup []interface{} `json:"getAllInCompletedStudentFeedbackPopup"`
	Typename                              string        `json:"__typename"`
}


func (ctx *RequestContext) GetAllInCompletedStudentFeedbackPopup(variables GetAllInCompletedStudentFeedbackPopup_Variables) (GetAllInCompletedStudentFeedbackPopup_Data, error) {
	request := gql.NewQueryRequest[GetAllInCompletedStudentFeedbackPopup_Variables](
		"query getAllInCompletedStudentFeedbackPopup {\n  sc21StudentTaskCheck {\n    getAllInCompletedStudentFeedbackPopup {\n      studentFeedbackId\n      goalName\n      resultAttemptDate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetAllInCompletedStudentFeedbackPopup_Data](ctx, request)
}
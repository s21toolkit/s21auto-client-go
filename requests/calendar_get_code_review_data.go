package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetCodeReviewData struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_CalendarGetCodeReviewData struct {
	Data_Student_CalendarGetCodeReviewData Data_Student_CalendarGetCodeReviewData `json:"student"`
}

type Data_Student_CalendarGetCodeReviewData struct {
	Data_GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData Data_GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type Data_GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData struct {
	ModuleTitle string      `json:"moduleTitle"`
	Data_CurrentTask_CalendarGetCodeReviewData Data_CurrentTask_CalendarGetCodeReviewData `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type Data_CurrentTask_CalendarGetCodeReviewData struct {
	Data_Task_CalendarGetCodeReviewData     Data_Task_CalendarGetCodeReviewData   `json:"task"`
	Typename string `json:"__typename"`
}

type Data_Task_CalendarGetCodeReviewData struct {
	Data_StudentTaskAdditionalAttributes_CalendarGetCodeReviewData Data_StudentTaskAdditionalAttributes_CalendarGetCodeReviewData `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_CalendarGetCodeReviewData struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetCodeReviewData(variables Variables_CalendarGetCodeReviewData) (Data_CalendarGetCodeReviewData, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetCodeReviewData](
		"query calendarGetCodeReviewData($studentGoalId: ID!) {\n  student {\n    getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {\n      moduleTitle\n      currentTask {\n        task {\n          studentTaskAdditionalAttributes {\n            codeReviewDuration\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetCodeReviewData](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetCodeReviewData_Variables struct {
	StudentGoalID string `json:"studentGoalId"`
}


type CalendarGetCodeReviewData_Data struct {
	Student CalendarGetCodeReviewData_Data_Student `json:"student"`
}

type CalendarGetCodeReviewData_Data_Student struct {
	GetStudentModuleByStudentGoalID CalendarGetCodeReviewData_Data_GetStudentModuleByStudentGoalID `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetCodeReviewData_Data_GetStudentModuleByStudentGoalID struct {
	ModuleTitle string      `json:"moduleTitle"`
	CurrentTask CalendarGetCodeReviewData_Data_CurrentTask `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type CalendarGetCodeReviewData_Data_CurrentTask struct {
	Task     CalendarGetCodeReviewData_Data_Task   `json:"task"`
	Typename string `json:"__typename"`
}

type CalendarGetCodeReviewData_Data_Task struct {
	StudentTaskAdditionalAttributes CalendarGetCodeReviewData_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetCodeReviewData_Data_StudentTaskAdditionalAttributes struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetCodeReviewData(variables CalendarGetCodeReviewData_Variables) (CalendarGetCodeReviewData_Data, error) {
	request := gql.NewQueryRequest[CalendarGetCodeReviewData_Variables](
		"query calendarGetCodeReviewData($studentGoalId: ID!) {\n  student {\n    getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {\n      moduleTitle\n      currentTask {\n        task {\n          studentTaskAdditionalAttributes {\n            codeReviewDuration\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetCodeReviewData_Data](ctx, request)
}
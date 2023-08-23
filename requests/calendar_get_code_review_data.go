package requests

import "s21client/gql"

type Variables_CalendarGetCodeReviewData struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_CalendarGetCodeReviewData struct {
	Student_CalendarGetCodeReviewData Student_CalendarGetCodeReviewData `json:"student"`
}

type Student_CalendarGetCodeReviewData struct {
	GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type GetStudentModuleByStudentGoalID_CalendarGetCodeReviewData struct {
	ModuleTitle string      `json:"moduleTitle"`
	CurrentTask_CalendarGetCodeReviewData CurrentTask_CalendarGetCodeReviewData `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type CurrentTask_CalendarGetCodeReviewData struct {
	Task_CalendarGetCodeReviewData     Task_CalendarGetCodeReviewData   `json:"task"`
	Typename string `json:"__typename"`
}

type Task_CalendarGetCodeReviewData struct {
	StudentTaskAdditionalAttributes_CalendarGetCodeReviewData StudentTaskAdditionalAttributes_CalendarGetCodeReviewData `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type StudentTaskAdditionalAttributes_CalendarGetCodeReviewData struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetCodeReviewData(variables Variables_CalendarGetCodeReviewData) (Data_CalendarGetCodeReviewData, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetCodeReviewData](
		"query calendarGetCodeReviewData($studentGoalId: ID!) {   student {     getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {       moduleTitle       currentTask {         task {           studentTaskAdditionalAttributes {             codeReviewDuration             __typename           }           __typename         }         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_CalendarGetCodeReviewData](ctx, request)
}
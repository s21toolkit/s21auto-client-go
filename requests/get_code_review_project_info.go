package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCodeReviewProjectInfo struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_GetCodeReviewProjectInfo struct {
	Data_Student_GetCodeReviewProjectInfo Data_Student_GetCodeReviewProjectInfo `json:"student"`
}

type Data_Student_GetCodeReviewProjectInfo struct {
	Data_GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo Data_GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type Data_GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo struct {
	ID          string      `json:"id"`
	ModuleTitle string      `json:"moduleTitle"`
	Data_StudyModule_GetCodeReviewProjectInfo Data_StudyModule_GetCodeReviewProjectInfo `json:"studyModule"`
	Data_CurrentTask_GetCodeReviewProjectInfo Data_CurrentTask_GetCodeReviewProjectInfo `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type Data_CurrentTask_GetCodeReviewProjectInfo struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Data_Task_GetCodeReviewProjectInfo     Data_Task_GetCodeReviewProjectInfo   `json:"task"`
	Typename string `json:"__typename"`
}

type Data_Task_GetCodeReviewProjectInfo struct {
	Data_Content_GetCodeReviewProjectInfo                         Data_Content_GetCodeReviewProjectInfo                         `json:"content"`
	AssignmentType                  string                          `json:"assignmentType"`
	Data_StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo Data_StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Data_Content_GetCodeReviewProjectInfo struct {
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	Typename           string `json:"__typename"`
}

type Data_StudyModule_GetCodeReviewProjectInfo struct {
	Duration int64  `json:"duration"`
	Data_Stage_GetCodeReviewProjectInfo    Data_Stage_GetCodeReviewProjectInfo  `json:"stage"`
	Typename string `json:"__typename"`
}

type Data_Stage_GetCodeReviewProjectInfo struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewProjectInfo(variables Variables_GetCodeReviewProjectInfo) (Data_GetCodeReviewProjectInfo, error) {
	request := gql.NewQueryRequest[Variables_GetCodeReviewProjectInfo](
		"query getCodeReviewProjectInfo($studentGoalId: ID!) {\n  student {\n    getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {\n      ...CodeReviewProjectInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProjectInfo on StudentModule {\n  id\n  moduleTitle\n  studyModule {\n    duration\n    stage {\n      name\n      __typename\n    }\n    __typename\n  }\n  currentTask {\n    ...CodeReviewCurrentTaskInfo\n    __typename\n  }\n  __typename\n}\n\nfragment CodeReviewCurrentTaskInfo on StudentTask {\n  id\n  taskId\n  task {\n    content {\n      body\n      __typename\n    }\n    assignmentType\n    studentTaskAdditionalAttributes {\n      codeReviewDuration\n      codeReviewCost\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewProjectInfo](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type GetCodeReviewProjectInfo_Variables struct {
	StudentGoalID string `json:"studentGoalId"`
}


type GetCodeReviewProjectInfo_Data struct {
	Student GetCodeReviewProjectInfo_Data_Student `json:"student"`
}

type GetCodeReviewProjectInfo_Data_Student struct {
	GetStudentModuleByStudentGoalID GetCodeReviewProjectInfo_Data_GetStudentModuleByStudentGoalID `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_GetStudentModuleByStudentGoalID struct {
	ID          string      `json:"id"`
	ModuleTitle string      `json:"moduleTitle"`
	StudyModule GetCodeReviewProjectInfo_Data_StudyModule `json:"studyModule"`
	CurrentTask GetCodeReviewProjectInfo_Data_CurrentTask `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_CurrentTask struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Task     GetCodeReviewProjectInfo_Data_Task   `json:"task"`
	Typename string `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_Task struct {
	Content                         GetCodeReviewProjectInfo_Data_Content                         `json:"content"`
	AssignmentType                  string                          `json:"assignmentType"`
	StudentTaskAdditionalAttributes GetCodeReviewProjectInfo_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_Content struct {
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_StudentTaskAdditionalAttributes struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	Typename           string `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_StudyModule struct {
	Duration int64  `json:"duration"`
	Stage    GetCodeReviewProjectInfo_Data_Stage  `json:"stage"`
	Typename string `json:"__typename"`
}

type GetCodeReviewProjectInfo_Data_Stage struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetCodeReviewProjectInfo(variables GetCodeReviewProjectInfo_Variables) (GetCodeReviewProjectInfo_Data, error) {
	request := gql.NewQueryRequest[GetCodeReviewProjectInfo_Variables](
		"query getCodeReviewProjectInfo($studentGoalId: ID!) {\n  student {\n    getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {\n      ...CodeReviewProjectInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProjectInfo on StudentModule {\n  id\n  moduleTitle\n  studyModule {\n    duration\n    stage {\n      name\n      __typename\n    }\n    __typename\n  }\n  currentTask {\n    ...CodeReviewCurrentTaskInfo\n    __typename\n  }\n  __typename\n}\n\nfragment CodeReviewCurrentTaskInfo on StudentTask {\n  id\n  taskId\n  task {\n    content {\n      body\n      __typename\n    }\n    assignmentType\n    studentTaskAdditionalAttributes {\n      codeReviewDuration\n      codeReviewCost\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetCodeReviewProjectInfo_Data](ctx, request)
}
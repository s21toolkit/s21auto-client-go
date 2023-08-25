package requests

import "s21client/gql"

type Variables_GetCodeReviewProjectInfo struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_GetCodeReviewProjectInfo struct {
	Student_GetCodeReviewProjectInfo Student_GetCodeReviewProjectInfo `json:"student"`
}

type Student_GetCodeReviewProjectInfo struct {
	GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo `json:"getStudentModuleByStudentGoalId"`
	Typename                        string                          `json:"__typename"`
}

type GetStudentModuleByStudentGoalID_GetCodeReviewProjectInfo struct {
	ID          string      `json:"id"`
	ModuleTitle string      `json:"moduleTitle"`
	StudyModule_GetCodeReviewProjectInfo StudyModule_GetCodeReviewProjectInfo `json:"studyModule"`
	CurrentTask_GetCodeReviewProjectInfo CurrentTask_GetCodeReviewProjectInfo `json:"currentTask"`
	Typename    string      `json:"__typename"`
}

type CurrentTask_GetCodeReviewProjectInfo struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Task_GetCodeReviewProjectInfo     Task_GetCodeReviewProjectInfo   `json:"task"`
	Typename string `json:"__typename"`
}

type Task_GetCodeReviewProjectInfo struct {
	Content_GetCodeReviewProjectInfo                         Content_GetCodeReviewProjectInfo                         `json:"content"`
	AssignmentType                  string                          `json:"assignmentType"`
	StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Content_GetCodeReviewProjectInfo struct {
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}

type StudentTaskAdditionalAttributes_GetCodeReviewProjectInfo struct {
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	Typename           string `json:"__typename"`
}

type StudyModule_GetCodeReviewProjectInfo struct {
	Duration int64  `json:"duration"`
	Stage_GetCodeReviewProjectInfo    Stage_GetCodeReviewProjectInfo  `json:"stage"`
	Typename string `json:"__typename"`
}

type Stage_GetCodeReviewProjectInfo struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetCodeReviewProjectInfo(variables Variables_GetCodeReviewProjectInfo) (Data_GetCodeReviewProjectInfo, error) {
	request := gql.NewQueryRequest[Variables_GetCodeReviewProjectInfo](
		"query getCodeReviewProjectInfo($studentGoalId: ID!) {   student {     getStudentModuleByStudentGoalId(studentGoalId: $studentGoalId) {       ...CodeReviewProjectInfo       __typename     }     __typename   } }  fragment CodeReviewProjectInfo on StudentModule {   id   moduleTitle   studyModule {     duration     stage {       name       __typename     }     __typename   }   currentTask {     ...CodeReviewCurrentTaskInfo     __typename   }   __typename }  fragment CodeReviewCurrentTaskInfo on StudentTask {   id   taskId   task {     content {       body       __typename     }     assignmentType     studentTaskAdditionalAttributes {       codeReviewDuration       codeReviewCost       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetCodeReviewProjectInfo](ctx, request)
}
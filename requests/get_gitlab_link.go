package requests

import "github.com/s21toolkit/s21client/gql"

type GetGitlabLink_Variables struct {
	TaskID string `json:"taskId"`
}


type GetGitlabLink_Data struct {
	Student GetGitlabLink_Data_Student `json:"student"`
}

type GetGitlabLink_Data_Student struct {
	GetLinkToPrivateStudentGitlabProjectByTaskID GetGitlabLink_Data_GetLinkToPrivateStudentGitlabProjectByTaskID `json:"getLinkToPrivateStudentGitlabProjectByTaskId"`
	Typename                                     string                                       `json:"__typename"`
}

type GetGitlabLink_Data_GetLinkToPrivateStudentGitlabProjectByTaskID struct {
	SSHLink     string `json:"sshLink"`
	HasOpenedMR bool   `json:"hasOpenedMR"`
	HTTPSLink   string `json:"httpsLink"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) GetGitlabLink(variables GetGitlabLink_Variables) (GetGitlabLink_Data, error) {
	request := gql.NewQueryRequest[GetGitlabLink_Variables](
		"query getGitlabLink($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      sshLink\n      hasOpenedMR\n      httpsLink\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetGitlabLink_Data](ctx, request)
}
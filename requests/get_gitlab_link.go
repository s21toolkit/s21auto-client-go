package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetGitlabLink struct {
	TaskID string `json:"taskId"`
}


type Data_GetGitlabLink struct {
	Data_Student_GetGitlabLink Data_Student_GetGitlabLink `json:"student"`
}

type Data_Student_GetGitlabLink struct {
	Data_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink Data_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink `json:"getLinkToPrivateStudentGitlabProjectByTaskId"`
	Typename                                     string                                       `json:"__typename"`
}

type Data_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink struct {
	SSHLink     string `json:"sshLink"`
	HasOpenedMR bool   `json:"hasOpenedMR"`
	HTTPSLink   string `json:"httpsLink"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) GetGitlabLink(variables Variables_GetGitlabLink) (Data_GetGitlabLink, error) {
	request := gql.NewQueryRequest[Variables_GetGitlabLink](
		"query getGitlabLink($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      sshLink\n      hasOpenedMR\n      httpsLink\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetGitlabLink](ctx, request)
}
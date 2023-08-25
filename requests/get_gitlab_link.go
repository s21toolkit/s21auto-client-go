package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetGitlabLink struct {
	TaskID string `json:"taskId"`
}


type Response_Data_GetGitlabLink struct {
	Response_Student_GetGitlabLink Response_Student_GetGitlabLink `json:"student"`
}

type Response_Student_GetGitlabLink struct {
	Response_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink Response_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink `json:"getLinkToPrivateStudentGitlabProjectByTaskId"`
	Typename                                     string                                       `json:"__typename"`
}

type Response_GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink struct {
	SSHLink     string `json:"sshLink"`
	HasOpenedMR bool   `json:"hasOpenedMR"`
	HTTPSLink   string `json:"httpsLink"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) GetGitlabLink(variables Request_Variables_GetGitlabLink) (Response_Data_GetGitlabLink, error) {
	request := gql.NewQueryRequest[Request_Variables_GetGitlabLink](
		"query getGitlabLink($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      sshLink\n      hasOpenedMR\n      httpsLink\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetGitlabLink](ctx, request)
}
package requests

import "s21client/gql"

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
		"query getGitlabLink($taskId: ID!) {   student {     getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {       sshLink       hasOpenedMR       httpsLink       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetGitlabLink](ctx, request)
}
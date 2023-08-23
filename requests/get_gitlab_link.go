package requests

import "s21client/gql"

type Variables_GetGitlabLink struct {
	TaskID string `json:"taskId"`
}


type Data_GetGitlabLink struct {
	Student_GetGitlabLink Student_GetGitlabLink `json:"student"`
}

type Student_GetGitlabLink struct {
	GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink `json:"getLinkToPrivateStudentGitlabProjectByTaskId"`
	Typename                                     string                                       `json:"__typename"`
}

type GetLinkToPrivateStudentGitlabProjectByTaskID_GetGitlabLink struct {
	SSHLink     string `json:"sshLink"`
	HasOpenedMR bool   `json:"hasOpenedMR"`
	HTTPSLink   string `json:"httpsLink"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) GetGitlabLink(variables Variables_GetGitlabLink) (Data_GetGitlabLink, error) {
	request := gql.NewQueryRequest[Variables_GetGitlabLink](
		"query getGitlabLink($taskId: ID!) {   student {     getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {       sshLink       hasOpenedMR       httpsLink       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetGitlabLink](ctx, request)
}
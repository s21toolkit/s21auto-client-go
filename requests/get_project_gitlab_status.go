package requests

import "s21client/gql"

type Variables_GetProjectGitlabStatus struct {
	TaskID string `json:"taskId"`
}


type Data_GetProjectGitlabStatus struct {
	Student_GetProjectGitlabStatus Student_GetProjectGitlabStatus `json:"student"`
}

type Student_GetProjectGitlabStatus struct {
	GetGitlabLinksWithStatus_GetProjectGitlabStatus GetGitlabLinksWithStatus_GetProjectGitlabStatus `json:"getGitlabLinksWithStatus"`
	Typename                 string                   `json:"__typename"`
}

type GetGitlabLinksWithStatus_GetProjectGitlabStatus struct {
	ID              int64  `json:"id"`
	SSHLink         string `json:"sshLink"`
	HTTPSLink       string `json:"httpsLink"`
	ReadyToUse      string `json:"readyToUse"`
	RestartsCounter int64  `json:"restartsCounter"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectGitlabStatus(variables Variables_GetProjectGitlabStatus) (Data_GetProjectGitlabStatus, error) {
	request := gql.NewQueryRequest[Variables_GetProjectGitlabStatus](
		"query getProjectGitlabStatus($taskId: ID!) {   student {     getGitlabLinksWithStatus(taskId: $taskId) {       id       sshLink       httpsLink       readyToUse       restartsCounter       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetProjectGitlabStatus](ctx, request)
}
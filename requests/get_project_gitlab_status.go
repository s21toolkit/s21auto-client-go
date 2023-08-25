package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetProjectGitlabStatus struct {
	TaskID string `json:"taskId"`
}


type Response_Data_GetProjectGitlabStatus struct {
	Response_Student_GetProjectGitlabStatus Response_Student_GetProjectGitlabStatus `json:"student"`
}

type Response_Student_GetProjectGitlabStatus struct {
	Response_GetGitlabLinksWithStatus_GetProjectGitlabStatus Response_GetGitlabLinksWithStatus_GetProjectGitlabStatus `json:"getGitlabLinksWithStatus"`
	Typename                 string                   `json:"__typename"`
}

type Response_GetGitlabLinksWithStatus_GetProjectGitlabStatus struct {
	ID              int64  `json:"id"`
	SSHLink         string `json:"sshLink"`
	HTTPSLink       string `json:"httpsLink"`
	ReadyToUse      string `json:"readyToUse"`
	RestartsCounter int64  `json:"restartsCounter"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectGitlabStatus(variables Request_Variables_GetProjectGitlabStatus) (Response_Data_GetProjectGitlabStatus, error) {
	request := gql.NewQueryRequest[Request_Variables_GetProjectGitlabStatus](
		"query getProjectGitlabStatus($taskId: ID!) {   student {     getGitlabLinksWithStatus(taskId: $taskId) {       id       sshLink       httpsLink       readyToUse       restartsCounter       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetProjectGitlabStatus](ctx, request)
}
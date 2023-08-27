package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetProjectGitlabStatus struct {
	TaskID string `json:"taskId"`
}


type Data_GetProjectGitlabStatus struct {
	Data_Student_GetProjectGitlabStatus Data_Student_GetProjectGitlabStatus `json:"student"`
}

type Data_Student_GetProjectGitlabStatus struct {
	Data_GetGitlabLinksWithStatus_GetProjectGitlabStatus Data_GetGitlabLinksWithStatus_GetProjectGitlabStatus `json:"getGitlabLinksWithStatus"`
	Typename                 string                   `json:"__typename"`
}

type Data_GetGitlabLinksWithStatus_GetProjectGitlabStatus struct {
	ID              int64  `json:"id"`
	SSHLink         string `json:"sshLink"`
	HTTPSLink       string `json:"httpsLink"`
	ReadyToUse      string `json:"readyToUse"`
	RestartsCounter int64  `json:"restartsCounter"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectGitlabStatus(variables Variables_GetProjectGitlabStatus) (Data_GetProjectGitlabStatus, error) {
	request := gql.NewQueryRequest[Variables_GetProjectGitlabStatus](
		"query getProjectGitlabStatus($taskId: ID!) {\n  student {\n    getGitlabLinksWithStatus(taskId: $taskId) {\n      id\n      sshLink\n      httpsLink\n      readyToUse\n      restartsCounter\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetProjectGitlabStatus](ctx, request)
}
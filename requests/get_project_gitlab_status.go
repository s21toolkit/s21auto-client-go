package requests

import "github.com/s21toolkit/s21client/gql"

type GetProjectGitlabStatus_Variables struct {
	TaskID string `json:"taskId"`
}


type GetProjectGitlabStatus_Data struct {
	Student GetProjectGitlabStatus_Data_Student `json:"student"`
}

type GetProjectGitlabStatus_Data_Student struct {
	GetGitlabLinksWithStatus GetProjectGitlabStatus_Data_GetGitlabLinksWithStatus `json:"getGitlabLinksWithStatus"`
	Typename                 string                   `json:"__typename"`
}

type GetProjectGitlabStatus_Data_GetGitlabLinksWithStatus struct {
	ID              int64  `json:"id"`
	SSHLink         string `json:"sshLink"`
	HTTPSLink       string `json:"httpsLink"`
	ReadyToUse      string `json:"readyToUse"`
	RestartsCounter int64  `json:"restartsCounter"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectGitlabStatus(variables GetProjectGitlabStatus_Variables) (GetProjectGitlabStatus_Data, error) {
	request := gql.NewQueryRequest[GetProjectGitlabStatus_Variables](
		"query getProjectGitlabStatus($taskId: ID!) {\n  student {\n    getGitlabLinksWithStatus(taskId: $taskId) {\n      id\n      sshLink\n      httpsLink\n      readyToUse\n      restartsCounter\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetProjectGitlabStatus_Data](ctx, request)
}
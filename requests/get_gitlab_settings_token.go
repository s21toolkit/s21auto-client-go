package requests

import "github.com/s21toolkit/s21client/gql"

type GetGitlabSettingsToken_Variables struct {
	TaskID string `json:"taskId"`
}


type GetGitlabSettingsToken_Data struct {
	Student interface{} `json:"student"`
}


func (ctx *RequestContext) GetGitlabSettingsToken(variables GetGitlabSettingsToken_Variables) (GetGitlabSettingsToken_Data, error) {
	request := gql.NewQueryRequest[GetGitlabSettingsToken_Variables](
		"query getGitlabSettingsToken($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      runnersToken\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetGitlabSettingsToken_Data](ctx, request)
}
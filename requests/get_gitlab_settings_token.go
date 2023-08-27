package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetGitlabSettingsToken struct {
	TaskID string `json:"taskId"`
}


type Data_GetGitlabSettingsToken struct {
	Student interface{} `json:"student"`
}


func (ctx *RequestContext) GetGitlabSettingsToken(variables Variables_GetGitlabSettingsToken) (Data_GetGitlabSettingsToken, error) {
	request := gql.NewQueryRequest[Variables_GetGitlabSettingsToken](
		"query getGitlabSettingsToken($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      runnersToken\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetGitlabSettingsToken](ctx, request)
}
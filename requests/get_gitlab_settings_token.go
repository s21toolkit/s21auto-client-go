package requests

import "s21client/gql"

type Variables_GetGitlabSettingsToken struct {
	TaskID string `json:"taskId"`
}


type Data_GetGitlabSettingsToken struct {
	Student interface{} `json:"student"`
}

func (ctx *RequestContext) GetGitlabSettingsToken(variables Variables_GetGitlabSettingsToken) (Data_GetGitlabSettingsToken, error) {
	request := gql.NewQueryRequest[Variables_GetGitlabSettingsToken](
		"query getGitlabSettingsToken($taskId: ID!) {   student {     getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {       runnersToken       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetGitlabSettingsToken](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetInvitationsCount struct {
}


type Response_Data_GetInvitationsCount struct {
	Response_Team_GetInvitationsCount Response_Team_GetInvitationsCount `json:"team"`
}

type Response_Team_GetInvitationsCount struct {
	GetCreatedJoinTeamRequestCount int64  `json:"getCreatedJoinTeamRequestCount"`
	Typename                       string `json:"__typename"`
}

func (ctx *RequestContext) GetInvitationsCount(variables Request_Variables_GetInvitationsCount) (Response_Data_GetInvitationsCount, error) {
	request := gql.NewQueryRequest[Request_Variables_GetInvitationsCount](
		"query getInvitationsCount {\n  team {\n    getCreatedJoinTeamRequestCount\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetInvitationsCount](ctx, request)
}
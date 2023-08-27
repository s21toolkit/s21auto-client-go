package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetInvitationsCount struct {
}


type Data_GetInvitationsCount struct {
	Data_Team_GetInvitationsCount Data_Team_GetInvitationsCount `json:"team"`
}

type Data_Team_GetInvitationsCount struct {
	GetCreatedJoinTeamRequestCount int64  `json:"getCreatedJoinTeamRequestCount"`
	Typename                       string `json:"__typename"`
}


func (ctx *RequestContext) GetInvitationsCount(variables Variables_GetInvitationsCount) (Data_GetInvitationsCount, error) {
	request := gql.NewQueryRequest[Variables_GetInvitationsCount](
		"query getInvitationsCount {\n  team {\n    getCreatedJoinTeamRequestCount\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetInvitationsCount](ctx, request)
}
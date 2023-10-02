package requests

import "github.com/s21toolkit/s21client/gql"

type GetInvitationsCount_Variables struct {
}


type GetInvitationsCount_Data struct {
	Team GetInvitationsCount_Data_Team `json:"team"`
}

type GetInvitationsCount_Data_Team struct {
	GetCreatedJoinTeamRequestCount int64  `json:"getCreatedJoinTeamRequestCount"`
	Typename                       string `json:"__typename"`
}


func (ctx *RequestContext) GetInvitationsCount(variables GetInvitationsCount_Variables) (GetInvitationsCount_Data, error) {
	request := gql.NewQueryRequest[GetInvitationsCount_Variables](
		"query getInvitationsCount {\n  team {\n    getCreatedJoinTeamRequestCount\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetInvitationsCount_Data](ctx, request)
}
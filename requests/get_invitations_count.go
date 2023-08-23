package requests

import "s21client/gql"

type Variables_GetInvitationsCount struct {
}


type Data_GetInvitationsCount struct {
	Team_GetInvitationsCount Team_GetInvitationsCount `json:"team"`
}

type Team_GetInvitationsCount struct {
	GetCreatedJoinTeamRequestCount int64  `json:"getCreatedJoinTeamRequestCount"`
	Typename                       string `json:"__typename"`
}

func (ctx *RequestContext) GetInvitationsCount(variables Variables_GetInvitationsCount) (Data_GetInvitationsCount, error) {
	request := gql.NewQueryRequest[Variables_GetInvitationsCount](
		"query getInvitationsCount {   team {     getCreatedJoinTeamRequestCount     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetInvitationsCount](ctx, request)
}
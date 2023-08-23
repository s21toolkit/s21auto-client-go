package requests

import "s21client/gql"

type Variables_GetIsDisbandRequestAlreadySent struct {
	TeamID string `json:"teamId"`
}


type Data_GetIsDisbandRequestAlreadySent struct {
	Student_GetIsDisbandRequestAlreadySent Student_GetIsDisbandRequestAlreadySent `json:"student"`
}

type Student_GetIsDisbandRequestAlreadySent struct {
	IsRequestBeenSentToTeamDisband bool   `json:"isRequestBeenSentToTeamDisband"`
	Typename                       string `json:"__typename"`
}

func (ctx *RequestContext) GetIsDisbandRequestAlreadySent(variables Variables_GetIsDisbandRequestAlreadySent) (Data_GetIsDisbandRequestAlreadySent, error) {
	request := gql.NewQueryRequest[Variables_GetIsDisbandRequestAlreadySent](
		"query getIsDisbandRequestAlreadySent($teamId: UUID!) {   student {     isRequestBeenSentToTeamDisband(teamId: $teamId)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetIsDisbandRequestAlreadySent](ctx, request)
}
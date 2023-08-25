package requests

import "s21client/gql"

type Request_Variables_GetIsDisbandRequestAlreadySent struct {
	TeamID string `json:"teamId"`
}


type Response_Data_GetIsDisbandRequestAlreadySent struct {
	Response_Student_GetIsDisbandRequestAlreadySent Response_Student_GetIsDisbandRequestAlreadySent `json:"student"`
}

type Response_Student_GetIsDisbandRequestAlreadySent struct {
	IsRequestBeenSentToTeamDisband bool   `json:"isRequestBeenSentToTeamDisband"`
	Typename                       string `json:"__typename"`
}

func (ctx *RequestContext) GetIsDisbandRequestAlreadySent(variables Request_Variables_GetIsDisbandRequestAlreadySent) (Response_Data_GetIsDisbandRequestAlreadySent, error) {
	request := gql.NewQueryRequest[Request_Variables_GetIsDisbandRequestAlreadySent](
		"query getIsDisbandRequestAlreadySent($teamId: UUID!) {   student {     isRequestBeenSentToTeamDisband(teamId: $teamId)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetIsDisbandRequestAlreadySent](ctx, request)
}
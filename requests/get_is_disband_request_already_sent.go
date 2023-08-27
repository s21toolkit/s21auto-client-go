package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetIsDisbandRequestAlreadySent struct {
	TeamID string `json:"teamId"`
}


type Data_GetIsDisbandRequestAlreadySent struct {
	Data_Student_GetIsDisbandRequestAlreadySent Data_Student_GetIsDisbandRequestAlreadySent `json:"student"`
}

type Data_Student_GetIsDisbandRequestAlreadySent struct {
	IsRequestBeenSentToTeamDisband bool   `json:"isRequestBeenSentToTeamDisband"`
	Typename                       string `json:"__typename"`
}


func (ctx *RequestContext) GetIsDisbandRequestAlreadySent(variables Variables_GetIsDisbandRequestAlreadySent) (Data_GetIsDisbandRequestAlreadySent, error) {
	request := gql.NewQueryRequest[Variables_GetIsDisbandRequestAlreadySent](
		"query getIsDisbandRequestAlreadySent($teamId: UUID!) {\n  student {\n    isRequestBeenSentToTeamDisband(teamId: $teamId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetIsDisbandRequestAlreadySent](ctx, request)
}
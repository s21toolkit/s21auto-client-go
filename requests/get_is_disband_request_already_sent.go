package requests

import "github.com/s21toolkit/s21client/gql"

type GetIsDisbandRequestAlreadySent_Variables struct {
	TeamID string `json:"teamId"`
}


type GetIsDisbandRequestAlreadySent_Data struct {
	Student GetIsDisbandRequestAlreadySent_Data_Student `json:"student"`
}

type GetIsDisbandRequestAlreadySent_Data_Student struct {
	IsRequestBeenSentToTeamDisband bool   `json:"isRequestBeenSentToTeamDisband"`
	Typename                       string `json:"__typename"`
}


func (ctx *RequestContext) GetIsDisbandRequestAlreadySent(variables GetIsDisbandRequestAlreadySent_Variables) (GetIsDisbandRequestAlreadySent_Data, error) {
	request := gql.NewQueryRequest[GetIsDisbandRequestAlreadySent_Variables](
		"query getIsDisbandRequestAlreadySent($teamId: UUID!) {\n  student {\n    isRequestBeenSentToTeamDisband(teamId: $teamId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetIsDisbandRequestAlreadySent_Data](ctx, request)
}
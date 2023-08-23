package requests

import "s21client/gql"

type Variables_GetTournamentNotificationResults struct {
}


type Data_GetTournamentNotificationResults struct {
	Student_GetTournamentNotificationResults Student_GetTournamentNotificationResults `json:"student"`
}

type Student_GetTournamentNotificationResults struct {
	GetTournamentResults []interface{} `json:"getTournamentResults"`
	Typename             string        `json:"__typename"`
}

func (ctx *RequestContext) GetTournamentNotificationResults(variables Variables_GetTournamentNotificationResults) (Data_GetTournamentNotificationResults, error) {
	request := gql.NewQueryRequest[Variables_GetTournamentNotificationResults](
		"query getTournamentNotificationResults {   student {     getTournamentResults(isShown: false) {       id       power       coalitionRank       userRank       firstCoalitionName       coalitionName       timeClosed       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetTournamentNotificationResults](ctx, request)
}
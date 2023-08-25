package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetTournamentNotificationResults struct {
}


type Response_Data_GetTournamentNotificationResults struct {
	Response_Student_GetTournamentNotificationResults Response_Student_GetTournamentNotificationResults `json:"student"`
}

type Response_Student_GetTournamentNotificationResults struct {
	GetTournamentResults []interface{} `json:"getTournamentResults"`
	Typename             string        `json:"__typename"`
}

func (ctx *RequestContext) GetTournamentNotificationResults(variables Request_Variables_GetTournamentNotificationResults) (Response_Data_GetTournamentNotificationResults, error) {
	request := gql.NewQueryRequest[Request_Variables_GetTournamentNotificationResults](
		"query getTournamentNotificationResults {   student {     getTournamentResults(isShown: false) {       id       power       coalitionRank       userRank       firstCoalitionName       coalitionName       timeClosed       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetTournamentNotificationResults](ctx, request)
}
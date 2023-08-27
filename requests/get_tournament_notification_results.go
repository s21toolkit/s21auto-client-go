package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetTournamentNotificationResults struct {
}


type Data_GetTournamentNotificationResults struct {
	Data_Student_GetTournamentNotificationResults Data_Student_GetTournamentNotificationResults `json:"student"`
}

type Data_Student_GetTournamentNotificationResults struct {
	GetTournamentResults []interface{} `json:"getTournamentResults"`
	Typename             string        `json:"__typename"`
}


func (ctx *RequestContext) GetTournamentNotificationResults(variables Variables_GetTournamentNotificationResults) (Data_GetTournamentNotificationResults, error) {
	request := gql.NewQueryRequest[Variables_GetTournamentNotificationResults](
		"query getTournamentNotificationResults {\n  student {\n    getTournamentResults(isShown: false) {\n      id\n      power\n      coalitionRank\n      userRank\n      firstCoalitionName\n      coalitionName\n      timeClosed\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetTournamentNotificationResults](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type GetTournamentNotificationResults_Variables struct {
}


type GetTournamentNotificationResults_Data struct {
	Student GetTournamentNotificationResults_Data_Student `json:"student"`
}

type GetTournamentNotificationResults_Data_Student struct {
	GetTournamentResults []interface{} `json:"getTournamentResults"`
	Typename             string        `json:"__typename"`
}


func (ctx *RequestContext) GetTournamentNotificationResults(variables GetTournamentNotificationResults_Variables) (GetTournamentNotificationResults_Data, error) {
	request := gql.NewQueryRequest[GetTournamentNotificationResults_Variables](
		"query getTournamentNotificationResults {\n  student {\n    getTournamentResults(isShown: false) {\n      id\n      power\n      coalitionRank\n      userRank\n      firstCoalitionName\n      coalitionName\n      timeClosed\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetTournamentNotificationResults_Data](ctx, request)
}
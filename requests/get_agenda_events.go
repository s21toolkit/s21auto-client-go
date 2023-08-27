package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetAgendaEvents struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Limit int64  `json:"limit"`
}


type Data_GetAgendaEvents struct {
	Data_Student_GetAgendaEvents Data_Student_GetAgendaEvents `json:"student"`
}

type Data_Student_GetAgendaEvents struct {
	GetMyAgendaEvents []interface{} `json:"getMyAgendaEvents"`
	Typename          string        `json:"__typename"`
}


func (ctx *RequestContext) GetAgendaEvents(variables Variables_GetAgendaEvents) (Data_GetAgendaEvents, error) {
	request := gql.NewQueryRequest[Variables_GetAgendaEvents](
		"query getAgendaEvents($from: DateTime!, $to: DateTime!, $limit: Int!) {\n  student {\n    getMyAgendaEvents(from: $from, to: $to, limit: $limit) {\n      agendaItemContext {\n        entityId\n        entityType\n        __typename\n      }\n      start\n      end\n      label\n      description\n      agendaEventType\n      additionalInfo {\n        key\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetAgendaEvents](ctx, request)
}
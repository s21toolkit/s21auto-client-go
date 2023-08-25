package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetAgendaEvents struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Limit int64  `json:"limit"`
}


type Response_Data_GetAgendaEvents struct {
	Response_Student_GetAgendaEvents Response_Student_GetAgendaEvents `json:"student"`
}

type Response_Student_GetAgendaEvents struct {
	GetMyAgendaEvents []interface{} `json:"getMyAgendaEvents"`
	Typename          string        `json:"__typename"`
}

func (ctx *RequestContext) GetAgendaEvents(variables Request_Variables_GetAgendaEvents) (Response_Data_GetAgendaEvents, error) {
	request := gql.NewQueryRequest[Request_Variables_GetAgendaEvents](
		"query getAgendaEvents($from: DateTime!, $to: DateTime!, $limit: Int!) {\n  student {\n    getMyAgendaEvents(from: $from, to: $to, limit: $limit) {\n      agendaItemContext {\n        entityId\n        entityType\n        __typename\n      }\n      start\n      end\n      label\n      description\n      agendaEventType\n      additionalInfo {\n        key\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetAgendaEvents](ctx, request)
}
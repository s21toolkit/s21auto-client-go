package requests

import "s21client/gql"

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
		"query getAgendaEvents($from: DateTime!, $to: DateTime!, $limit: Int!) {   student {     getMyAgendaEvents(from: $from, to: $to, limit: $limit) {       agendaItemContext {         entityId         entityType         __typename       }       start       end       label       description       agendaEventType       additionalInfo {         key         value         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetAgendaEvents](ctx, request)
}
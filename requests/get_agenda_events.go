package requests

import "s21client/gql"

type Variables_GetAgendaEvents struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Limit int64  `json:"limit"`
}


type Data_GetAgendaEvents struct {
	Student_GetAgendaEvents Student_GetAgendaEvents `json:"student"`
}

type Student_GetAgendaEvents struct {
	GetMyAgendaEvents []interface{} `json:"getMyAgendaEvents"`
	Typename          string        `json:"__typename"`
}

func (ctx *RequestContext) GetAgendaEvents(variables Variables_GetAgendaEvents) (Data_GetAgendaEvents, error) {
	request := gql.NewQueryRequest[Variables_GetAgendaEvents](
		"query getAgendaEvents($from: DateTime!, $to: DateTime!, $limit: Int!) {   student {     getMyAgendaEvents(from: $from, to: $to, limit: $limit) {       agendaItemContext {         entityId         entityType         __typename       }       start       end       label       description       agendaEventType       additionalInfo {         key         value         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetAgendaEvents](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type GetAgendaEvents_Variables struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Limit int64  `json:"limit"`
}


type GetAgendaEvents_Data struct {
	CalendarEventS21 GetAgendaEvents_Data_CalendarEventS21 `json:"calendarEventS21"`
}

type GetAgendaEvents_Data_CalendarEventS21 struct {
	GetMyAgendaEvents []GetAgendaEvents_Data_GetMyAgendaEvent `json:"getMyAgendaEvents"`
	Typename          string             `json:"__typename"`
}

type GetAgendaEvents_Data_GetMyAgendaEvent struct {
	AgendaItemContext GetAgendaEvents_Data_AgendaItemContext `json:"agendaItemContext"`
	Start             string            `json:"start"`
	End               string            `json:"end"`
	Label             string            `json:"label"`
	Description       string            `json:"description"`
	AgendaEventType   string            `json:"agendaEventType"`
	AdditionalInfo    []GetAgendaEvents_Data_AdditionalInfo  `json:"additionalInfo"`
	Typename          string            `json:"__typename"`
}

type GetAgendaEvents_Data_AdditionalInfo struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Typename string `json:"__typename"`
}

type GetAgendaEvents_Data_AgendaItemContext struct {
	EntityID   string `json:"entityId"`
	EntityType string `json:"entityType"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) GetAgendaEvents(variables GetAgendaEvents_Variables) (GetAgendaEvents_Data, error) {
	request := gql.NewQueryRequest[GetAgendaEvents_Variables](
		"query getAgendaEvents($from: DateTime!, $to: DateTime!, $limit: Int!) {\n  calendarEventS21 {\n    getMyAgendaEvents(from: $from, to: $to, limit: $limit) {\n      agendaItemContext {\n        entityId\n        entityType\n        __typename\n      }\n      start\n      end\n      label\n      description\n      agendaEventType\n      additionalInfo {\n        key\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetAgendaEvents_Data](ctx, request)
}
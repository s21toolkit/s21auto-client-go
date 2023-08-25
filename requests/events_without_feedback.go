package requests

import "s21client/gql"

type Request_Variables_EventsWithoutFeedback struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Response_Data_EventsWithoutFeedback struct {
	Response_Student_EventsWithoutFeedback Response_Student_EventsWithoutFeedback `json:"student"`
}

type Response_Student_EventsWithoutFeedback struct {
	GetCalendarEventsWithoutFeedback []interface{} `json:"getCalendarEventsWithoutFeedback"`
	Typename                         string        `json:"__typename"`
}

func (ctx *RequestContext) EventsWithoutFeedback(variables Request_Variables_EventsWithoutFeedback) (Response_Data_EventsWithoutFeedback, error) {
	request := gql.NewQueryRequest[Request_Variables_EventsWithoutFeedback](
		"query eventsWithoutFeedback($from: DateTime!, $to: DateTime!) {   student {     getCalendarEventsWithoutFeedback(from: $from, to: $to) {       id       activity {         eventId         name         endDate         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_EventsWithoutFeedback](ctx, request)
}
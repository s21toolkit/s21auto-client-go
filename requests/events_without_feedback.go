package requests

import "s21client/gql"

type Variables_EventsWithoutFeedback struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_EventsWithoutFeedback struct {
	Student_EventsWithoutFeedback Student_EventsWithoutFeedback `json:"student"`
}

type Student_EventsWithoutFeedback struct {
	GetCalendarEventsWithoutFeedback []interface{} `json:"getCalendarEventsWithoutFeedback"`
	Typename                         string        `json:"__typename"`
}

func (ctx *RequestContext) EventsWithoutFeedback(variables Variables_EventsWithoutFeedback) (Data_EventsWithoutFeedback, error) {
	request := gql.NewQueryRequest[Variables_EventsWithoutFeedback](
		"query eventsWithoutFeedback($from: DateTime!, $to: DateTime!) {   student {     getCalendarEventsWithoutFeedback(from: $from, to: $to) {       id       activity {         eventId         name         endDate         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_EventsWithoutFeedback](ctx, request)
}
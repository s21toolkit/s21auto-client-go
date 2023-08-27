package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_EventsWithoutFeedback struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_EventsWithoutFeedback struct {
	Data_Student_EventsWithoutFeedback Data_Student_EventsWithoutFeedback `json:"student"`
}

type Data_Student_EventsWithoutFeedback struct {
	GetCalendarEventsWithoutFeedback []interface{} `json:"getCalendarEventsWithoutFeedback"`
	Typename                         string        `json:"__typename"`
}


func (ctx *RequestContext) EventsWithoutFeedback(variables Variables_EventsWithoutFeedback) (Data_EventsWithoutFeedback, error) {
	request := gql.NewQueryRequest[Variables_EventsWithoutFeedback](
		"query eventsWithoutFeedback($from: DateTime!, $to: DateTime!) {\n  student {\n    getCalendarEventsWithoutFeedback(from: $from, to: $to) {\n      id\n      activity {\n        eventId\n        name\n        endDate\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_EventsWithoutFeedback](ctx, request)
}
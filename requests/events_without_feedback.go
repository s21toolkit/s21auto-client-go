package requests

import "github.com/s21toolkit/s21client/gql"

type EventsWithoutFeedback_Variables struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type EventsWithoutFeedback_Data struct {
	Student EventsWithoutFeedback_Data_Student `json:"student"`
}

type EventsWithoutFeedback_Data_Student struct {
	GetCalendarEventsWithoutFeedback []interface{} `json:"getCalendarEventsWithoutFeedback"`
	Typename                         string        `json:"__typename"`
}


func (ctx *RequestContext) EventsWithoutFeedback(variables EventsWithoutFeedback_Variables) (EventsWithoutFeedback_Data, error) {
	request := gql.NewQueryRequest[EventsWithoutFeedback_Variables](
		"query eventsWithoutFeedback($from: DateTime!, $to: DateTime!) {\n  student {\n    getCalendarEventsWithoutFeedback(from: $from, to: $to) {\n      id\n      activity {\n        eventId\n        name\n        endDate\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[EventsWithoutFeedback_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CalendarDeleteEventSlot struct {
	EventSlotID string `json:"eventSlotId"`
}


type Response_Data_CalendarDeleteEventSlot struct {
	Response_Student_CalendarDeleteEventSlot Response_Student_CalendarDeleteEventSlot `json:"student"`
}

type Response_Student_CalendarDeleteEventSlot struct {
	DeleteEventSlot bool   `json:"deleteEventSlot"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) CalendarDeleteEventSlot(variables Request_Variables_CalendarDeleteEventSlot) (Response_Data_CalendarDeleteEventSlot, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarDeleteEventSlot](
		"mutation calendarDeleteEventSlot($eventSlotId: ID!) {\n  student {\n    deleteEventSlot(eventSlotId: $eventSlotId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarDeleteEventSlot](ctx, request)
}
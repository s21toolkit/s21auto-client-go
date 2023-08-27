package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarDeleteEventSlot struct {
	EventSlotID string `json:"eventSlotId"`
}


type Data_CalendarDeleteEventSlot struct {
	Data_Student_CalendarDeleteEventSlot Data_Student_CalendarDeleteEventSlot `json:"student"`
}

type Data_Student_CalendarDeleteEventSlot struct {
	DeleteEventSlot bool   `json:"deleteEventSlot"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) CalendarDeleteEventSlot(variables Variables_CalendarDeleteEventSlot) (Data_CalendarDeleteEventSlot, error) {
	request := gql.NewQueryRequest[Variables_CalendarDeleteEventSlot](
		"mutation calendarDeleteEventSlot($eventSlotId: ID!) {\n  student {\n    deleteEventSlot(eventSlotId: $eventSlotId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarDeleteEventSlot](ctx, request)
}
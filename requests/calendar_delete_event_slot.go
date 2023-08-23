package requests

import "s21client/gql"

type Variables_CalendarDeleteEventSlot struct {
	EventSlotID string `json:"eventSlotId"`
}


type Data_CalendarDeleteEventSlot struct {
	Student_CalendarDeleteEventSlot Student_CalendarDeleteEventSlot `json:"student"`
}

type Student_CalendarDeleteEventSlot struct {
	DeleteEventSlot bool   `json:"deleteEventSlot"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) CalendarDeleteEventSlot(variables Variables_CalendarDeleteEventSlot) (Data_CalendarDeleteEventSlot, error) {
	request := gql.NewQueryRequest[Variables_CalendarDeleteEventSlot](
		"mutation calendarDeleteEventSlot($eventSlotId: ID!) {   student {     deleteEventSlot(eventSlotId: $eventSlotId)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_CalendarDeleteEventSlot](ctx, request)
}
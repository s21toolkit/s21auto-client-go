package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarDeleteEventSlot_Variables struct {
	EventSlotID string `json:"eventSlotId"`
}


type CalendarDeleteEventSlot_Data struct {
	Student CalendarDeleteEventSlot_Data_Student `json:"student"`
}

type CalendarDeleteEventSlot_Data_Student struct {
	DeleteEventSlot bool   `json:"deleteEventSlot"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) CalendarDeleteEventSlot(variables CalendarDeleteEventSlot_Variables) (CalendarDeleteEventSlot_Data, error) {
	request := gql.NewQueryRequest[CalendarDeleteEventSlot_Variables](
		"mutation calendarDeleteEventSlot($eventSlotId: ID!) {\n  student {\n    deleteEventSlot(eventSlotId: $eventSlotId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[CalendarDeleteEventSlot_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarAddBookingToEventSlot struct {
	AnswerID           string `json:"answerId"`
	StartTime          string `json:"startTime"`
	WasStaffSlotChosen string `json:"wasStaffSlotChosen"`
	IsOnline           bool   `json:"isOnline"`
}


type Data_CalendarAddBookingToEventSlot struct {
	Data_Student_CalendarAddBookingToEventSlot Data_Student_CalendarAddBookingToEventSlot `json:"student"`
}

type Data_Student_CalendarAddBookingToEventSlot struct {
	Data_AddBookingP2PToEventSlot_CalendarAddBookingToEventSlot Data_AddBookingP2PToEventSlot_CalendarAddBookingToEventSlot `json:"addBookingP2PToEventSlot"`
	Typename                 string                   `json:"__typename"`
}

type Data_AddBookingP2PToEventSlot_CalendarAddBookingToEventSlot struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarAddBookingToEventSlot(variables Variables_CalendarAddBookingToEventSlot) (Data_CalendarAddBookingToEventSlot, error) {
	request := gql.NewQueryRequest[Variables_CalendarAddBookingToEventSlot](
		"mutation calendarAddBookingToEventSlot(\n	$answerId: ID!\n	$startTime: DateTime!\n	$wasStaffSlotChosen: Boolean!\n	$isOnline: Boolean\n) {\n	student {\n		addBookingP2PToEventSlot(\n			answerId: $answerId\n			startTime: $startTime\n			wasStaffSlotChosen: $wasStaffSlotChosen\n			isOnline: $isOnline\n		) {\n			id\n			__typename\n		}\n		__typename\n	}\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarAddBookingToEventSlot](ctx, request)
}
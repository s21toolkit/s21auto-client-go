package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarAddBookingToEventSlot_Variables struct {
	AnswerID           string `json:"answerId"`
	StartTime          string `json:"startTime"`
	WasStaffSlotChosen string `json:"wasStaffSlotChosen"`
	IsOnline           bool   `json:"isOnline"`
}


type CalendarAddBookingToEventSlot_Data struct {
	Student CalendarAddBookingToEventSlot_Data_Student `json:"student"`
}

type CalendarAddBookingToEventSlot_Data_Student struct {
	AddBookingP2PToEventSlot CalendarAddBookingToEventSlot_Data_AddBookingP2PToEventSlot `json:"addBookingP2PToEventSlot"`
	Typename                 string                   `json:"__typename"`
}

type CalendarAddBookingToEventSlot_Data_AddBookingP2PToEventSlot struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarAddBookingToEventSlot(variables CalendarAddBookingToEventSlot_Variables) (CalendarAddBookingToEventSlot_Data, error) {
	request := gql.NewQueryRequest[CalendarAddBookingToEventSlot_Variables](
		"mutation calendarAddBookingToEventSlot(\n	$answerId: ID!\n	$startTime: DateTime!\n	$wasStaffSlotChosen: Boolean!\n	$isOnline: Boolean\n) {\n	student {\n		addBookingP2PToEventSlot(\n			answerId: $answerId\n			startTime: $startTime\n			wasStaffSlotChosen: $wasStaffSlotChosen\n			isOnline: $isOnline\n		) {\n			id\n			__typename\n		}\n		__typename\n	}\n}\n",
		variables,
	)

	return GqlRequest[CalendarAddBookingToEventSlot_Data](ctx, request)
}
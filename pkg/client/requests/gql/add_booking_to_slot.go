package gql

import (
	"school-21/pkg/client/requests"
)

type AddBookingToSlot struct {
	query         string
	operationName string
	variables     requests.BookingRequest
}

func NewAddBookingToSlot(r requests.BookingRequest) *AddBookingToSlot {
	b := &AddBookingToSlot{
		query:         "mutation calendarAddBookingToEventSlot(\n\t$answerId: ID!\n\t$startTime: DateTime!\n\t$wasStaffSlotChosen: Boolean!\n\t$isOnline: Boolean\n) {\n\tstudent {\n\t\taddBookingP2PToEventSlot(\n\t\t\tanswerId: $answerId\n\t\t\tstartTime: $startTime\n\t\t\twasStaffSlotChosen: $wasStaffSlotChosen\n\t\t\tisOnline: $isOnline\n\t\t) {\n\t\t\tid\n\t\t\t__typename\n\t\t}\n\t\t__typename\n\t}\n}\n",
		operationName: "calendarAddBookingToEventSlot",
		variables:     r,
	}

	return b
}

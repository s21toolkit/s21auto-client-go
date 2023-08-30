package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_RemoveP2P struct {
	BookingID string `json:"bookingId"`
}


type Data_RemoveP2P struct {
	Data_Student_RemoveP2P Data_Student_RemoveP2P `json:"student"`
}

type Data_Student_RemoveP2P struct {
	RemoveBookingFromEventSlot bool   `json:"removeBookingFromEventSlot"`
	Typename                   string `json:"__typename"`
}


func (ctx *RequestContext) RemoveP2P(variables Variables_RemoveP2P) (Data_RemoveP2P, error) {
	request := gql.NewQueryRequest[Variables_RemoveP2P](
		"mutation removeP2P($bookingId: ID!) {\n  student {\n    removeBookingFromEventSlot(bookingId: $bookingId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_RemoveP2P](ctx, request)
}
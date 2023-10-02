package requests

import "github.com/s21toolkit/s21client/gql"

type RemoveP2P_Variables struct {
	BookingID string `json:"bookingId"`
}


type RemoveP2P_Data struct {
	Student RemoveP2P_Data_Student `json:"student"`
}

type RemoveP2P_Data_Student struct {
	RemoveBookingFromEventSlot bool   `json:"removeBookingFromEventSlot"`
	Typename                   string `json:"__typename"`
}


func (ctx *RequestContext) RemoveP2P(variables RemoveP2P_Variables) (RemoveP2P_Data, error) {
	request := gql.NewQueryRequest[RemoveP2P_Variables](
		"mutation removeP2P($bookingId: ID!) {\n  student {\n    removeBookingFromEventSlot(bookingId: $bookingId)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[RemoveP2P_Data](ctx, request)
}
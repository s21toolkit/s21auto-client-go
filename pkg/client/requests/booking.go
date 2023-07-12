package requests

import "encoding/json"

func UnmarshalBookingRequest(data []byte) (BookingRequest, error) {
	var r BookingRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BookingRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BookingRequest struct {
	AnswerID           string `json:"answerId"`
	StartTime          string `json:"startTime"`
	WasStaffSlotChosen string `json:"wasStaffSlotChosen"`
	IsOnline           bool   `json:"isOnline"`
}

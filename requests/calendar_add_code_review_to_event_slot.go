package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarAddCodeReviewToEventSlot struct {
	StudentGoalID string `json:"studentGoalId"`
	StartTime     string `json:"startTime"`
}


type Data_CalendarAddCodeReviewToEventSlot struct {
	Data_Student_CalendarAddCodeReviewToEventSlot Data_Student_CalendarAddCodeReviewToEventSlot `json:"student"`
}

type Data_Student_CalendarAddCodeReviewToEventSlot struct {
	Data_AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot Data_AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot `json:"addBookingCodeReviewToEventSlot"`
	Typename                        string                          `json:"__typename"`
}

type Data_AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarAddCodeReviewToEventSlot(variables Variables_CalendarAddCodeReviewToEventSlot) (Data_CalendarAddCodeReviewToEventSlot, error) {
	request := gql.NewQueryRequest[Variables_CalendarAddCodeReviewToEventSlot](
		"mutation calendarAddCodeReviewToEventSlot($studentGoalId: ID!, $startTime: DateTime!) {\n  student {\n    addBookingCodeReviewToEventSlot(\n      studentGoalId: $studentGoalId\n      startTime: $startTime\n    ) {\n      id\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarAddCodeReviewToEventSlot](ctx, request)
}
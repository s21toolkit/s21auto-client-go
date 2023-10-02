package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarAddCodeReviewToEventSlot_Variables struct {
	StudentGoalID string `json:"studentGoalId"`
	StartTime     string `json:"startTime"`
}


type CalendarAddCodeReviewToEventSlot_Data struct {
	Student CalendarAddCodeReviewToEventSlot_Data_Student `json:"student"`
}

type CalendarAddCodeReviewToEventSlot_Data_Student struct {
	AddBookingCodeReviewToEventSlot CalendarAddCodeReviewToEventSlot_Data_AddBookingCodeReviewToEventSlot `json:"addBookingCodeReviewToEventSlot"`
	Typename                        string                          `json:"__typename"`
}

type CalendarAddCodeReviewToEventSlot_Data_AddBookingCodeReviewToEventSlot struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarAddCodeReviewToEventSlot(variables CalendarAddCodeReviewToEventSlot_Variables) (CalendarAddCodeReviewToEventSlot_Data, error) {
	request := gql.NewQueryRequest[CalendarAddCodeReviewToEventSlot_Variables](
		"mutation calendarAddCodeReviewToEventSlot($studentGoalId: ID!, $startTime: DateTime!) {\n  student {\n    addBookingCodeReviewToEventSlot(\n      studentGoalId: $studentGoalId\n      startTime: $startTime\n    ) {\n      id\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[CalendarAddCodeReviewToEventSlot_Data](ctx, request)
}
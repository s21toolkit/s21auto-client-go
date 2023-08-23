package requests

import "s21client/gql"

type Variables_CalendarAddCodeReviewToEventSlot struct {
	StudentGoalID string `json:"studentGoalId"`
	StartTime     string `json:"startTime"`
}


type Data_CalendarAddCodeReviewToEventSlot struct {
	Student_CalendarAddCodeReviewToEventSlot Student_CalendarAddCodeReviewToEventSlot `json:"student"`
}

type Student_CalendarAddCodeReviewToEventSlot struct {
	AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot `json:"addBookingCodeReviewToEventSlot"`
	Typename                        string                          `json:"__typename"`
}

type AddBookingCodeReviewToEventSlot_CalendarAddCodeReviewToEventSlot struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CalendarAddCodeReviewToEventSlot(variables Variables_CalendarAddCodeReviewToEventSlot) (Data_CalendarAddCodeReviewToEventSlot, error) {
	request := gql.NewQueryRequest[Variables_CalendarAddCodeReviewToEventSlot](
		"mutation calendarAddCodeReviewToEventSlot($studentGoalId: ID!, $startTime: DateTime!) {   student {     addBookingCodeReviewToEventSlot(       studentGoalId: $studentGoalId       startTime: $startTime     ) {       id       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_CalendarAddCodeReviewToEventSlot](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetStudentCodeReviews struct {
	StudentGoalID string `json:"studentGoalId"`
}

type Data_CalendarGetStudentCodeReviews struct {
	Student_CalendarGetStudentCodeReviews Student_CalendarGetStudentCodeReviews `json:"student"`
}

type Student_CalendarGetStudentCodeReviews struct {
	GetStudentCodeReviews_CalendarGetStudentCodeReviews GetStudentCodeReviews_CalendarGetStudentCodeReviews `json:"getStudentCodeReviews"`
	Typename                                            string                                              `json:"__typename"`
}

type GetStudentCodeReviews_CalendarGetStudentCodeReviews struct {
	SecondRoundStartDate string `json:"secondRoundStartDate"`
	Typename             string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetStudentCodeReviews(variables Variables_CalendarGetStudentCodeReviews) (Data_CalendarGetStudentCodeReviews, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetStudentCodeReviews](
		"query calendarGetStudentCodeReviews($studentGoalId: ID!) {   student {     getStudentCodeReviews(studentGoalId: $studentGoalId) {       secondRoundStartDate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_CalendarGetStudentCodeReviews](ctx, request)
}

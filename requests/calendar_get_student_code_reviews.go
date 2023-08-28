package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetStudentCodeReviews struct {
	StudentGoalID string `json:"studentGoalId"`
}


type Data_CalendarGetStudentCodeReviews struct {
	Data_Student_CalendarGetStudentCodeReviews Data_Student_CalendarGetStudentCodeReviews `json:"student"`
}

type Data_Student_CalendarGetStudentCodeReviews struct {
	Data_GetStudentCodeReviews_CalendarGetStudentCodeReviews Data_GetStudentCodeReviews_CalendarGetStudentCodeReviews `json:"getStudentCodeReviews"`
	Typename              string                `json:"__typename"`
}

type Data_GetStudentCodeReviews_CalendarGetStudentCodeReviews struct {
	SecondRoundStartDate string `json:"secondRoundStartDate"`
	Typename             string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetStudentCodeReviews(variables Variables_CalendarGetStudentCodeReviews) (Data_CalendarGetStudentCodeReviews, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetStudentCodeReviews](
		"query calendarGetStudentCodeReviews($studentGoalId: ID!) {\n  student {\n    getStudentCodeReviews(studentGoalId: $studentGoalId) {\n      secondRoundStartDate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetStudentCodeReviews](ctx, request)
}
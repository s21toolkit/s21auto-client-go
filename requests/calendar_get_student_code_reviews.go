package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetStudentCodeReviews_Variables struct {
	StudentGoalID string `json:"studentGoalId"`
}


type CalendarGetStudentCodeReviews_Data struct {
	Student CalendarGetStudentCodeReviews_Data_Student `json:"student"`
}

type CalendarGetStudentCodeReviews_Data_Student struct {
	GetStudentCodeReviews CalendarGetStudentCodeReviews_Data_GetStudentCodeReviews `json:"getStudentCodeReviews"`
	Typename              string                `json:"__typename"`
}

type CalendarGetStudentCodeReviews_Data_GetStudentCodeReviews struct {
	SecondRoundStartDate string `json:"secondRoundStartDate"`
	Typename             string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetStudentCodeReviews(variables CalendarGetStudentCodeReviews_Variables) (CalendarGetStudentCodeReviews_Data, error) {
	request := gql.NewQueryRequest[CalendarGetStudentCodeReviews_Variables](
		"query calendarGetStudentCodeReviews($studentGoalId: ID!) {\n  student {\n    getStudentCodeReviews(studentGoalId: $studentGoalId) {\n      secondRoundStartDate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetStudentCodeReviews_Data](ctx, request)
}
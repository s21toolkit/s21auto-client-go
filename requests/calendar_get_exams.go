package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetExams_Variables struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type CalendarGetExams_Data struct {
	Student CalendarGetExams_Data_Student `json:"student"`
}

type CalendarGetExams_Data_Student struct {
	GetExams []interface{} `json:"getExams"`
	Typename string        `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetExams(variables CalendarGetExams_Variables) (CalendarGetExams_Data, error) {
	request := gql.NewQueryRequest[CalendarGetExams_Variables](
		"query calendarGetExams($from: DateTime!, $to: DateTime!) {\n  student {\n    getExams(from: $from, to: $to) {\n      ...CalendarExam\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarExam on Exam {\n  examId\n  eventId\n  beginDate\n  endDate\n  name\n  location\n  maxStudentCount\n  currentStudentsCount\n  updateDate\n  goalId\n  goalName\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetExams_Data](ctx, request)
}
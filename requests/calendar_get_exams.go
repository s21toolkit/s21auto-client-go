package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetExams struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_CalendarGetExams struct {
	Data_Student_CalendarGetExams Data_Student_CalendarGetExams `json:"student"`
}

type Data_Student_CalendarGetExams struct {
	GetExams []interface{} `json:"getExams"`
	Typename string        `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetExams(variables Variables_CalendarGetExams) (Data_CalendarGetExams, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetExams](
		"query calendarGetExams($from: DateTime!, $to: DateTime!) {\n  student {\n    getExams(from: $from, to: $to) {\n      ...CalendarExam\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarExam on Exam {\n  examId\n  eventId\n  beginDate\n  endDate\n  name\n  location\n  maxStudentCount\n  currentStudentsCount\n  updateDate\n  goalId\n  goalName\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetExams](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CalendarGetNameLessStudentTimeslotsForReview struct {
	TaskID string `json:"taskId"`
	From   string `json:"from"`
	To     string `json:"to"`
}


type Response_Data_CalendarGetNameLessStudentTimeslotsForReview struct {
	Response_Student_CalendarGetNameLessStudentTimeslotsForReview Response_Student_CalendarGetNameLessStudentTimeslotsForReview `json:"student"`
}

type Response_Student_CalendarGetNameLessStudentTimeslotsForReview struct {
	Response_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview Response_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview `json:"getNameLessStudentTimeslotsForReview"`
	Typename                             string                               `json:"__typename"`
}

type Response_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview struct {
	CheckDuration      int64              `json:"checkDuration"`
	Response_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview Response_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview `json:"projectReviewsInfo"`
	TimeSlots          []interface{}      `json:"timeSlots"`
	Typename           string             `json:"__typename"`
}

type Response_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetNameLessStudentTimeslotsForReview(variables Request_Variables_CalendarGetNameLessStudentTimeslotsForReview) (Response_Data_CalendarGetNameLessStudentTimeslotsForReview, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetNameLessStudentTimeslotsForReview](
		"query calendarGetNameLessStudentTimeslotsForReview($from: DateTime!, $taskId: ID!, $to: DateTime!) {\n  student {\n    getNameLessStudentTimeslotsForReview(from: $from, taskId: $taskId, to: $to) {\n      checkDuration\n      projectReviewsInfo {\n        ...ProjectReviewsInfo\n        __typename\n      }\n      timeSlots {\n        ...CalendarNameLessTimeslot\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectReviewsInfo on ProjectReviewsInfo {\n  reviewByStudentCount\n  relevantReviewByStudentsCount\n  reviewByInspectionStaffCount\n  relevantReviewByInspectionStaffCount\n  __typename\n}\n\nfragment CalendarNameLessTimeslot on CalendarNamelessTimeSlot {\n  start\n  end\n  validStartTimes\n  staffSlot\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetNameLessStudentTimeslotsForReview](ctx, request)
}
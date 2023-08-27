package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetNameLessStudentTimeslotsForReview struct {
	TaskID string `json:"taskId"`
	From   string `json:"from"`
	To     string `json:"to"`
}


type Data_CalendarGetNameLessStudentTimeslotsForReview struct {
	Data_Student_CalendarGetNameLessStudentTimeslotsForReview Data_Student_CalendarGetNameLessStudentTimeslotsForReview `json:"student"`
}

type Data_Student_CalendarGetNameLessStudentTimeslotsForReview struct {
	Data_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview Data_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview `json:"getNameLessStudentTimeslotsForReview"`
	Typename                             string                               `json:"__typename"`
}

type Data_GetNameLessStudentTimeslotsForReview_CalendarGetNameLessStudentTimeslotsForReview struct {
	CheckDuration      int64              `json:"checkDuration"`
	Data_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview Data_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview `json:"projectReviewsInfo"`
	TimeSlots          []Data_TimeSlot_CalendarGetNameLessStudentTimeslotsForReview         `json:"timeSlots"`
	Typename           string             `json:"__typename"`
}

type Data_ProjectReviewsInfo_CalendarGetNameLessStudentTimeslotsForReview struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type Data_TimeSlot_CalendarGetNameLessStudentTimeslotsForReview struct {
	Start           string   `json:"start"`
	End             string   `json:"end"`
	ValidStartTimes []string `json:"validStartTimes"`
	StaffSlot       bool     `json:"staffSlot"`
	Typename        string   `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetNameLessStudentTimeslotsForReview(variables Variables_CalendarGetNameLessStudentTimeslotsForReview) (Data_CalendarGetNameLessStudentTimeslotsForReview, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetNameLessStudentTimeslotsForReview](
		"query calendarGetNameLessStudentTimeslotsForReview($from: DateTime!, $taskId: ID!, $to: DateTime!) {\n  student {\n    getNameLessStudentTimeslotsForReview(from: $from, taskId: $taskId, to: $to) {\n      checkDuration\n      projectReviewsInfo {\n        ...ProjectReviewsInfo\n        __typename\n      }\n      timeSlots {\n        ...CalendarNameLessTimeslot\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectReviewsInfo on ProjectReviewsInfo {\n  reviewByStudentCount\n  relevantReviewByStudentsCount\n  reviewByInspectionStaffCount\n  relevantReviewByInspectionStaffCount\n  __typename\n}\n\nfragment CalendarNameLessTimeslot on CalendarNamelessTimeSlot {\n  start\n  end\n  validStartTimes\n  staffSlot\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetNameLessStudentTimeslotsForReview](ctx, request)
}
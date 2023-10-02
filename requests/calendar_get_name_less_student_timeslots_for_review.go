package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetNameLessStudentTimeslotsForReview_Variables struct {
	TaskID string `json:"taskId"`
	From   string `json:"from"`
	To     string `json:"to"`
}


type CalendarGetNameLessStudentTimeslotsForReview_Data struct {
	Student CalendarGetNameLessStudentTimeslotsForReview_Data_Student `json:"student"`
}

type CalendarGetNameLessStudentTimeslotsForReview_Data_Student struct {
	GetNameLessStudentTimeslotsForReview CalendarGetNameLessStudentTimeslotsForReview_Data_GetNameLessStudentTimeslotsForReview `json:"getNameLessStudentTimeslotsForReview"`
	Typename                             string                               `json:"__typename"`
}

type CalendarGetNameLessStudentTimeslotsForReview_Data_GetNameLessStudentTimeslotsForReview struct {
	CheckDuration      int64              `json:"checkDuration"`
	ProjectReviewsInfo CalendarGetNameLessStudentTimeslotsForReview_Data_ProjectReviewsInfo `json:"projectReviewsInfo"`
	TimeSlots          []CalendarGetNameLessStudentTimeslotsForReview_Data_TimeSlot         `json:"timeSlots"`
	Typename           string             `json:"__typename"`
}

type CalendarGetNameLessStudentTimeslotsForReview_Data_ProjectReviewsInfo struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type CalendarGetNameLessStudentTimeslotsForReview_Data_TimeSlot struct {
	Start           string   `json:"start"`
	End             string   `json:"end"`
	ValidStartTimes []string `json:"validStartTimes"`
	StaffSlot       bool     `json:"staffSlot"`
	Typename        string   `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetNameLessStudentTimeslotsForReview(variables CalendarGetNameLessStudentTimeslotsForReview_Variables) (CalendarGetNameLessStudentTimeslotsForReview_Data, error) {
	request := gql.NewQueryRequest[CalendarGetNameLessStudentTimeslotsForReview_Variables](
		"query calendarGetNameLessStudentTimeslotsForReview($from: DateTime!, $taskId: ID!, $to: DateTime!) {\n  student {\n    getNameLessStudentTimeslotsForReview(from: $from, taskId: $taskId, to: $to) {\n      checkDuration\n      projectReviewsInfo {\n        ...ProjectReviewsInfo\n        __typename\n      }\n      timeSlots {\n        ...CalendarNameLessTimeslot\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectReviewsInfo on ProjectReviewsInfo {\n  reviewByStudentCount\n  relevantReviewByStudentsCount\n  reviewByInspectionStaffCount\n  relevantReviewByInspectionStaffCount\n  __typename\n}\n\nfragment CalendarNameLessTimeslot on CalendarNamelessTimeSlot {\n  start\n  end\n  validStartTimes\n  staffSlot\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetNameLessStudentTimeslotsForReview_Data](ctx, request)
}
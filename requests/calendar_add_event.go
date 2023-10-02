package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarAddEvent_Variables struct {
	Start string `json:"start"`
	End   string `json:"end"`
}


type CalendarAddEvent_Data struct {
	Student CalendarAddEvent_Data_Student `json:"student"`
}

type CalendarAddEvent_Data_Student struct {
	AddEventToTimetable []CalendarAddEvent_Data_AddEventToTimetable `json:"addEventToTimetable"`
	Typename            string                `json:"__typename"`
}

type CalendarAddEvent_Data_AddEventToTimetable struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         interface{}   `json:"eventCode"`
	EventSlots        []CalendarAddEvent_Data_EventSlot   `json:"eventSlots"`
	Bookings          []interface{} `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Activity          interface{}   `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type CalendarAddEvent_Data_EventSlot struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarAddEvent(variables CalendarAddEvent_Variables) (CalendarAddEvent_Data, error) {
	request := gql.NewQueryRequest[CalendarAddEvent_Variables](
		"mutation calendarAddEvent($start: DateTime!, $end: DateTime!) {\n  student {\n    addEventToTimetable(start: $start, end: $end) {\n      ...CalendarEvent\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarEvent on CalendarEvent {\n  id\n  start\n  end\n  description\n  eventType\n  eventCode\n  eventSlots {\n    id\n    type\n    start\n    end\n    __typename\n  }\n  bookings {\n    ...CalendarReviewBooking\n    __typename\n  }\n  exam {\n    ...CalendarEventExam\n    __typename\n  }\n  studentCodeReview {\n    studentGoalId\n    __typename\n  }\n  activity {\n    ...CalendarEventActivity\n    studentFeedback {\n      id\n      rating\n      comment\n      __typename\n    }\n    status\n    activityType\n    isMandatory\n    isWaitListActive\n    isVisible\n    comments {\n      type\n      createTs\n      comment\n      __typename\n    }\n    organizers {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n  goals {\n    goalId\n    goalName\n    __typename\n  }\n  penalty {\n    ...Penalty\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarEventExam on Exam {\n  examId\n  eventId\n  beginDate\n  endDate\n  name\n  location\n  currentStudentsCount\n  maxStudentCount\n  updateDate\n  goalId\n  goalName\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment CalendarEventActivity on ActivityEvent {\n  activityEventId\n  eventId\n  name\n  beginDate\n  endDate\n  isRegistered\n  description\n  currentStudentsCount\n  maxStudentCount\n  location\n  updateDate\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarAddEvent_Data](ctx, request)
}
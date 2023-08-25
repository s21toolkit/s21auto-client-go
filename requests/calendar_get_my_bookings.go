package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CalendarGetMyBookings struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Response_Data_CalendarGetMyBookings struct {
	Response_Student_CalendarGetMyBookings Response_Student_CalendarGetMyBookings `json:"student"`
}

type Response_Student_CalendarGetMyBookings struct {
	GetMyCalendarBookings []Response_GetMyCalendarBooking_CalendarGetMyBookings `json:"getMyCalendarBookings"`
	Typename              string                 `json:"__typename"`
}

type Response_GetMyCalendarBooking_CalendarGetMyBookings struct {
	ID                string            `json:"id"`
	AnswerID          string            `json:"answerId"`
	EventSlotID       string            `json:"eventSlotId"`
	Response_Task_CalendarGetMyBookings              Response_Task_CalendarGetMyBookings              `json:"task"`
	Response_EventSlot_CalendarGetMyBookings         Response_EventSlot_CalendarGetMyBookings         `json:"eventSlot"`
	VerifierUser      Response_User_CalendarGetMyBookings              `json:"verifierUser"`
	Response_VerifiableStudent_CalendarGetMyBookings Response_VerifiableStudent_CalendarGetMyBookings `json:"verifiableStudent"`
	BookingStatus     string            `json:"bookingStatus"`
	Team              interface{}       `json:"team"`
	IsOnline          bool              `json:"isOnline"`
	VcLinkURL         interface{}       `json:"vcLinkUrl"`
	Typename          string            `json:"__typename"`
}

type Response_EventSlot_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Response_Event_CalendarGetMyBookings    Response_Event_CalendarGetMyBookings  `json:"event"`
	Typename string `json:"__typename"`
}

type Response_Event_CalendarGetMyBookings struct {
	EventUserRole string `json:"eventUserRole"`
	Typename      string `json:"__typename"`
}

type Response_Task_CalendarGetMyBookings struct {
	ID                              string                          `json:"id"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	Response_StudentTaskAdditionalAttributes_CalendarGetMyBookings Response_StudentTaskAdditionalAttributes_CalendarGetMyBookings `json:"studentTaskAdditionalAttributes"`
	AssignmentType                  string                          `json:"assignmentType"`
	Typename                        string                          `json:"__typename"`
}

type Response_StudentTaskAdditionalAttributes_CalendarGetMyBookings struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type Response_VerifiableStudent_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Response_User_CalendarGetMyBookings     Response_User_CalendarGetMyBookings   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetMyBookings(variables Request_Variables_CalendarGetMyBookings) (Response_Data_CalendarGetMyBookings, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetMyBookings](
		"query calendarGetMyBookings($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarBookings(from: $from, to: $to) {\n      ...CalendarReviewBooking\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetMyBookings](ctx, request)
}
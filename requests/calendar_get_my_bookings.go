package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetMyBookings struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_CalendarGetMyBookings struct {
	Data_Student_CalendarGetMyBookings Data_Student_CalendarGetMyBookings `json:"student"`
}

type Data_Student_CalendarGetMyBookings struct {
	GetMyCalendarBookings []Data_GetMyCalendarBooking_CalendarGetMyBookings `json:"getMyCalendarBookings"`
	Typename              string                 `json:"__typename"`
}

type Data_GetMyCalendarBooking_CalendarGetMyBookings struct {
	ID                string            `json:"id"`
	AnswerID          string            `json:"answerId"`
	EventSlotID       string            `json:"eventSlotId"`
	Data_Task_CalendarGetMyBookings              Data_Task_CalendarGetMyBookings              `json:"task"`
	Data_EventSlot_CalendarGetMyBookings         Data_EventSlot_CalendarGetMyBookings         `json:"eventSlot"`
	VerifierUser      Data_User_CalendarGetMyBookings              `json:"verifierUser"`
	Data_VerifiableStudent_CalendarGetMyBookings Data_VerifiableStudent_CalendarGetMyBookings `json:"verifiableStudent"`
	BookingStatus     string            `json:"bookingStatus"`
	Team              interface{}       `json:"team"`
	IsOnline          bool              `json:"isOnline"`
	VcLinkURL         interface{}       `json:"vcLinkUrl"`
	Typename          string            `json:"__typename"`
}

type Data_EventSlot_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Data_Event_CalendarGetMyBookings    Data_Event_CalendarGetMyBookings  `json:"event"`
	Typename string `json:"__typename"`
}

type Data_Event_CalendarGetMyBookings struct {
	EventUserRole string `json:"eventUserRole"`
	Typename      string `json:"__typename"`
}

type Data_Task_CalendarGetMyBookings struct {
	ID                              string                          `json:"id"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	Data_StudentTaskAdditionalAttributes_CalendarGetMyBookings Data_StudentTaskAdditionalAttributes_CalendarGetMyBookings `json:"studentTaskAdditionalAttributes"`
	AssignmentType                  string                          `json:"assignmentType"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_CalendarGetMyBookings struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type Data_VerifiableStudent_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Data_User_CalendarGetMyBookings     Data_User_CalendarGetMyBookings   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_CalendarGetMyBookings struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetMyBookings(variables Variables_CalendarGetMyBookings) (Data_CalendarGetMyBookings, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetMyBookings](
		"query calendarGetMyBookings($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarBookings(from: $from, to: $to) {\n      ...CalendarReviewBooking\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetMyBookings](ctx, request)
}
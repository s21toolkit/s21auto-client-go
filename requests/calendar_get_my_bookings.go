package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetMyBookings_Variables struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type CalendarGetMyBookings_Data struct {
	Student CalendarGetMyBookings_Data_Student `json:"student"`
}

type CalendarGetMyBookings_Data_Student struct {
	GetMyCalendarBookings []CalendarGetMyBookings_Data_GetMyCalendarBooking `json:"getMyCalendarBookings"`
	Typename              string                 `json:"__typename"`
}

type CalendarGetMyBookings_Data_GetMyCalendarBooking struct {
	ID                string            `json:"id"`
	AnswerID          string            `json:"answerId"`
	EventSlotID       string            `json:"eventSlotId"`
	Task              CalendarGetMyBookings_Data_Task              `json:"task"`
	EventSlot         CalendarGetMyBookings_Data_EventSlot         `json:"eventSlot"`
	VerifierUser      CalendarGetMyBookings_Data_User              `json:"verifierUser"`
	VerifiableStudent CalendarGetMyBookings_Data_VerifiableStudent `json:"verifiableStudent"`
	BookingStatus     string            `json:"bookingStatus"`
	Team              interface{}       `json:"team"`
	IsOnline          bool              `json:"isOnline"`
	VcLinkURL         interface{}       `json:"vcLinkUrl"`
	Typename          string            `json:"__typename"`
}

type CalendarGetMyBookings_Data_EventSlot struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Event    CalendarGetMyBookings_Data_Event  `json:"event"`
	Typename string `json:"__typename"`
}

type CalendarGetMyBookings_Data_Event struct {
	EventUserRole string `json:"eventUserRole"`
	Typename      string `json:"__typename"`
}

type CalendarGetMyBookings_Data_Task struct {
	ID                              string                          `json:"id"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	StudentTaskAdditionalAttributes CalendarGetMyBookings_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	AssignmentType                  string                          `json:"assignmentType"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetMyBookings_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type CalendarGetMyBookings_Data_VerifiableStudent struct {
	ID       string `json:"id"`
	User     CalendarGetMyBookings_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type CalendarGetMyBookings_Data_User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetMyBookings(variables CalendarGetMyBookings_Variables) (CalendarGetMyBookings_Data, error) {
	request := gql.NewQueryRequest[CalendarGetMyBookings_Variables](
		"query calendarGetMyBookings($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarBookings(from: $from, to: $to) {\n      ...CalendarReviewBooking\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetMyBookings_Data](ctx, request)
}
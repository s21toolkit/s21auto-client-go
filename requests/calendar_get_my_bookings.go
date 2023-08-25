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
	GetMyCalendarBookings []interface{} `json:"getMyCalendarBookings"`
	Typename              string        `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetMyBookings(variables Request_Variables_CalendarGetMyBookings) (Response_Data_CalendarGetMyBookings, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetMyBookings](
		"query calendarGetMyBookings($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarBookings(from: $from, to: $to) {\n      ...CalendarReviewBooking\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetMyBookings](ctx, request)
}
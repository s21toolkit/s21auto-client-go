package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CalendarGetMyReviews struct {
	Limit int64 `json:"limit"`
}


type Response_Data_CalendarGetMyReviews struct {
	Response_Student_CalendarGetMyReviews Response_Student_CalendarGetMyReviews `json:"student"`
}

type Response_Student_CalendarGetMyReviews struct {
	GetMyUpcomingBookings []interface{} `json:"getMyUpcomingBookings"`
	Typename              string        `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetMyReviews(variables Request_Variables_CalendarGetMyReviews) (Response_Data_CalendarGetMyReviews, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetMyReviews](
		"query calendarGetMyReviews($to: DateTime, $limit: Int) {\n  student {\n    getMyUpcomingBookings(to: $to, limit: $limit) {\n      ...Review\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment Review on CalendarBooking {\n  id\n  answerId\n  eventSlot {\n    id\n    start\n    end\n    __typename\n  }\n  task {\n    id\n    title\n    assignmentType\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...UserInBooking\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...UserInBooking\n      __typename\n    }\n    __typename\n  }\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  bookingStatus\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment UserInBooking on User {\n  id\n  login\n  avatarUrl\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetMyReviews](ctx, request)
}
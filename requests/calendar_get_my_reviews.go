package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetMyReviews struct {
	Limit int64 `json:"limit"`
}


type Data_CalendarGetMyReviews struct {
	Data_Student_CalendarGetMyReviews Data_Student_CalendarGetMyReviews `json:"student"`
}

type Data_Student_CalendarGetMyReviews struct {
	GetMyUpcomingBookings []Data_GetMyUpcomingBooking_CalendarGetMyReviews `json:"getMyUpcomingBookings"`
	Typename              string                 `json:"__typename"`
}

type Data_GetMyUpcomingBooking_CalendarGetMyReviews struct {
	ID                string       `json:"id"`
	AnswerID          interface{}  `json:"answerId"`
	Data_EventSlot_CalendarGetMyReviews         Data_EventSlot_CalendarGetMyReviews    `json:"eventSlot"`
	Task              interface{}  `json:"task"`
	Data_VerifierUser_CalendarGetMyReviews      Data_VerifierUser_CalendarGetMyReviews `json:"verifierUser"`
	VerifiableStudent interface{}  `json:"verifiableStudent"`
	Team              interface{}  `json:"team"`
	BookingStatus     string       `json:"bookingStatus"`
	IsOnline          bool         `json:"isOnline"`
	VcLinkURL         interface{}  `json:"vcLinkUrl"`
	Typename          string       `json:"__typename"`
}

type Data_EventSlot_CalendarGetMyReviews struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}

type Data_VerifierUser_CalendarGetMyReviews struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_CalendarGetMyReviews Data_UserExperience_CalendarGetMyReviews `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_CalendarGetMyReviews struct {
	Data_Level_CalendarGetMyReviews    Data_Level_CalendarGetMyReviews  `json:"level"`
	Typename string `json:"__typename"`
}

type Data_Level_CalendarGetMyReviews struct {
	ID       int64  `json:"id"`
	Data_Range_CalendarGetMyReviews    Data_Range_CalendarGetMyReviews  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_CalendarGetMyReviews struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetMyReviews(variables Variables_CalendarGetMyReviews) (Data_CalendarGetMyReviews, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetMyReviews](
		"query calendarGetMyReviews($to: DateTime, $limit: Int) {\n  student {\n    getMyUpcomingBookings(to: $to, limit: $limit) {\n      ...Review\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment Review on CalendarBooking {\n  id\n  answerId\n  eventSlot {\n    id\n    start\n    end\n    __typename\n  }\n  task {\n    id\n    title\n    assignmentType\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...UserInBooking\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...UserInBooking\n      __typename\n    }\n    __typename\n  }\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  bookingStatus\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment UserInBooking on User {\n  id\n  login\n  avatarUrl\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetMyReviews](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetMyReviews_Variables struct {
	Limit int64 `json:"limit"`
}


type CalendarGetMyReviews_Data struct {
	Student CalendarGetMyReviews_Data_Student `json:"student"`
}

type CalendarGetMyReviews_Data_Student struct {
	GetMyUpcomingBookings []CalendarGetMyReviews_Data_GetMyUpcomingBooking `json:"getMyUpcomingBookings"`
	Typename              string                 `json:"__typename"`
}

type CalendarGetMyReviews_Data_GetMyUpcomingBooking struct {
	ID                string             `json:"id"`
	AnswerID          *string            `json:"answerId"`
	EventSlot         CalendarGetMyReviews_Data_EventSlot          `json:"eventSlot"`
	Task              *CalendarGetMyReviews_Data_Task              `json:"task"`
	VerifierUser      CalendarGetMyReviews_Data_User               `json:"verifierUser"`
	VerifiableStudent *CalendarGetMyReviews_Data_VerifiableStudent `json:"verifiableStudent"`
	Team              *CalendarGetMyReviews_Data_Team              `json:"team"`
	BookingStatus     string             `json:"bookingStatus"`
	IsOnline          bool               `json:"isOnline"`
	VcLinkURL         interface{}        `json:"vcLinkUrl"`
	Typename          string             `json:"__typename"`
}

type CalendarGetMyReviews_Data_EventSlot struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}

type CalendarGetMyReviews_Data_Task struct {
	ID                              string                          `json:"id"`
	Title                           string                          `json:"title"`
	AssignmentType                  string                          `json:"assignmentType"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	StudentTaskAdditionalAttributes CalendarGetMyReviews_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetMyReviews_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type CalendarGetMyReviews_Data_Team struct {
	ID                 string        `json:"id"`
	TeamLead           CalendarGetMyReviews_Data_TeamLead      `json:"teamLead"`
	Members            []CalendarGetMyReviews_Data_TeamLead    `json:"members"`
	InvitedUsers       []interface{} `json:"invitedUsers"`
	TeamName           string        `json:"teamName"`
	TeamStatus         string        `json:"teamStatus"`
	MinTeamMemberCount int64         `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64         `json:"maxTeamMemberCount"`
	Typename           string        `json:"__typename"`
}

type CalendarGetMyReviews_Data_TeamLead struct {
	ID             string                 `json:"id"`
	AvatarURL      string                 `json:"avatarUrl"`
	Login          string                 `json:"login"`
	UserExperience CalendarGetMyReviews_Data_TeamLeadUserExperience `json:"userExperience"`
	Typename       string                 `json:"__typename"`
}

type CalendarGetMyReviews_Data_TeamLeadUserExperience struct {
	Level            CalendarGetMyReviews_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type CalendarGetMyReviews_Data_Level struct {
	ID       int64  `json:"id"`
	Range    CalendarGetMyReviews_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type CalendarGetMyReviews_Data_Range struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type CalendarGetMyReviews_Data_VerifiableStudent struct {
	ID       string `json:"id"`
	User     CalendarGetMyReviews_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type CalendarGetMyReviews_Data_User struct {
	ID             string                     `json:"id"`
	Login          string                     `json:"login"`
	AvatarURL      string                     `json:"avatarUrl"`
	UserExperience CalendarGetMyReviews_Data_VerifierUserUserExperience `json:"userExperience"`
	Typename       string                     `json:"__typename"`
}

type CalendarGetMyReviews_Data_VerifierUserUserExperience struct {
	Level    CalendarGetMyReviews_Data_Level  `json:"level"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetMyReviews(variables CalendarGetMyReviews_Variables) (CalendarGetMyReviews_Data, error) {
	request := gql.NewQueryRequest[CalendarGetMyReviews_Variables](
		"query calendarGetMyReviews($to: DateTime, $limit: Int) {\n  student {\n    getMyUpcomingBookings(to: $to, limit: $limit) {\n      ...Review\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment Review on CalendarBooking {\n  id\n  answerId\n  eventSlot {\n    id\n    start\n    end\n    __typename\n  }\n  task {\n    id\n    title\n    assignmentType\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...UserInBooking\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...UserInBooking\n      __typename\n    }\n    __typename\n  }\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  bookingStatus\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment UserInBooking on User {\n  id\n  login\n  avatarUrl\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetMyReviews_Data](ctx, request)
}
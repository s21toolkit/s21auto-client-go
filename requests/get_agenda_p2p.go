package requests

import "github.com/s21toolkit/s21client/gql"

type GetAgendaP2P_Variables struct {
	BookingID string `json:"bookingId"`
}


type GetAgendaP2P_Data struct {
	Student GetAgendaP2P_Data_Student `json:"student"`
}

type GetAgendaP2P_Data_Student struct {
	GetEnrichedBooking GetAgendaP2P_Data_GetEnrichedBooking `json:"getEnrichedBooking"`
	Typename           string             `json:"__typename"`
}

type GetAgendaP2P_Data_GetEnrichedBooking struct {
	ID                string             `json:"id"`
	EventSlot         GetAgendaP2P_Data_EventSlot          `json:"eventSlot"`
	Task              *GetAgendaP2P_Data_Task              `json:"task"`
	VerifierUser      GetAgendaP2P_Data_VerifierUser       `json:"verifierUser"`
	VerifiableStudent *GetAgendaP2P_Data_VerifiableStudent `json:"verifiableStudent"`
	AnswerID          *string            `json:"answerId"`
	Team              *GetAgendaP2P_Data_Team              `json:"team"`
	BookingStatus     string             `json:"bookingStatus"`
	IsOnline          bool               `json:"isOnline"`
	VcLinkURL         interface{}        `json:"vcLinkUrl"`
	Typename          string             `json:"__typename"`
}

type GetAgendaP2P_Data_EventSlot struct {
	Start    string `json:"start"`
	Typename string `json:"__typename"`
}

type GetAgendaP2P_Data_Task struct {
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	AssignmentType                  string                          `json:"assignmentType"`
	StudentTaskAdditionalAttributes GetAgendaP2P_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type GetAgendaP2P_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type GetAgendaP2P_Data_Team struct {
	TeamName string         `json:"teamName"`
	TeamLead GetAgendaP2P_Data_VerifierUser   `json:"teamLead"`
	Members  []GetAgendaP2P_Data_VerifierUser `json:"members"`
	Typename string         `json:"__typename"`
}

type GetAgendaP2P_Data_VerifierUser struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience GetAgendaP2P_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type GetAgendaP2P_Data_UserExperience struct {
	Level    GetAgendaP2P_Data_Level  `json:"level"`
	Typename string `json:"__typename"`
}

type GetAgendaP2P_Data_Level struct {
	Range    GetAgendaP2P_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type GetAgendaP2P_Data_Range struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type GetAgendaP2P_Data_VerifiableStudent struct {
	User     GetAgendaP2P_Data_VerifierUser `json:"user"`
	Typename string       `json:"__typename"`
}


func (ctx *RequestContext) GetAgendaP2P(variables GetAgendaP2P_Variables) (GetAgendaP2P_Data, error) {
	request := gql.NewQueryRequest[GetAgendaP2P_Variables](
		"query getAgendaP2P($bookingId: ID!) {\n  student {\n    getEnrichedBooking(bookingId: $bookingId) {\n      id\n      eventSlot {\n        start\n        __typename\n      }\n      task {\n        goalId\n        goalName\n        assignmentType\n        studentTaskAdditionalAttributes {\n          cookiesCount\n          __typename\n        }\n        __typename\n      }\n      verifierUser {\n        id\n        login\n        avatarUrl\n        userExperience {\n          level {\n            range {\n              levelCode\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      verifiableStudent {\n        user {\n          id\n          login\n          avatarUrl\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      answerId\n      team {\n        teamName\n        teamLead {\n          id\n          avatarUrl\n          login\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        members {\n          id\n          avatarUrl\n          login\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      bookingStatus\n      isOnline\n      vcLinkUrl\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetAgendaP2P_Data](ctx, request)
}
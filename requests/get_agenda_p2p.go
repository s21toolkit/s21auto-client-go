package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetAgendaP2P struct {
	BookingID string `json:"bookingId"`
}


type Data_GetAgendaP2P struct {
	Data_Student_GetAgendaP2P Data_Student_GetAgendaP2P `json:"student"`
}

type Data_Student_GetAgendaP2P struct {
	Data_GetEnrichedBooking_GetAgendaP2P Data_GetEnrichedBooking_GetAgendaP2P `json:"getEnrichedBooking"`
	Typename           string             `json:"__typename"`
}

type Data_GetEnrichedBooking_GetAgendaP2P struct {
	ID                string             `json:"id"`
	Data_EventSlot_GetAgendaP2P         Data_EventSlot_GetAgendaP2P          `json:"eventSlot"`
	Data_Task_GetAgendaP2P              *Data_Task_GetAgendaP2P              `json:"task"`
	Data_VerifierUser_GetAgendaP2P      Data_VerifierUser_GetAgendaP2P       `json:"verifierUser"`
	Data_VerifiableStudent_GetAgendaP2P *Data_VerifiableStudent_GetAgendaP2P `json:"verifiableStudent"`
	AnswerID          *string            `json:"answerId"`
	Data_Team_GetAgendaP2P              *Data_Team_GetAgendaP2P              `json:"team"`
	BookingStatus     string             `json:"bookingStatus"`
	IsOnline          bool               `json:"isOnline"`
	VcLinkURL         interface{}        `json:"vcLinkUrl"`
	Typename          string             `json:"__typename"`
}

type Data_EventSlot_GetAgendaP2P struct {
	Start    string `json:"start"`
	Typename string `json:"__typename"`
}

type Data_Task_GetAgendaP2P struct {
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	AssignmentType                  string                          `json:"assignmentType"`
	Data_StudentTaskAdditionalAttributes_GetAgendaP2P Data_StudentTaskAdditionalAttributes_GetAgendaP2P `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_GetAgendaP2P struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type Data_Team_GetAgendaP2P struct {
	TeamName string         `json:"teamName"`
	TeamLead Data_VerifierUser_GetAgendaP2P   `json:"teamLead"`
	Members  []Data_VerifierUser_GetAgendaP2P `json:"members"`
	Typename string         `json:"__typename"`
}

type Data_VerifierUser_GetAgendaP2P struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_GetAgendaP2P Data_UserExperience_GetAgendaP2P `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_GetAgendaP2P struct {
	Data_Level_GetAgendaP2P    Data_Level_GetAgendaP2P  `json:"level"`
	Typename string `json:"__typename"`
}

type Data_Level_GetAgendaP2P struct {
	Data_Range_GetAgendaP2P    Data_Range_GetAgendaP2P  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_GetAgendaP2P struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Data_VerifiableStudent_GetAgendaP2P struct {
	User     Data_VerifierUser_GetAgendaP2P `json:"user"`
	Typename string       `json:"__typename"`
}


func (ctx *RequestContext) GetAgendaP2P(variables Variables_GetAgendaP2P) (Data_GetAgendaP2P, error) {
	request := gql.NewQueryRequest[Variables_GetAgendaP2P](
		"query getAgendaP2P($bookingId: ID!) {\n  student {\n    getEnrichedBooking(bookingId: $bookingId) {\n      id\n      eventSlot {\n        start\n        __typename\n      }\n      task {\n        goalId\n        goalName\n        assignmentType\n        studentTaskAdditionalAttributes {\n          cookiesCount\n          __typename\n        }\n        __typename\n      }\n      verifierUser {\n        id\n        login\n        avatarUrl\n        userExperience {\n          level {\n            range {\n              levelCode\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      verifiableStudent {\n        user {\n          id\n          login\n          avatarUrl\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      answerId\n      team {\n        teamName\n        teamLead {\n          id\n          avatarUrl\n          login\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        members {\n          id\n          avatarUrl\n          login\n          userExperience {\n            level {\n              range {\n                levelCode\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      bookingStatus\n      isOnline\n      vcLinkUrl\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetAgendaP2P](ctx, request)
}
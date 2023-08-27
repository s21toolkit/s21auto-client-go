package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_SendInvitation struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type Data_SendInvitation struct {
	Student Data_DataStudent_SendInvitation `json:"student"`
}

type Data_DataStudent_SendInvitation struct {
	Data_SendInvitation_SendInvitation Data_SendInvitation_SendInvitation `json:"sendInvitation"`
	Typename       string         `json:"__typename"`
}

type Data_SendInvitation_SendInvitation struct {
	Student          Data_SendInvitationStudent_SendInvitation `json:"student"`
	InvitationStatus string                `json:"invitationStatus"`
	Typename         string                `json:"__typename"`
}

type Data_SendInvitationStudent_SendInvitation struct {
	ID       string `json:"id"`
	Data_User_SendInvitation     Data_User_SendInvitation   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_SendInvitation struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_SendInvitation Data_UserExperience_SendInvitation `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_SendInvitation struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Data_Level_SendInvitation            Data_Level_SendInvitation  `json:"level"`
	Typename         string `json:"__typename"`
}

type Data_Level_SendInvitation struct {
	ID       int64  `json:"id"`
	Data_Range_SendInvitation    Data_Range_SendInvitation  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_SendInvitation struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) SendInvitation(variables Variables_SendInvitation) (Data_SendInvitation, error) {
	request := gql.NewQueryRequest[Variables_SendInvitation](
		"mutation sendInvitation($teamId: UUID!, $userId: ID!) {\n  student {\n    sendInvitation(teamId: $teamId, userId: $userId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_SendInvitation](ctx, request)
}
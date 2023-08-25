package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_SendInvitation struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type Response_Data_SendInvitation struct {
	Student Response_DataStudent_SendInvitation `json:"student"`
}

type Response_DataStudent_SendInvitation struct {
	Response_SendInvitation_SendInvitation Response_SendInvitation_SendInvitation `json:"sendInvitation"`
	Typename       string         `json:"__typename"`
}

type Response_SendInvitation_SendInvitation struct {
	Student          Response_SendInvitationStudent_SendInvitation `json:"student"`
	InvitationStatus string                `json:"invitationStatus"`
	Typename         string                `json:"__typename"`
}

type Response_SendInvitationStudent_SendInvitation struct {
	ID       string `json:"id"`
	Response_User_SendInvitation     Response_User_SendInvitation   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_SendInvitation struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Response_UserExperience_SendInvitation Response_UserExperience_SendInvitation `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Response_UserExperience_SendInvitation struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Response_Level_SendInvitation            Response_Level_SendInvitation  `json:"level"`
	Typename         string `json:"__typename"`
}

type Response_Level_SendInvitation struct {
	ID       int64  `json:"id"`
	Response_Range_SendInvitation    Response_Range_SendInvitation  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_SendInvitation struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) SendInvitation(variables Request_Variables_SendInvitation) (Response_Data_SendInvitation, error) {
	request := gql.NewQueryRequest[Request_Variables_SendInvitation](
		"mutation sendInvitation($teamId: UUID!, $userId: ID!) {\n  student {\n    sendInvitation(teamId: $teamId, userId: $userId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_SendInvitation](ctx, request)
}
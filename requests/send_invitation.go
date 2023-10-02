package requests

import "github.com/s21toolkit/s21client/gql"

type SendInvitation_Variables struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type SendInvitation_Data struct {
	Student SendInvitation_Data_DataStudent `json:"student"`
}

type SendInvitation_Data_DataStudent struct {
	SendInvitation SendInvitation_Data_SendInvitation `json:"sendInvitation"`
	Typename       string         `json:"__typename"`
}

type SendInvitation_Data_SendInvitation struct {
	Student          SendInvitation_Data_SendInvitationStudent `json:"student"`
	InvitationStatus string                `json:"invitationStatus"`
	Typename         string                `json:"__typename"`
}

type SendInvitation_Data_SendInvitationStudent struct {
	ID       string `json:"id"`
	User     SendInvitation_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type SendInvitation_Data_User struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience SendInvitation_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type SendInvitation_Data_UserExperience struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level            SendInvitation_Data_Level  `json:"level"`
	Typename         string `json:"__typename"`
}

type SendInvitation_Data_Level struct {
	ID       int64  `json:"id"`
	Range    SendInvitation_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type SendInvitation_Data_Range struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) SendInvitation(variables SendInvitation_Variables) (SendInvitation_Data, error) {
	request := gql.NewQueryRequest[SendInvitation_Variables](
		"mutation sendInvitation($teamId: UUID!, $userId: ID!) {\n  student {\n    sendInvitation(teamId: $teamId, userId: $userId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[SendInvitation_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CancelInvitation struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type Data_CancelInvitation struct {
	Student Data_DataStudent_CancelInvitation `json:"student"`
}

type Data_DataStudent_CancelInvitation struct {
	Data_CancelInvitation_CancelInvitation Data_CancelInvitation_CancelInvitation `json:"cancelInvitation"`
	Typename         string           `json:"__typename"`
}

type Data_CancelInvitation_CancelInvitation struct {
	Student          Data_CancelInvitationStudent_CancelInvitation `json:"student"`
	InvitationStatus string                  `json:"invitationStatus"`
	Typename         string                  `json:"__typename"`
}

type Data_CancelInvitationStudent_CancelInvitation struct {
	ID       string `json:"id"`
	Data_User_CancelInvitation     Data_User_CancelInvitation   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_CancelInvitation struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_CancelInvitation Data_UserExperience_CancelInvitation `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_CancelInvitation struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Data_Level_CancelInvitation            Data_Level_CancelInvitation  `json:"level"`
	Typename         string `json:"__typename"`
}

type Data_Level_CancelInvitation struct {
	ID       int64  `json:"id"`
	Data_Range_CancelInvitation    Data_Range_CancelInvitation  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_CancelInvitation struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) CancelInvitation(variables Variables_CancelInvitation) (Data_CancelInvitation, error) {
	request := gql.NewQueryRequest[Variables_CancelInvitation](
		"mutation cancelInvitation($teamId: UUID!, $userId: ID!) {\n  student {\n    cancelInvitation(teamId: $teamId, userId: $userId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CancelInvitation](ctx, request)
}
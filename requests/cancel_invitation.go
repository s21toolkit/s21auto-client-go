package requests

import "github.com/s21toolkit/s21client/gql"

type CancelInvitation_Variables struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type CancelInvitation_Data struct {
	Student CancelInvitation_Data_DataStudent `json:"student"`
}

type CancelInvitation_Data_DataStudent struct {
	CancelInvitation CancelInvitation_Data_CancelInvitation `json:"cancelInvitation"`
	Typename         string           `json:"__typename"`
}

type CancelInvitation_Data_CancelInvitation struct {
	Student          CancelInvitation_Data_CancelInvitationStudent `json:"student"`
	InvitationStatus string                  `json:"invitationStatus"`
	Typename         string                  `json:"__typename"`
}

type CancelInvitation_Data_CancelInvitationStudent struct {
	ID       string `json:"id"`
	User     CancelInvitation_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type CancelInvitation_Data_User struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience CancelInvitation_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type CancelInvitation_Data_UserExperience struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level            CancelInvitation_Data_Level  `json:"level"`
	Typename         string `json:"__typename"`
}

type CancelInvitation_Data_Level struct {
	ID       int64  `json:"id"`
	Range    CancelInvitation_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type CancelInvitation_Data_Range struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) CancelInvitation(variables CancelInvitation_Variables) (CancelInvitation_Data, error) {
	request := gql.NewQueryRequest[CancelInvitation_Variables](
		"mutation cancelInvitation($teamId: UUID!, $userId: ID!) {\n  student {\n    cancelInvitation(teamId: $teamId, userId: $userId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CancelInvitation_Data](ctx, request)
}
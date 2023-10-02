package requests

import "github.com/s21toolkit/s21client/gql"

type GetAvailableStudentsForTeams_Variables struct {
	GoalID string `json:"goalId"`
}


type GetAvailableStudentsForTeams_Data struct {
	Student GetAvailableStudentsForTeams_Data_DataStudent `json:"student"`
}

type GetAvailableStudentsForTeams_Data_DataStudent struct {
	GetAvailableStudentsForTeam []GetAvailableStudentsForTeams_Data_GetAvailableStudentsForTeam `json:"getAvailableStudentsForTeam"`
	Typename                    string                        `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_GetAvailableStudentsForTeam struct {
	Student          GetAvailableStudentsForTeams_Data_GetAvailableStudentsForTeamStudent `json:"student"`
	InvitationStatus string                             `json:"invitationStatus"`
	Typename         string                             `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_GetAvailableStudentsForTeamStudent struct {
	ID       string `json:"id"`
	User     GetAvailableStudentsForTeams_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_User struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience GetAvailableStudentsForTeams_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_UserExperience struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level            GetAvailableStudentsForTeams_Data_Level  `json:"level"`
	Typename         string `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_Level struct {
	ID       int64  `json:"id"`
	Range    GetAvailableStudentsForTeams_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type GetAvailableStudentsForTeams_Data_Range struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetAvailableStudentsForTeams(variables GetAvailableStudentsForTeams_Variables) (GetAvailableStudentsForTeams_Data, error) {
	request := gql.NewQueryRequest[GetAvailableStudentsForTeams_Variables](
		"query getAvailableStudentsForTeams($goalId: ID!) {\n  student {\n    getAvailableStudentsForTeam(goalId: $goalId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetAvailableStudentsForTeams_Data](ctx, request)
}
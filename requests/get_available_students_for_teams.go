package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetAvailableStudentsForTeams struct {
	GoalID string `json:"goalId"`
}


type Data_GetAvailableStudentsForTeams struct {
	Student Data_DataStudent_GetAvailableStudentsForTeams `json:"student"`
}

type Data_DataStudent_GetAvailableStudentsForTeams struct {
	Data_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams []Data_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams `json:"getAvailableStudentsForTeam"`
	Typename                    string                        `json:"__typename"`
}

type Data_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams struct {
	Student          Data_GetAvailableStudentsForTeamStudent_GetAvailableStudentsForTeams `json:"student"`
	InvitationStatus string                             `json:"invitationStatus"`
	Typename         string                             `json:"__typename"`
}

type Data_GetAvailableStudentsForTeamStudent_GetAvailableStudentsForTeams struct {
	ID       string `json:"id"`
	Data_User_GetAvailableStudentsForTeams     Data_User_GetAvailableStudentsForTeams   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_GetAvailableStudentsForTeams struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_GetAvailableStudentsForTeams Data_UserExperience_GetAvailableStudentsForTeams `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_GetAvailableStudentsForTeams struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Data_Level_GetAvailableStudentsForTeams            Data_Level_GetAvailableStudentsForTeams  `json:"level"`
	Typename         string `json:"__typename"`
}

type Data_Level_GetAvailableStudentsForTeams struct {
	ID       int64  `json:"id"`
	Data_Range_GetAvailableStudentsForTeams    Data_Range_GetAvailableStudentsForTeams  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_GetAvailableStudentsForTeams struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetAvailableStudentsForTeams(variables Variables_GetAvailableStudentsForTeams) (Data_GetAvailableStudentsForTeams, error) {
	request := gql.NewQueryRequest[Variables_GetAvailableStudentsForTeams](
		"query getAvailableStudentsForTeams($goalId: ID!) {\n  student {\n    getAvailableStudentsForTeam(goalId: $goalId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetAvailableStudentsForTeams](ctx, request)
}
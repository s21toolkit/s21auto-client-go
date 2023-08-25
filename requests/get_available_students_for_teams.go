package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetAvailableStudentsForTeams struct {
	GoalID string `json:"goalId"`
}


type Response_Data_GetAvailableStudentsForTeams struct {
	Student Response_DataStudent_GetAvailableStudentsForTeams `json:"student"`
}

type Response_DataStudent_GetAvailableStudentsForTeams struct {
	Response_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams []Response_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams `json:"getAvailableStudentsForTeam"`
	Typename                    string                        `json:"__typename"`
}

type Response_GetAvailableStudentsForTeam_GetAvailableStudentsForTeams struct {
	Student          Response_GetAvailableStudentsForTeamStudent_GetAvailableStudentsForTeams `json:"student"`
	InvitationStatus string                             `json:"invitationStatus"`
	Typename         string                             `json:"__typename"`
}

type Response_GetAvailableStudentsForTeamStudent_GetAvailableStudentsForTeams struct {
	ID       string `json:"id"`
	Response_User_GetAvailableStudentsForTeams     Response_User_GetAvailableStudentsForTeams   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_GetAvailableStudentsForTeams struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Response_UserExperience_GetAvailableStudentsForTeams Response_UserExperience_GetAvailableStudentsForTeams `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Response_UserExperience_GetAvailableStudentsForTeams struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Response_Level_GetAvailableStudentsForTeams            Response_Level_GetAvailableStudentsForTeams  `json:"level"`
	Typename         string `json:"__typename"`
}

type Response_Level_GetAvailableStudentsForTeams struct {
	ID       int64  `json:"id"`
	Response_Range_GetAvailableStudentsForTeams    Response_Range_GetAvailableStudentsForTeams  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_GetAvailableStudentsForTeams struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetAvailableStudentsForTeams(variables Request_Variables_GetAvailableStudentsForTeams) (Response_Data_GetAvailableStudentsForTeams, error) {
	request := gql.NewQueryRequest[Request_Variables_GetAvailableStudentsForTeams](
		"query getAvailableStudentsForTeams($goalId: ID!) {\n  student {\n    getAvailableStudentsForTeam(goalId: $goalId) {\n      ...StudentInvitationInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentInvitationInfo on StudentInvitationInfo {\n  student {\n    ...AvailableStudentForTeam\n    __typename\n  }\n  invitationStatus\n  __typename\n}\n\nfragment AvailableStudentForTeam on Student {\n  id\n  user {\n    id\n    login\n    avatarUrl\n    userExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetAvailableStudentsForTeams](ctx, request)
}
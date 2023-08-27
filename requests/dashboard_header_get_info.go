package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_DashboardHeaderGetInfo struct {
}


type Data_DashboardHeaderGetInfo struct {
	Data_User_DashboardHeaderGetInfo    Data_User_DashboardHeaderGetInfo    `json:"user"`
	Data_Student_DashboardHeaderGetInfo Data_Student_DashboardHeaderGetInfo `json:"student"`
}

type Data_Student_DashboardHeaderGetInfo struct {
	Data_GetUserTournamentWidget_DashboardHeaderGetInfo Data_GetUserTournamentWidget_DashboardHeaderGetInfo `json:"getUserTournamentWidget"`
	Data_GetExperience_DashboardHeaderGetInfo           Data_GetExperience_DashboardHeaderGetInfo           `json:"getExperience"`
	Typename                string                  `json:"__typename"`
}

type Data_GetExperience_DashboardHeaderGetInfo struct {
	ID               string `json:"id"`
	Value            int64  `json:"value"`
	Data_Level_DashboardHeaderGetInfo            Data_Level_DashboardHeaderGetInfo  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Data_Level_DashboardHeaderGetInfo struct {
	ID       int64  `json:"id"`
	Data_Range_DashboardHeaderGetInfo    Data_Range_DashboardHeaderGetInfo  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_DashboardHeaderGetInfo struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type Data_GetUserTournamentWidget_DashboardHeaderGetInfo struct {
	Data_CoalitionMember_DashboardHeaderGetInfo      Data_CoalitionMember_DashboardHeaderGetInfo      `json:"coalitionMember"`
	Data_LastTournamentResult_DashboardHeaderGetInfo Data_LastTournamentResult_DashboardHeaderGetInfo `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Data_CoalitionMember_DashboardHeaderGetInfo struct {
	Data_Coalition_DashboardHeaderGetInfo                  Data_Coalition_DashboardHeaderGetInfo                  `json:"coalition"`
	Data_CurrentTournamentPowerRank_DashboardHeaderGetInfo Data_CurrentTournamentPowerRank_DashboardHeaderGetInfo `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Data_Coalition_DashboardHeaderGetInfo struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type Data_CurrentTournamentPowerRank_DashboardHeaderGetInfo struct {
	Rank     int64  `json:"rank"`
	Typename string `json:"__typename"`
}

type Data_LastTournamentResult_DashboardHeaderGetInfo struct {
	UserRank int64  `json:"userRank"`
	Typename string `json:"__typename"`
}

type Data_User_DashboardHeaderGetInfo struct {
	Data_GetCurrentUser_DashboardHeaderGetInfo Data_GetCurrentUser_DashboardHeaderGetInfo `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type Data_GetCurrentUser_DashboardHeaderGetInfo struct {
	ID                     string        `json:"id"`
	Login                  string        `json:"login"`
	AvatarURL              string        `json:"avatarUrl"`
	FirstName              string        `json:"firstName"`
	MiddleName             string        `json:"middleName"`
	LastName               string        `json:"lastName"`
	CurrentSchoolStudentID string        `json:"currentSchoolStudentId"`
	StudentRoles           []Data_StudentRole_DashboardHeaderGetInfo `json:"studentRoles"`
	Typename               string        `json:"__typename"`
}

type Data_StudentRole_DashboardHeaderGetInfo struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Data_School_DashboardHeaderGetInfo   Data_School_DashboardHeaderGetInfo `json:"school"`
	Typename string `json:"__typename"`
}

type Data_School_DashboardHeaderGetInfo struct {
	ID        string `json:"id"`
	ShortName string `json:"shortName"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) DashboardHeaderGetInfo(variables Variables_DashboardHeaderGetInfo) (Data_DashboardHeaderGetInfo, error) {
	request := gql.NewQueryRequest[Variables_DashboardHeaderGetInfo](
		"query dashboardHeaderGetInfo {\n  user {\n    getCurrentUser {\n      id\n      login\n      avatarUrl\n      firstName\n      middleName\n      lastName\n      currentSchoolStudentId\n      studentRoles {\n        id\n        status\n        school {\n          id\n          shortName\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  student {\n    getUserTournamentWidget {\n      coalitionMember {\n        coalition {\n          avatarUrl\n          color\n          name\n          memberCount\n          __typename\n        }\n        currentTournamentPowerRank {\n          rank\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        userRank\n        __typename\n      }\n      __typename\n    }\n    getExperience {\n      id\n      value\n      level {\n        id\n        range {\n          id\n          levelCode\n          rightBorder\n          leftBorder\n          __typename\n        }\n        __typename\n      }\n      cookiesCount\n      coinsCount\n      codeReviewPoints\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_DashboardHeaderGetInfo](ctx, request)
}
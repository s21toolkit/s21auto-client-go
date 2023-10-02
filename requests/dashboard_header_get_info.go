package requests

import "github.com/s21toolkit/s21client/gql"

type DashboardHeaderGetInfo_Variables struct {
}


type DashboardHeaderGetInfo_Data struct {
	User    DashboardHeaderGetInfo_Data_User    `json:"user"`
	Student DashboardHeaderGetInfo_Data_Student `json:"student"`
}

type DashboardHeaderGetInfo_Data_Student struct {
	GetUserTournamentWidget DashboardHeaderGetInfo_Data_GetUserTournamentWidget `json:"getUserTournamentWidget"`
	GetExperience           DashboardHeaderGetInfo_Data_GetExperience           `json:"getExperience"`
	Typename                string                  `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_GetExperience struct {
	ID               string `json:"id"`
	Value            int64  `json:"value"`
	Level            DashboardHeaderGetInfo_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_Level struct {
	ID       int64  `json:"id"`
	Range    DashboardHeaderGetInfo_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_Range struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_GetUserTournamentWidget struct {
	CoalitionMember      DashboardHeaderGetInfo_Data_CoalitionMember      `json:"coalitionMember"`
	LastTournamentResult DashboardHeaderGetInfo_Data_LastTournamentResult `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_CoalitionMember struct {
	Coalition                  DashboardHeaderGetInfo_Data_Coalition                  `json:"coalition"`
	CurrentTournamentPowerRank DashboardHeaderGetInfo_Data_CurrentTournamentPowerRank `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_Coalition struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_CurrentTournamentPowerRank struct {
	Rank     int64  `json:"rank"`
	Typename string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_LastTournamentResult struct {
	UserRank int64  `json:"userRank"`
	Typename string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_User struct {
	GetCurrentUser DashboardHeaderGetInfo_Data_GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_GetCurrentUser struct {
	ID                     string        `json:"id"`
	Login                  string        `json:"login"`
	AvatarURL              string        `json:"avatarUrl"`
	FirstName              string        `json:"firstName"`
	MiddleName             string        `json:"middleName"`
	LastName               string        `json:"lastName"`
	CurrentSchoolStudentID string        `json:"currentSchoolStudentId"`
	StudentRoles           []DashboardHeaderGetInfo_Data_StudentRole `json:"studentRoles"`
	Typename               string        `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_StudentRole struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	School   DashboardHeaderGetInfo_Data_School `json:"school"`
	Typename string `json:"__typename"`
}

type DashboardHeaderGetInfo_Data_School struct {
	ID        string `json:"id"`
	ShortName string `json:"shortName"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) DashboardHeaderGetInfo(variables DashboardHeaderGetInfo_Variables) (DashboardHeaderGetInfo_Data, error) {
	request := gql.NewQueryRequest[DashboardHeaderGetInfo_Variables](
		"query dashboardHeaderGetInfo {\n  user {\n    getCurrentUser {\n      id\n      login\n      avatarUrl\n      firstName\n      middleName\n      lastName\n      currentSchoolStudentId\n      studentRoles {\n        id\n        status\n        school {\n          id\n          shortName\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  student {\n    getUserTournamentWidget {\n      coalitionMember {\n        coalition {\n          avatarUrl\n          color\n          name\n          memberCount\n          __typename\n        }\n        currentTournamentPowerRank {\n          rank\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        userRank\n        __typename\n      }\n      __typename\n    }\n    getExperience {\n      id\n      value\n      level {\n        id\n        range {\n          id\n          levelCode\n          rightBorder\n          leftBorder\n          __typename\n        }\n        __typename\n      }\n      cookiesCount\n      coinsCount\n      codeReviewPoints\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[DashboardHeaderGetInfo_Data](ctx, request)
}
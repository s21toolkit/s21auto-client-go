package requests

import "s21client/gql"

type Request_Variables_DashboardHeaderGetInfo struct {
}


type Response_Data_DashboardHeaderGetInfo struct {
	Response_User_DashboardHeaderGetInfo    Response_User_DashboardHeaderGetInfo    `json:"user"`
	Response_Student_DashboardHeaderGetInfo Response_Student_DashboardHeaderGetInfo `json:"student"`
}

type Response_Student_DashboardHeaderGetInfo struct {
	Response_GetUserTournamentWidget_DashboardHeaderGetInfo Response_GetUserTournamentWidget_DashboardHeaderGetInfo `json:"getUserTournamentWidget"`
	Response_GetExperience_DashboardHeaderGetInfo           Response_GetExperience_DashboardHeaderGetInfo           `json:"getExperience"`
	Typename                string                  `json:"__typename"`
}

type Response_GetExperience_DashboardHeaderGetInfo struct {
	ID               string `json:"id"`
	Value            int64  `json:"value"`
	Response_Level_DashboardHeaderGetInfo            Response_Level_DashboardHeaderGetInfo  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Response_Level_DashboardHeaderGetInfo struct {
	ID       int64  `json:"id"`
	Response_Range_DashboardHeaderGetInfo    Response_Range_DashboardHeaderGetInfo  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_DashboardHeaderGetInfo struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type Response_GetUserTournamentWidget_DashboardHeaderGetInfo struct {
	Response_CoalitionMember_DashboardHeaderGetInfo      Response_CoalitionMember_DashboardHeaderGetInfo      `json:"coalitionMember"`
	Response_LastTournamentResult_DashboardHeaderGetInfo Response_LastTournamentResult_DashboardHeaderGetInfo `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Response_CoalitionMember_DashboardHeaderGetInfo struct {
	Response_Coalition_DashboardHeaderGetInfo                  Response_Coalition_DashboardHeaderGetInfo                  `json:"coalition"`
	Response_CurrentTournamentPowerRank_DashboardHeaderGetInfo Response_CurrentTournamentPowerRank_DashboardHeaderGetInfo `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Response_Coalition_DashboardHeaderGetInfo struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type Response_CurrentTournamentPowerRank_DashboardHeaderGetInfo struct {
	Rank     int64  `json:"rank"`
	Typename string `json:"__typename"`
}

type Response_LastTournamentResult_DashboardHeaderGetInfo struct {
	UserRank int64  `json:"userRank"`
	Typename string `json:"__typename"`
}

type Response_User_DashboardHeaderGetInfo struct {
	Response_GetCurrentUser_DashboardHeaderGetInfo Response_GetCurrentUser_DashboardHeaderGetInfo `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type Response_GetCurrentUser_DashboardHeaderGetInfo struct {
	ID                     string        `json:"id"`
	Login                  string        `json:"login"`
	AvatarURL              string        `json:"avatarUrl"`
	FirstName              string        `json:"firstName"`
	MiddleName             string        `json:"middleName"`
	LastName               string        `json:"lastName"`
	CurrentSchoolStudentID string        `json:"currentSchoolStudentId"`
	StudentRoles           []Response_StudentRole_DashboardHeaderGetInfo `json:"studentRoles"`
	Typename               string        `json:"__typename"`
}

type Response_StudentRole_DashboardHeaderGetInfo struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Response_School_DashboardHeaderGetInfo   Response_School_DashboardHeaderGetInfo `json:"school"`
	Typename string `json:"__typename"`
}

type Response_School_DashboardHeaderGetInfo struct {
	ID        string `json:"id"`
	ShortName string `json:"shortName"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) DashboardHeaderGetInfo(variables Request_Variables_DashboardHeaderGetInfo) (Response_Data_DashboardHeaderGetInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_DashboardHeaderGetInfo](
		"query dashboardHeaderGetInfo {   user {     getCurrentUser {       id       login       avatarUrl       firstName       middleName       lastName       currentSchoolStudentId       studentRoles {         id         status         school {           id           shortName           __typename         }         __typename       }       __typename     }     __typename   }   student {     getUserTournamentWidget {       coalitionMember {         coalition {           avatarUrl           color           name           memberCount           __typename         }         currentTournamentPowerRank {           rank           __typename         }         __typename       }       lastTournamentResult {         userRank         __typename       }       __typename     }     getExperience {       id       value       level {         id         range {           id           levelCode           rightBorder           leftBorder           __typename         }         __typename       }       cookiesCount       coinsCount       codeReviewPoints       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_DashboardHeaderGetInfo](ctx, request)
}
package requests

import "s21client/gql"

type Variables_DashboardHeaderGetInfo struct {
}


type Data_DashboardHeaderGetInfo struct {
	User_DashboardHeaderGetInfo    User_DashboardHeaderGetInfo    `json:"user"`
	Student_DashboardHeaderGetInfo Student_DashboardHeaderGetInfo `json:"student"`
}

type Student_DashboardHeaderGetInfo struct {
	GetUserTournamentWidget_DashboardHeaderGetInfo GetUserTournamentWidget_DashboardHeaderGetInfo `json:"getUserTournamentWidget"`
	GetExperience_DashboardHeaderGetInfo           GetExperience_DashboardHeaderGetInfo           `json:"getExperience"`
	Typename                string                  `json:"__typename"`
}

type GetExperience_DashboardHeaderGetInfo struct {
	ID               string `json:"id"`
	Value            int64  `json:"value"`
	Level_DashboardHeaderGetInfo            Level_DashboardHeaderGetInfo  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Level_DashboardHeaderGetInfo struct {
	ID       int64  `json:"id"`
	Range_DashboardHeaderGetInfo    Range_DashboardHeaderGetInfo  `json:"range"`
	Typename string `json:"__typename"`
}

type Range_DashboardHeaderGetInfo struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type GetUserTournamentWidget_DashboardHeaderGetInfo struct {
	CoalitionMember_DashboardHeaderGetInfo      CoalitionMember_DashboardHeaderGetInfo      `json:"coalitionMember"`
	LastTournamentResult_DashboardHeaderGetInfo LastTournamentResult_DashboardHeaderGetInfo `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type CoalitionMember_DashboardHeaderGetInfo struct {
	Coalition_DashboardHeaderGetInfo                  Coalition_DashboardHeaderGetInfo                  `json:"coalition"`
	CurrentTournamentPowerRank_DashboardHeaderGetInfo CurrentTournamentPowerRank_DashboardHeaderGetInfo `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Coalition_DashboardHeaderGetInfo struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type CurrentTournamentPowerRank_DashboardHeaderGetInfo struct {
	Rank     int64  `json:"rank"`
	Typename string `json:"__typename"`
}

type LastTournamentResult_DashboardHeaderGetInfo struct {
	UserRank int64  `json:"userRank"`
	Typename string `json:"__typename"`
}

type User_DashboardHeaderGetInfo struct {
	GetCurrentUser_DashboardHeaderGetInfo GetCurrentUser_DashboardHeaderGetInfo `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCurrentUser_DashboardHeaderGetInfo struct {
	ID                     string        `json:"id"`
	Login                  string        `json:"login"`
	AvatarURL              string        `json:"avatarUrl"`
	FirstName              string        `json:"firstName"`
	MiddleName             string        `json:"middleName"`
	LastName               string        `json:"lastName"`
	CurrentSchoolStudentID string        `json:"currentSchoolStudentId"`
	StudentRoles           []StudentRole_DashboardHeaderGetInfo `json:"studentRoles"`
	Typename               string        `json:"__typename"`
}

type StudentRole_DashboardHeaderGetInfo struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	School_DashboardHeaderGetInfo   School_DashboardHeaderGetInfo `json:"school"`
	Typename string `json:"__typename"`
}

type School_DashboardHeaderGetInfo struct {
	ID        string `json:"id"`
	ShortName string `json:"shortName"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) DashboardHeaderGetInfo(variables Variables_DashboardHeaderGetInfo) (Data_DashboardHeaderGetInfo, error) {
	request := gql.NewQueryRequest[Variables_DashboardHeaderGetInfo](
		"query dashboardHeaderGetInfo {   user {     getCurrentUser {       id       login       avatarUrl       firstName       middleName       lastName       currentSchoolStudentId       studentRoles {         id         status         school {           id           shortName           __typename         }         __typename       }       __typename     }     __typename   }   student {     getUserTournamentWidget {       coalitionMember {         coalition {           avatarUrl           color           name           memberCount           __typename         }         currentTournamentPowerRank {           rank           __typename         }         __typename       }       lastTournamentResult {         userRank         __typename       }       __typename     }     getExperience {       id       value       level {         id         range {           id           levelCode           rightBorder           leftBorder           __typename         }         __typename       }       cookiesCount       coinsCount       codeReviewPoints       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_DashboardHeaderGetInfo](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileGetCoalition struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetCoalition struct {
	Data_Student_PublicProfileGetCoalition Data_Student_PublicProfileGetCoalition `json:"student"`
}

type Data_Student_PublicProfileGetCoalition struct {
	Data_GetUserTournamentWidget_PublicProfileGetCoalition Data_GetUserTournamentWidget_PublicProfileGetCoalition `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Data_GetUserTournamentWidget_PublicProfileGetCoalition struct {
	Data_CoalitionMember_PublicProfileGetCoalition      Data_CoalitionMember_PublicProfileGetCoalition      `json:"coalitionMember"`
	Data_LastTournamentResult_PublicProfileGetCoalition Data_LastTournamentResult_PublicProfileGetCoalition `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Data_CoalitionMember_PublicProfileGetCoalition struct {
	Data_Coalition_PublicProfileGetCoalition                  Data_Coalition_PublicProfileGetCoalition                  `json:"coalition"`
	Data_CurrentTournamentPowerRank_PublicProfileGetCoalition Data_CurrentTournamentPowerRank_PublicProfileGetCoalition `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Data_Coalition_PublicProfileGetCoalition struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type Data_CurrentTournamentPowerRank_PublicProfileGetCoalition struct {
	Rank     int64  `json:"rank"`
	Data_Power_PublicProfileGetCoalition    Data_Power_PublicProfileGetCoalition  `json:"power"`
	Typename string `json:"__typename"`
}

type Data_Power_PublicProfileGetCoalition struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Typename string `json:"__typename"`
}

type Data_LastTournamentResult_PublicProfileGetCoalition struct {
	UserRank int64  `json:"userRank"`
	Data_Power_PublicProfileGetCoalition    int64  `json:"power"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetCoalition(variables Variables_PublicProfileGetCoalition) (Data_PublicProfileGetCoalition, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetCoalition](
		"query publicProfileGetCoalition($userId: UUID!) {\n  student {\n    getUserTournamentWidget(userId: $userId) {\n      coalitionMember {\n        coalition {\n          avatarUrl\n          color\n          name\n          memberCount\n          __typename\n        }\n        currentTournamentPowerRank {\n          rank\n          power {\n            id\n            points\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        userRank\n        power\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetCoalition](ctx, request)
}
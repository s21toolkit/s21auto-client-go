package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetCoalition_Variables struct {
	UserID string `json:"userId"`
}


type PublicProfileGetCoalition_Data struct {
	Student PublicProfileGetCoalition_Data_Student `json:"student"`
}

type PublicProfileGetCoalition_Data_Student struct {
	GetUserTournamentWidget PublicProfileGetCoalition_Data_GetUserTournamentWidget `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type PublicProfileGetCoalition_Data_GetUserTournamentWidget struct {
	CoalitionMember      PublicProfileGetCoalition_Data_CoalitionMember      `json:"coalitionMember"`
	LastTournamentResult PublicProfileGetCoalition_Data_LastTournamentResult `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type PublicProfileGetCoalition_Data_CoalitionMember struct {
	Coalition                  PublicProfileGetCoalition_Data_Coalition                  `json:"coalition"`
	CurrentTournamentPowerRank PublicProfileGetCoalition_Data_CurrentTournamentPowerRank `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type PublicProfileGetCoalition_Data_Coalition struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type PublicProfileGetCoalition_Data_CurrentTournamentPowerRank struct {
	Rank     int64  `json:"rank"`
	Power    PublicProfileGetCoalition_Data_Power  `json:"power"`
	Typename string `json:"__typename"`
}

type PublicProfileGetCoalition_Data_Power struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Typename string `json:"__typename"`
}

type PublicProfileGetCoalition_Data_LastTournamentResult struct {
	UserRank int64  `json:"userRank"`
	Power    int64  `json:"power"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetCoalition(variables PublicProfileGetCoalition_Variables) (PublicProfileGetCoalition_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetCoalition_Variables](
		"query publicProfileGetCoalition($userId: UUID!) {\n  student {\n    getUserTournamentWidget(userId: $userId) {\n      coalitionMember {\n        coalition {\n          avatarUrl\n          color\n          name\n          memberCount\n          __typename\n        }\n        currentTournamentPowerRank {\n          rank\n          power {\n            id\n            points\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        userRank\n        power\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetCoalition_Data](ctx, request)
}
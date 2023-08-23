package requests

import "s21client/gql"

type Variables_PublicProfileGetCoalition struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetCoalition struct {
	Student_PublicProfileGetCoalition Student_PublicProfileGetCoalition `json:"student"`
}

type Student_PublicProfileGetCoalition struct {
	GetUserTournamentWidget_PublicProfileGetCoalition GetUserTournamentWidget_PublicProfileGetCoalition `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type GetUserTournamentWidget_PublicProfileGetCoalition struct {
	CoalitionMember_PublicProfileGetCoalition      CoalitionMember_PublicProfileGetCoalition      `json:"coalitionMember"`
	LastTournamentResult_PublicProfileGetCoalition LastTournamentResult_PublicProfileGetCoalition `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type CoalitionMember_PublicProfileGetCoalition struct {
	Coalition_PublicProfileGetCoalition                  Coalition_PublicProfileGetCoalition                  `json:"coalition"`
	CurrentTournamentPowerRank_PublicProfileGetCoalition CurrentTournamentPowerRank_PublicProfileGetCoalition `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Coalition_PublicProfileGetCoalition struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type CurrentTournamentPowerRank_PublicProfileGetCoalition struct {
	Rank     int64  `json:"rank"`
	Power_PublicProfileGetCoalition    Power_PublicProfileGetCoalition  `json:"power"`
	Typename string `json:"__typename"`
}

type Power_PublicProfileGetCoalition struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Typename string `json:"__typename"`
}

type LastTournamentResult_PublicProfileGetCoalition struct {
	UserRank int64  `json:"userRank"`
	Power_PublicProfileGetCoalition    int64  `json:"power"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetCoalition(variables Variables_PublicProfileGetCoalition) (Data_PublicProfileGetCoalition, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetCoalition](
		"query publicProfileGetCoalition($userId: UUID!) {   student {     getUserTournamentWidget(userId: $userId) {       coalitionMember {         coalition {           avatarUrl           color           name           memberCount           __typename         }         currentTournamentPowerRank {           rank           power {             id             points             __typename           }           __typename         }         __typename       }       lastTournamentResult {         userRank         power         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetCoalition](ctx, request)
}
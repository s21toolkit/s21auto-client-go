package requests

import "s21client/gql"

type Request_Variables_PublicProfileGetCoalition struct {
	UserID string `json:"userId"`
}


type Response_Data_PublicProfileGetCoalition struct {
	Response_Student_PublicProfileGetCoalition Response_Student_PublicProfileGetCoalition `json:"student"`
}

type Response_Student_PublicProfileGetCoalition struct {
	Response_GetUserTournamentWidget_PublicProfileGetCoalition Response_GetUserTournamentWidget_PublicProfileGetCoalition `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Response_GetUserTournamentWidget_PublicProfileGetCoalition struct {
	Response_CoalitionMember_PublicProfileGetCoalition      Response_CoalitionMember_PublicProfileGetCoalition      `json:"coalitionMember"`
	Response_LastTournamentResult_PublicProfileGetCoalition Response_LastTournamentResult_PublicProfileGetCoalition `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Response_CoalitionMember_PublicProfileGetCoalition struct {
	Response_Coalition_PublicProfileGetCoalition                  Response_Coalition_PublicProfileGetCoalition                  `json:"coalition"`
	Response_CurrentTournamentPowerRank_PublicProfileGetCoalition Response_CurrentTournamentPowerRank_PublicProfileGetCoalition `json:"currentTournamentPowerRank"`
	Typename                   string                     `json:"__typename"`
}

type Response_Coalition_PublicProfileGetCoalition struct {
	AvatarURL   string `json:"avatarUrl"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
	Typename    string `json:"__typename"`
}

type Response_CurrentTournamentPowerRank_PublicProfileGetCoalition struct {
	Rank     int64  `json:"rank"`
	Response_Power_PublicProfileGetCoalition    Response_Power_PublicProfileGetCoalition  `json:"power"`
	Typename string `json:"__typename"`
}

type Response_Power_PublicProfileGetCoalition struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Typename string `json:"__typename"`
}

type Response_LastTournamentResult_PublicProfileGetCoalition struct {
	UserRank int64  `json:"userRank"`
	Response_Power_PublicProfileGetCoalition    int64  `json:"power"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetCoalition(variables Request_Variables_PublicProfileGetCoalition) (Response_Data_PublicProfileGetCoalition, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetCoalition](
		"query publicProfileGetCoalition($userId: UUID!) {   student {     getUserTournamentWidget(userId: $userId) {       coalitionMember {         coalition {           avatarUrl           color           name           memberCount           __typename         }         currentTournamentPowerRank {           rank           power {             id             points             __typename           }           __typename         }         __typename       }       lastTournamentResult {         userRank         power         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetCoalition](ctx, request)
}
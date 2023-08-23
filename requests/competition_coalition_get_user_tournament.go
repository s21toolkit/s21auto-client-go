package requests

import "s21client/gql"

type Variables_CompetitionCoalitionGetUserTournament struct {
}


type Data_CompetitionCoalitionGetUserTournament struct {
	Student_CompetitionCoalitionGetUserTournament Student_CompetitionCoalitionGetUserTournament `json:"student"`
}

type Student_CompetitionCoalitionGetUserTournament struct {
	GetUserTournamentWidget_CompetitionCoalitionGetUserTournament GetUserTournamentWidget_CompetitionCoalitionGetUserTournament `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type GetUserTournamentWidget_CompetitionCoalitionGetUserTournament struct {
	CoalitionMember_CompetitionCoalitionGetUserTournament      CoalitionMember_CompetitionCoalitionGetUserTournament      `json:"coalitionMember"`
	LastTournamentResult_CompetitionCoalitionGetUserTournament LastTournamentResult_CompetitionCoalitionGetUserTournament `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type CoalitionMember_CompetitionCoalitionGetUserTournament struct {
	Coalition_CompetitionCoalitionGetUserTournament Coalition_CompetitionCoalitionGetUserTournament `json:"coalition"`
	Typename  string    `json:"__typename"`
}

type Coalition_CompetitionCoalitionGetUserTournament struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	AvatarURL         string            `json:"avatarUrl"`
	BackgroundURL     string            `json:"backgroundUrl"`
	BackgroundURLBig  string            `json:"backgroundUrlBig"`
	MemberCount       int64             `json:"memberCount"`
	Color             string            `json:"color"`
	CurrentTournament_CompetitionCoalitionGetUserTournament CurrentTournament_CompetitionCoalitionGetUserTournament `json:"currentTournament"`
	MasterUser_CompetitionCoalitionGetUserTournament        MasterUser_CompetitionCoalitionGetUserTournament        `json:"masterUser"`
	Typename          string            `json:"__typename"`
}

type CurrentTournament_CompetitionCoalitionGetUserTournament struct {
	Points     int64      `json:"points"`
	Tournament_CompetitionCoalitionGetUserTournament Tournament_CompetitionCoalitionGetUserTournament `json:"tournament"`
	Typename   string     `json:"__typename"`
}

type Tournament_CompetitionCoalitionGetUserTournament struct {
	Name      string `json:"name"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Typename  string `json:"__typename"`
}

type MasterUser_CompetitionCoalitionGetUserTournament struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

type LastTournamentResult_CompetitionCoalitionGetUserTournament struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CompetitionCoalitionGetUserTournament(variables Variables_CompetitionCoalitionGetUserTournament) (Data_CompetitionCoalitionGetUserTournament, error) {
	request := gql.NewQueryRequest[Variables_CompetitionCoalitionGetUserTournament](
		"query competitionCoalitionGetUserTournament {   student {     getUserTournamentWidget {       coalitionMember {         coalition {           ...CompetitionCurrentCoalition           __typename         }         __typename       }       lastTournamentResult {         id         __typename       }       __typename     }     __typename   } }  fragment CompetitionCurrentCoalition on GameCoalition {   id   name   avatarUrl   backgroundUrl   backgroundUrlBig   memberCount   color   currentTournament {     points     tournament {       name       timeStart       timeEnd       __typename     }     __typename   }   masterUser {     login     avatarUrl     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_CompetitionCoalitionGetUserTournament](ctx, request)
}
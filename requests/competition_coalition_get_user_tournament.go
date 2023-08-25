package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CompetitionCoalitionGetUserTournament struct {
}


type Response_Data_CompetitionCoalitionGetUserTournament struct {
	Response_Student_CompetitionCoalitionGetUserTournament Response_Student_CompetitionCoalitionGetUserTournament `json:"student"`
}

type Response_Student_CompetitionCoalitionGetUserTournament struct {
	Response_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament Response_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Response_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament struct {
	Response_CoalitionMember_CompetitionCoalitionGetUserTournament      Response_CoalitionMember_CompetitionCoalitionGetUserTournament      `json:"coalitionMember"`
	Response_LastTournamentResult_CompetitionCoalitionGetUserTournament Response_LastTournamentResult_CompetitionCoalitionGetUserTournament `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Response_CoalitionMember_CompetitionCoalitionGetUserTournament struct {
	Response_Coalition_CompetitionCoalitionGetUserTournament Response_Coalition_CompetitionCoalitionGetUserTournament `json:"coalition"`
	Typename  string    `json:"__typename"`
}

type Response_Coalition_CompetitionCoalitionGetUserTournament struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	AvatarURL         string            `json:"avatarUrl"`
	BackgroundURL     string            `json:"backgroundUrl"`
	BackgroundURLBig  string            `json:"backgroundUrlBig"`
	MemberCount       int64             `json:"memberCount"`
	Color             string            `json:"color"`
	Response_CurrentTournament_CompetitionCoalitionGetUserTournament Response_CurrentTournament_CompetitionCoalitionGetUserTournament `json:"currentTournament"`
	Response_MasterUser_CompetitionCoalitionGetUserTournament        Response_MasterUser_CompetitionCoalitionGetUserTournament        `json:"masterUser"`
	Typename          string            `json:"__typename"`
}

type Response_CurrentTournament_CompetitionCoalitionGetUserTournament struct {
	Points     int64      `json:"points"`
	Response_Tournament_CompetitionCoalitionGetUserTournament Response_Tournament_CompetitionCoalitionGetUserTournament `json:"tournament"`
	Typename   string     `json:"__typename"`
}

type Response_Tournament_CompetitionCoalitionGetUserTournament struct {
	Name      string `json:"name"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Typename  string `json:"__typename"`
}

type Response_MasterUser_CompetitionCoalitionGetUserTournament struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

type Response_LastTournamentResult_CompetitionCoalitionGetUserTournament struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CompetitionCoalitionGetUserTournament(variables Request_Variables_CompetitionCoalitionGetUserTournament) (Response_Data_CompetitionCoalitionGetUserTournament, error) {
	request := gql.NewQueryRequest[Request_Variables_CompetitionCoalitionGetUserTournament](
		"query competitionCoalitionGetUserTournament {   student {     getUserTournamentWidget {       coalitionMember {         coalition {           ...CompetitionCurrentCoalition           __typename         }         __typename       }       lastTournamentResult {         id         __typename       }       __typename     }     __typename   } }  fragment CompetitionCurrentCoalition on GameCoalition {   id   name   avatarUrl   backgroundUrl   backgroundUrlBig   memberCount   color   currentTournament {     points     tournament {       name       timeStart       timeEnd       __typename     }     __typename   }   masterUser {     login     avatarUrl     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_CompetitionCoalitionGetUserTournament](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CompetitionCoalitionGetUserTournament struct {
}


type Data_CompetitionCoalitionGetUserTournament struct {
	Data_Student_CompetitionCoalitionGetUserTournament Data_Student_CompetitionCoalitionGetUserTournament `json:"student"`
}

type Data_Student_CompetitionCoalitionGetUserTournament struct {
	Data_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament Data_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Data_GetUserTournamentWidget_CompetitionCoalitionGetUserTournament struct {
	Data_CoalitionMember_CompetitionCoalitionGetUserTournament      Data_CoalitionMember_CompetitionCoalitionGetUserTournament      `json:"coalitionMember"`
	Data_LastTournamentResult_CompetitionCoalitionGetUserTournament Data_LastTournamentResult_CompetitionCoalitionGetUserTournament `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type Data_CoalitionMember_CompetitionCoalitionGetUserTournament struct {
	Data_Coalition_CompetitionCoalitionGetUserTournament Data_Coalition_CompetitionCoalitionGetUserTournament `json:"coalition"`
	Typename  string    `json:"__typename"`
}

type Data_Coalition_CompetitionCoalitionGetUserTournament struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	AvatarURL         string            `json:"avatarUrl"`
	BackgroundURL     string            `json:"backgroundUrl"`
	BackgroundURLBig  string            `json:"backgroundUrlBig"`
	MemberCount       int64             `json:"memberCount"`
	Color             string            `json:"color"`
	Data_CurrentTournament_CompetitionCoalitionGetUserTournament Data_CurrentTournament_CompetitionCoalitionGetUserTournament `json:"currentTournament"`
	Data_MasterUser_CompetitionCoalitionGetUserTournament        Data_MasterUser_CompetitionCoalitionGetUserTournament        `json:"masterUser"`
	Typename          string            `json:"__typename"`
}

type Data_CurrentTournament_CompetitionCoalitionGetUserTournament struct {
	Points     int64      `json:"points"`
	Data_Tournament_CompetitionCoalitionGetUserTournament Data_Tournament_CompetitionCoalitionGetUserTournament `json:"tournament"`
	Typename   string     `json:"__typename"`
}

type Data_Tournament_CompetitionCoalitionGetUserTournament struct {
	Name      string `json:"name"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Typename  string `json:"__typename"`
}

type Data_MasterUser_CompetitionCoalitionGetUserTournament struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

type Data_LastTournamentResult_CompetitionCoalitionGetUserTournament struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CompetitionCoalitionGetUserTournament(variables Variables_CompetitionCoalitionGetUserTournament) (Data_CompetitionCoalitionGetUserTournament, error) {
	request := gql.NewQueryRequest[Variables_CompetitionCoalitionGetUserTournament](
		"query competitionCoalitionGetUserTournament {\n  student {\n    getUserTournamentWidget {\n      coalitionMember {\n        coalition {\n          ...CompetitionCurrentCoalition\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        id\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CompetitionCurrentCoalition on GameCoalition {\n  id\n  name\n  avatarUrl\n  backgroundUrl\n  backgroundUrlBig\n  memberCount\n  color\n  currentTournament {\n    points\n    tournament {\n      name\n      timeStart\n      timeEnd\n      __typename\n    }\n    __typename\n  }\n  masterUser {\n    login\n    avatarUrl\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CompetitionCoalitionGetUserTournament](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type CompetitionCoalitionGetUserTournament_Variables struct {
}


type CompetitionCoalitionGetUserTournament_Data struct {
	Student CompetitionCoalitionGetUserTournament_Data_Student `json:"student"`
}

type CompetitionCoalitionGetUserTournament_Data_Student struct {
	GetUserTournamentWidget CompetitionCoalitionGetUserTournament_Data_GetUserTournamentWidget `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_GetUserTournamentWidget struct {
	CoalitionMember      CompetitionCoalitionGetUserTournament_Data_CoalitionMember      `json:"coalitionMember"`
	LastTournamentResult CompetitionCoalitionGetUserTournament_Data_LastTournamentResult `json:"lastTournamentResult"`
	Typename             string               `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_CoalitionMember struct {
	Coalition CompetitionCoalitionGetUserTournament_Data_Coalition `json:"coalition"`
	Typename  string    `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_Coalition struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	AvatarURL         string            `json:"avatarUrl"`
	BackgroundURL     string            `json:"backgroundUrl"`
	BackgroundURLBig  string            `json:"backgroundUrlBig"`
	MemberCount       int64             `json:"memberCount"`
	Color             string            `json:"color"`
	CurrentTournament CompetitionCoalitionGetUserTournament_Data_CurrentTournament `json:"currentTournament"`
	MasterUser        CompetitionCoalitionGetUserTournament_Data_MasterUser        `json:"masterUser"`
	Typename          string            `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_CurrentTournament struct {
	Points     int64      `json:"points"`
	Tournament CompetitionCoalitionGetUserTournament_Data_Tournament `json:"tournament"`
	Typename   string     `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_Tournament struct {
	Name      string `json:"name"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Typename  string `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_MasterUser struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

type CompetitionCoalitionGetUserTournament_Data_LastTournamentResult struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CompetitionCoalitionGetUserTournament(variables CompetitionCoalitionGetUserTournament_Variables) (CompetitionCoalitionGetUserTournament_Data, error) {
	request := gql.NewQueryRequest[CompetitionCoalitionGetUserTournament_Variables](
		"query competitionCoalitionGetUserTournament {\n  student {\n    getUserTournamentWidget {\n      coalitionMember {\n        coalition {\n          ...CompetitionCurrentCoalition\n          __typename\n        }\n        __typename\n      }\n      lastTournamentResult {\n        id\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CompetitionCurrentCoalition on GameCoalition {\n  id\n  name\n  avatarUrl\n  backgroundUrl\n  backgroundUrlBig\n  memberCount\n  color\n  currentTournament {\n    points\n    tournament {\n      name\n      timeStart\n      timeEnd\n      __typename\n    }\n    __typename\n  }\n  masterUser {\n    login\n    avatarUrl\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CompetitionCoalitionGetUserTournament_Data](ctx, request)
}
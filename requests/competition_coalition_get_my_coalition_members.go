package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CompetitionCoalitionGetMyCoalitionMembers struct {
	Variables_Page_CompetitionCoalitionGetMyCoalitionMembers Variables_Page_CompetitionCoalitionGetMyCoalitionMembers `json:"page"`
}

type Variables_Page_CompetitionCoalitionGetMyCoalitionMembers struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_CompetitionCoalitionGetMyCoalitionMembers struct {
	Data_Student_CompetitionCoalitionGetMyCoalitionMembers Data_Student_CompetitionCoalitionGetMyCoalitionMembers `json:"student"`
}

type Data_Student_CompetitionCoalitionGetMyCoalitionMembers struct {
	Data_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers Data_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Data_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers struct {
	GetMyCoalitionMembers []Data_GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers `json:"getMyCoalitionMembers"`
	Typename              string                 `json:"__typename"`
}

type Data_GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers struct {
	Data_User_CompetitionCoalitionGetMyCoalitionMembers     Data_User_CompetitionCoalitionGetMyCoalitionMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Data_UserExperience_CompetitionCoalitionGetMyCoalitionMembers Data_UserExperience_CompetitionCoalitionGetMyCoalitionMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_CompetitionCoalitionGetMyCoalitionMembers struct {
	Data_Level_CompetitionCoalitionGetMyCoalitionMembers    Data_Level_CompetitionCoalitionGetMyCoalitionMembers  `json:"level"`
	Typename string `json:"__typename"`
}

type Data_Level_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) CompetitionCoalitionGetMyCoalitionMembers(variables Variables_CompetitionCoalitionGetMyCoalitionMembers) (Data_CompetitionCoalitionGetMyCoalitionMembers, error) {
	request := gql.NewQueryRequest[Variables_CompetitionCoalitionGetMyCoalitionMembers](
		"query competitionCoalitionGetMyCoalitionMembers($page: PagingInput) {\n  student {\n    getUserTournamentWidget {\n      getMyCoalitionMembers(page: $page) {\n        user {\n          id\n          login\n          avatarUrl\n          userExperience {\n            level {\n              id\n              levelCode\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_CompetitionCoalitionGetMyCoalitionMembers](ctx, request)
}
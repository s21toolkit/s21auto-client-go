package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CompetitionCoalitionGetMyCoalitionMembers struct {
	Request_Page_CompetitionCoalitionGetMyCoalitionMembers Request_Page_CompetitionCoalitionGetMyCoalitionMembers `json:"page"`
}

type Request_Page_CompetitionCoalitionGetMyCoalitionMembers struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Response_Data_CompetitionCoalitionGetMyCoalitionMembers struct {
	Response_Student_CompetitionCoalitionGetMyCoalitionMembers Response_Student_CompetitionCoalitionGetMyCoalitionMembers `json:"student"`
}

type Response_Student_CompetitionCoalitionGetMyCoalitionMembers struct {
	Response_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers Response_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type Response_GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers struct {
	GetMyCoalitionMembers []Response_GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers `json:"getMyCoalitionMembers"`
	Typename              string                 `json:"__typename"`
}

type Response_GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers struct {
	Response_User_CompetitionCoalitionGetMyCoalitionMembers     Response_User_CompetitionCoalitionGetMyCoalitionMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Response_UserExperience_CompetitionCoalitionGetMyCoalitionMembers Response_UserExperience_CompetitionCoalitionGetMyCoalitionMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Response_UserExperience_CompetitionCoalitionGetMyCoalitionMembers struct {
	Response_Level_CompetitionCoalitionGetMyCoalitionMembers    Response_Level_CompetitionCoalitionGetMyCoalitionMembers  `json:"level"`
	Typename string `json:"__typename"`
}

type Response_Level_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) CompetitionCoalitionGetMyCoalitionMembers(variables Request_Variables_CompetitionCoalitionGetMyCoalitionMembers) (Response_Data_CompetitionCoalitionGetMyCoalitionMembers, error) {
	request := gql.NewQueryRequest[Request_Variables_CompetitionCoalitionGetMyCoalitionMembers](
		"query competitionCoalitionGetMyCoalitionMembers($page: PagingInput) {\n  student {\n    getUserTournamentWidget {\n      getMyCoalitionMembers(page: $page) {\n        user {\n          id\n          login\n          avatarUrl\n          userExperience {\n            level {\n              id\n              levelCode\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CompetitionCoalitionGetMyCoalitionMembers](ctx, request)
}
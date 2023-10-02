package requests

import "github.com/s21toolkit/s21client/gql"

type CompetitionCoalitionGetMyCoalitionMembers_Variables struct {
	Page CompetitionCoalitionGetMyCoalitionMembers_Variables_Page `json:"page"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Variables_Page struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type CompetitionCoalitionGetMyCoalitionMembers_Data struct {
	Student CompetitionCoalitionGetMyCoalitionMembers_Data_Student `json:"student"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_Student struct {
	GetUserTournamentWidget CompetitionCoalitionGetMyCoalitionMembers_Data_GetUserTournamentWidget `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_GetUserTournamentWidget struct {
	GetMyCoalitionMembers []CompetitionCoalitionGetMyCoalitionMembers_Data_GetMyCoalitionMember `json:"getMyCoalitionMembers"`
	Typename              string                 `json:"__typename"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_GetMyCoalitionMember struct {
	User     CompetitionCoalitionGetMyCoalitionMembers_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_User struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience CompetitionCoalitionGetMyCoalitionMembers_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_UserExperience struct {
	Level    CompetitionCoalitionGetMyCoalitionMembers_Data_Level  `json:"level"`
	Typename string `json:"__typename"`
}

type CompetitionCoalitionGetMyCoalitionMembers_Data_Level struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) CompetitionCoalitionGetMyCoalitionMembers(variables CompetitionCoalitionGetMyCoalitionMembers_Variables) (CompetitionCoalitionGetMyCoalitionMembers_Data, error) {
	request := gql.NewQueryRequest[CompetitionCoalitionGetMyCoalitionMembers_Variables](
		"query competitionCoalitionGetMyCoalitionMembers($page: PagingInput) {\n  student {\n    getUserTournamentWidget {\n      getMyCoalitionMembers(page: $page) {\n        user {\n          id\n          login\n          avatarUrl\n          userExperience {\n            level {\n              id\n              levelCode\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[CompetitionCoalitionGetMyCoalitionMembers_Data](ctx, request)
}
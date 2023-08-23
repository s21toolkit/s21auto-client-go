package requests

import "s21client/gql"

type Variables_CompetitionCoalitionGetMyCoalitionMembers struct {
	Page_CompetitionCoalitionGetMyCoalitionMembers Page_CompetitionCoalitionGetMyCoalitionMembers `json:"page"`
}

type Page_CompetitionCoalitionGetMyCoalitionMembers struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_CompetitionCoalitionGetMyCoalitionMembers struct {
	Student_CompetitionCoalitionGetMyCoalitionMembers Student_CompetitionCoalitionGetMyCoalitionMembers `json:"student"`
}

type Student_CompetitionCoalitionGetMyCoalitionMembers struct {
	GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers `json:"getUserTournamentWidget"`
	Typename                string                  `json:"__typename"`
}

type GetUserTournamentWidget_CompetitionCoalitionGetMyCoalitionMembers struct {
	GetMyCoalitionMembers []GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers `json:"getMyCoalitionMembers"`
	Typename              string                 `json:"__typename"`
}

type GetMyCoalitionMember_CompetitionCoalitionGetMyCoalitionMembers struct {
	User_CompetitionCoalitionGetMyCoalitionMembers     User_CompetitionCoalitionGetMyCoalitionMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type User_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	UserExperience_CompetitionCoalitionGetMyCoalitionMembers UserExperience_CompetitionCoalitionGetMyCoalitionMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type UserExperience_CompetitionCoalitionGetMyCoalitionMembers struct {
	Level_CompetitionCoalitionGetMyCoalitionMembers    Level_CompetitionCoalitionGetMyCoalitionMembers  `json:"level"`
	Typename string `json:"__typename"`
}

type Level_CompetitionCoalitionGetMyCoalitionMembers struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) CompetitionCoalitionGetMyCoalitionMembers(variables Variables_CompetitionCoalitionGetMyCoalitionMembers) (Data_CompetitionCoalitionGetMyCoalitionMembers, error) {
	request := gql.NewQueryRequest[Variables_CompetitionCoalitionGetMyCoalitionMembers](
		"query competitionCoalitionGetMyCoalitionMembers($page: PagingInput) {   student {     getUserTournamentWidget {       getMyCoalitionMembers(page: $page) {         user {           id           login           avatarUrl           userExperience {             level {               id               levelCode               __typename             }             __typename           }           __typename         }         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_CompetitionCoalitionGetMyCoalitionMembers](ctx, request)
}
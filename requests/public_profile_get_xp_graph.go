package requests

import "s21client/gql"

type Variables_PublicProfileGetXpGraph struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetXpGraph struct {
	Student_PublicProfileGetXpGraph Student_PublicProfileGetXpGraph `json:"student"`
}

type Student_PublicProfileGetXpGraph struct {
	GetExperienceHistoryDate_PublicProfileGetXpGraph GetExperienceHistoryDate_PublicProfileGetXpGraph `json:"getExperienceHistoryDate"`
	Typename                 string                   `json:"__typename"`
}

type GetExperienceHistoryDate_PublicProfileGetXpGraph struct {
	History_PublicProfileGetXpGraph  []History_PublicProfileGetXpGraph `json:"history"`
	Typename string    `json:"__typename"`
}

type History_PublicProfileGetXpGraph struct {
	AwardDate string `json:"awardDate"`
	ExpValue  int64  `json:"expValue"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetXpGraph(variables Variables_PublicProfileGetXpGraph) (Data_PublicProfileGetXpGraph, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetXpGraph](
		"query publicProfileGetXpGraph($userId: UUID!) {   student {     getExperienceHistoryDate(userId: $userId) {       history {         awardDate         expValue         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetXpGraph](ctx, request)
}
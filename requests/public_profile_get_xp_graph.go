package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileGetXpGraph struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetXpGraph struct {
	Data_Student_PublicProfileGetXpGraph Data_Student_PublicProfileGetXpGraph `json:"student"`
}

type Data_Student_PublicProfileGetXpGraph struct {
	Data_GetExperienceHistoryDate_PublicProfileGetXpGraph Data_GetExperienceHistoryDate_PublicProfileGetXpGraph `json:"getExperienceHistoryDate"`
	Typename                 string                   `json:"__typename"`
}

type Data_GetExperienceHistoryDate_PublicProfileGetXpGraph struct {
	Data_History_PublicProfileGetXpGraph  []Data_History_PublicProfileGetXpGraph `json:"history"`
	Typename string    `json:"__typename"`
}

type Data_History_PublicProfileGetXpGraph struct {
	AwardDate string `json:"awardDate"`
	ExpValue  int64  `json:"expValue"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetXpGraph(variables Variables_PublicProfileGetXpGraph) (Data_PublicProfileGetXpGraph, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetXpGraph](
		"query publicProfileGetXpGraph($userId: UUID!) {\n  student {\n    getExperienceHistoryDate(userId: $userId) {\n      history {\n        awardDate\n        expValue\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetXpGraph](ctx, request)
}
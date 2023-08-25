package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetXpGraph struct {
	UserID string `json:"userId"`
}


type Response_Data_PublicProfileGetXpGraph struct {
	Response_Student_PublicProfileGetXpGraph Response_Student_PublicProfileGetXpGraph `json:"student"`
}

type Response_Student_PublicProfileGetXpGraph struct {
	Response_GetExperienceHistoryDate_PublicProfileGetXpGraph Response_GetExperienceHistoryDate_PublicProfileGetXpGraph `json:"getExperienceHistoryDate"`
	Typename                 string                   `json:"__typename"`
}

type Response_GetExperienceHistoryDate_PublicProfileGetXpGraph struct {
	Response_History_PublicProfileGetXpGraph  []Response_History_PublicProfileGetXpGraph `json:"history"`
	Typename string    `json:"__typename"`
}

type Response_History_PublicProfileGetXpGraph struct {
	AwardDate string `json:"awardDate"`
	ExpValue  int64  `json:"expValue"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetXpGraph(variables Request_Variables_PublicProfileGetXpGraph) (Response_Data_PublicProfileGetXpGraph, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetXpGraph](
		"query publicProfileGetXpGraph($userId: UUID!) {\n  student {\n    getExperienceHistoryDate(userId: $userId) {\n      history {\n        awardDate\n        expValue\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetXpGraph](ctx, request)
}
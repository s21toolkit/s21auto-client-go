package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetXpGraph_Variables struct {
	UserID string `json:"userId"`
}


type PublicProfileGetXpGraph_Data struct {
	Student PublicProfileGetXpGraph_Data_Student `json:"student"`
}

type PublicProfileGetXpGraph_Data_Student struct {
	GetExperienceHistoryDate PublicProfileGetXpGraph_Data_GetExperienceHistoryDate `json:"getExperienceHistoryDate"`
	Typename                 string                   `json:"__typename"`
}

type PublicProfileGetXpGraph_Data_GetExperienceHistoryDate struct {
	History  []PublicProfileGetXpGraph_Data_History `json:"history"`
	Typename string    `json:"__typename"`
}

type PublicProfileGetXpGraph_Data_History struct {
	AwardDate string `json:"awardDate"`
	ExpValue  int64  `json:"expValue"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetXpGraph(variables PublicProfileGetXpGraph_Variables) (PublicProfileGetXpGraph_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetXpGraph_Variables](
		"query publicProfileGetXpGraph($userId: UUID!) {\n  student {\n    getExperienceHistoryDate(userId: $userId) {\n      history {\n        awardDate\n        expValue\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetXpGraph_Data](ctx, request)
}
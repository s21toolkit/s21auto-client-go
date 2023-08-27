package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileLoadAverageLogtime struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type Data_PublicProfileLoadAverageLogtime struct {
	Data_School21_PublicProfileLoadAverageLogtime Data_School21_PublicProfileLoadAverageLogtime `json:"school21"`
}

type Data_School21_PublicProfileLoadAverageLogtime struct {
	Data_LoadAverageLogtime_PublicProfileLoadAverageLogtime Data_LoadAverageLogtime_PublicProfileLoadAverageLogtime `json:"loadAverageLogtime"`
	Typename           string             `json:"__typename"`
}

type Data_LoadAverageLogtime_PublicProfileLoadAverageLogtime struct {
	Week         int64  `json:"week"`
	Month        int64  `json:"month"`
	WeekPerMonth int64  `json:"weekPerMonth"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileLoadAverageLogtime(variables Variables_PublicProfileLoadAverageLogtime) (Data_PublicProfileLoadAverageLogtime, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileLoadAverageLogtime](
		"query publicProfileLoadAverageLogtime($login: String!, $schoolID: UUID!, $date: Date!) {\n  school21 {\n    loadAverageLogtime(login: $login, schoolID: $schoolID, date: $date) {\n      week\n      month\n      weekPerMonth\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileLoadAverageLogtime](ctx, request)
}
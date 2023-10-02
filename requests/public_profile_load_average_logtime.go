package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileLoadAverageLogtime_Variables struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type PublicProfileLoadAverageLogtime_Data struct {
	School21 PublicProfileLoadAverageLogtime_Data_School21 `json:"school21"`
}

type PublicProfileLoadAverageLogtime_Data_School21 struct {
	LoadAverageLogtime PublicProfileLoadAverageLogtime_Data_LoadAverageLogtime `json:"loadAverageLogtime"`
	Typename           string             `json:"__typename"`
}

type PublicProfileLoadAverageLogtime_Data_LoadAverageLogtime struct {
	Week         int64  `json:"week"`
	Month        int64  `json:"month"`
	WeekPerMonth int64  `json:"weekPerMonth"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileLoadAverageLogtime(variables PublicProfileLoadAverageLogtime_Variables) (PublicProfileLoadAverageLogtime_Data, error) {
	request := gql.NewQueryRequest[PublicProfileLoadAverageLogtime_Variables](
		"query publicProfileLoadAverageLogtime($login: String!, $schoolID: UUID!, $date: Date!) {\n  school21 {\n    loadAverageLogtime(login: $login, schoolID: $schoolID, date: $date) {\n      week\n      month\n      weekPerMonth\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileLoadAverageLogtime_Data](ctx, request)
}
package requests

import "s21client/gql"

type Variables_PublicProfileLoadAverageLogtime struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type Data_PublicProfileLoadAverageLogtime struct {
	School21_PublicProfileLoadAverageLogtime School21_PublicProfileLoadAverageLogtime `json:"school21"`
}

type School21_PublicProfileLoadAverageLogtime struct {
	LoadAverageLogtime_PublicProfileLoadAverageLogtime LoadAverageLogtime_PublicProfileLoadAverageLogtime `json:"loadAverageLogtime"`
	Typename           string             `json:"__typename"`
}

type LoadAverageLogtime_PublicProfileLoadAverageLogtime struct {
	Week         int64  `json:"week"`
	Month        int64  `json:"month"`
	WeekPerMonth int64  `json:"weekPerMonth"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileLoadAverageLogtime(variables Variables_PublicProfileLoadAverageLogtime) (Data_PublicProfileLoadAverageLogtime, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileLoadAverageLogtime](
		"query publicProfileLoadAverageLogtime($login: String!, $schoolID: UUID!, $date: Date!) {   school21 {     loadAverageLogtime(login: $login, schoolID: $schoolID, date: $date) {       week       month       weekPerMonth       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileLoadAverageLogtime](ctx, request)
}
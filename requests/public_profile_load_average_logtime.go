package requests

import "s21client/gql"

type Request_Variables_PublicProfileLoadAverageLogtime struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type Response_Data_PublicProfileLoadAverageLogtime struct {
	Response_School21_PublicProfileLoadAverageLogtime Response_School21_PublicProfileLoadAverageLogtime `json:"school21"`
}

type Response_School21_PublicProfileLoadAverageLogtime struct {
	Response_LoadAverageLogtime_PublicProfileLoadAverageLogtime Response_LoadAverageLogtime_PublicProfileLoadAverageLogtime `json:"loadAverageLogtime"`
	Typename           string             `json:"__typename"`
}

type Response_LoadAverageLogtime_PublicProfileLoadAverageLogtime struct {
	Week         int64  `json:"week"`
	Month        int64  `json:"month"`
	WeekPerMonth int64  `json:"weekPerMonth"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileLoadAverageLogtime(variables Request_Variables_PublicProfileLoadAverageLogtime) (Response_Data_PublicProfileLoadAverageLogtime, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileLoadAverageLogtime](
		"query publicProfileLoadAverageLogtime($login: String!, $schoolID: UUID!, $date: Date!) {   school21 {     loadAverageLogtime(login: $login, schoolID: $schoolID, date: $date) {       week       month       weekPerMonth       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileLoadAverageLogtime](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetStudentTraffic_Variables struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type PublicProfileGetStudentTraffic_Data struct {
	Student PublicProfileGetStudentTraffic_Data_Student `json:"student"`
}

type PublicProfileGetStudentTraffic_Data_Student struct {
	GetStudentTraffic PublicProfileGetStudentTraffic_Data_GetStudentTraffic `json:"getStudentTraffic"`
	Typename          string            `json:"__typename"`
}

type PublicProfileGetStudentTraffic_Data_GetStudentTraffic struct {
	Days      []PublicProfileGetStudentTraffic_Data_Day  `json:"days"`
	EndDate   string `json:"endDate"`
	StartDate string `json:"startDate"`
	Typename  string `json:"__typename"`
}

type PublicProfileGetStudentTraffic_Data_Day struct {
	Date               string `json:"date"`
	PeriodOnCampus     string `json:"periodOnCampus"`
	PeriodAuthorizSDP  string `json:"periodAuthorizSDP"`
	PeriodAuthorizIMAC string `json:"periodAuthorizIMac"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetStudentTraffic(variables PublicProfileGetStudentTraffic_Variables) (PublicProfileGetStudentTraffic_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetStudentTraffic_Variables](
		"query publicProfileGetStudentTraffic($login: String!, $schoolID: UUID!, $date: Date!) {\n  student {\n    getStudentTraffic(login: $login, schoolID: $schoolID, date: $date) {\n      days {\n        date\n        periodOnCampus\n        periodAuthorizSDP\n        periodAuthorizIMac\n        __typename\n      }\n      endDate\n      startDate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetStudentTraffic_Data](ctx, request)
}
package requests

import "s21client/gql"

type Variables_PublicProfileGetStudentTraffic struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type Data_PublicProfileGetStudentTraffic struct {
	Student_PublicProfileGetStudentTraffic Student_PublicProfileGetStudentTraffic `json:"student"`
}

type Student_PublicProfileGetStudentTraffic struct {
	GetStudentTraffic_PublicProfileGetStudentTraffic GetStudentTraffic_PublicProfileGetStudentTraffic `json:"getStudentTraffic"`
	Typename          string            `json:"__typename"`
}

type GetStudentTraffic_PublicProfileGetStudentTraffic struct {
	Days      []interface{} `json:"days"`
	EndDate   string        `json:"endDate"`
	StartDate string        `json:"startDate"`
	Typename  string        `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetStudentTraffic(variables Variables_PublicProfileGetStudentTraffic) (Data_PublicProfileGetStudentTraffic, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetStudentTraffic](
		"query publicProfileGetStudentTraffic($login: String!, $schoolID: UUID!, $date: Date!) {   student {     getStudentTraffic(login: $login, schoolID: $schoolID, date: $date) {       days {         date         periodOnCampus         periodAuthorizSDP         periodAuthorizIMac         __typename       }       endDate       startDate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetStudentTraffic](ctx, request)
}
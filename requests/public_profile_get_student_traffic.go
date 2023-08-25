package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetStudentTraffic struct {
	Login    string `json:"login"`
	SchoolID string `json:"schoolID"`
	Date     string `json:"date"`
}


type Response_Data_PublicProfileGetStudentTraffic struct {
	Response_Student_PublicProfileGetStudentTraffic Response_Student_PublicProfileGetStudentTraffic `json:"student"`
}

type Response_Student_PublicProfileGetStudentTraffic struct {
	Response_GetStudentTraffic_PublicProfileGetStudentTraffic Response_GetStudentTraffic_PublicProfileGetStudentTraffic `json:"getStudentTraffic"`
	Typename          string            `json:"__typename"`
}

type Response_GetStudentTraffic_PublicProfileGetStudentTraffic struct {
	Days      []interface{} `json:"days"`
	EndDate   string        `json:"endDate"`
	StartDate string        `json:"startDate"`
	Typename  string        `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetStudentTraffic(variables Request_Variables_PublicProfileGetStudentTraffic) (Response_Data_PublicProfileGetStudentTraffic, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetStudentTraffic](
		"query publicProfileGetStudentTraffic($login: String!, $schoolID: UUID!, $date: Date!) {   student {     getStudentTraffic(login: $login, schoolID: $schoolID, date: $date) {       days {         date         periodOnCampus         periodAuthorizSDP         periodAuthorizIMac         __typename       }       endDate       startDate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetStudentTraffic](ctx, request)
}
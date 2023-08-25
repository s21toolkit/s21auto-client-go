package requests

import "s21client/gql"

type Request_Variables_GetDashboardWorkstation struct {
	Login string `json:"login"`
}


type Response_Data_GetDashboardWorkstation struct {
	Response_Student_GetDashboardWorkstation Response_Student_GetDashboardWorkstation `json:"student"`
}

type Response_Student_GetDashboardWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}

func (ctx *RequestContext) GetDashboardWorkstation(variables Request_Variables_GetDashboardWorkstation) (Response_Data_GetDashboardWorkstation, error) {
	request := gql.NewQueryRequest[Request_Variables_GetDashboardWorkstation](
		"query getDashboardWorkstation($login: String!) {   student {     getWorkstationByLogin(login: $login) {       id       classroomId       hostName       classroom {         floor         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetDashboardWorkstation](ctx, request)
}
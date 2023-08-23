package requests

import "s21client/gql"

type Variables_GetDashboardWorkstation struct {
	Login string `json:"login"`
}


type Data_GetDashboardWorkstation struct {
	Student_GetDashboardWorkstation Student_GetDashboardWorkstation `json:"student"`
}

type Student_GetDashboardWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}

func (ctx *RequestContext) GetDashboardWorkstation(variables Variables_GetDashboardWorkstation) (Data_GetDashboardWorkstation, error) {
	request := gql.NewQueryRequest[Variables_GetDashboardWorkstation](
		"query getDashboardWorkstation($login: String!) {   student {     getWorkstationByLogin(login: $login) {       id       classroomId       hostName       classroom {         floor         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetDashboardWorkstation](ctx, request)
}
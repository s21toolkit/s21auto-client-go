package requests

import "github.com/s21toolkit/s21client/gql"

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
		"query getDashboardWorkstation($login: String!) {\n  student {\n    getWorkstationByLogin(login: $login) {\n      id\n      classroomId\n      hostName\n      classroom {\n        floor\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetDashboardWorkstation](ctx, request)
}
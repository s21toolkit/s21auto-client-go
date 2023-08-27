package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetDashboardWorkstation struct {
	Login string `json:"login"`
}


type Data_GetDashboardWorkstation struct {
	Data_Student_GetDashboardWorkstation Data_Student_GetDashboardWorkstation `json:"student"`
}

type Data_Student_GetDashboardWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}


func (ctx *RequestContext) GetDashboardWorkstation(variables Variables_GetDashboardWorkstation) (Data_GetDashboardWorkstation, error) {
	request := gql.NewQueryRequest[Variables_GetDashboardWorkstation](
		"query getDashboardWorkstation($login: String!) {\n  student {\n    getWorkstationByLogin(login: $login) {\n      id\n      classroomId\n      hostName\n      classroom {\n        floor\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetDashboardWorkstation](ctx, request)
}
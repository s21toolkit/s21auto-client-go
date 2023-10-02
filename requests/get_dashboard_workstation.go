package requests

import "github.com/s21toolkit/s21client/gql"

type GetDashboardWorkstation_Variables struct {
	Login string `json:"login"`
}


type GetDashboardWorkstation_Data struct {
	Student GetDashboardWorkstation_Data_Student `json:"student"`
}

type GetDashboardWorkstation_Data_Student struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}


func (ctx *RequestContext) GetDashboardWorkstation(variables GetDashboardWorkstation_Variables) (GetDashboardWorkstation_Data, error) {
	request := gql.NewQueryRequest[GetDashboardWorkstation_Variables](
		"query getDashboardWorkstation($login: String!) {\n  student {\n    getWorkstationByLogin(login: $login) {\n      id\n      classroomId\n      hostName\n      classroom {\n        floor\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetDashboardWorkstation_Data](ctx, request)
}
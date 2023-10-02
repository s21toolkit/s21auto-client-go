package requests

import "github.com/s21toolkit/s21client/gql"

type GetCampusWorkstation_Variables struct {
	Login string `json:"login"`
}


type GetCampusWorkstation_Data struct {
	Student GetCampusWorkstation_Data_Student `json:"student"`
}

type GetCampusWorkstation_Data_Student struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}


func (ctx *RequestContext) GetCampusWorkstation(variables GetCampusWorkstation_Variables) (GetCampusWorkstation_Data, error) {
	request := gql.NewQueryRequest[GetCampusWorkstation_Variables](
		"query getCampusWorkstation($login: String!) {\n  student {\n    getWorkstationByLogin(login: $login) {\n      id\n      classroomId\n      hostName\n      workstationRow\n      workstationNumber\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCampusWorkstation_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCampusWorkstation struct {
	Login string `json:"login"`
}


type Data_GetCampusWorkstation struct {
	Data_Student_GetCampusWorkstation Data_Student_GetCampusWorkstation `json:"student"`
}

type Data_Student_GetCampusWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}


func (ctx *RequestContext) GetCampusWorkstation(variables Variables_GetCampusWorkstation) (Data_GetCampusWorkstation, error) {
	request := gql.NewQueryRequest[Variables_GetCampusWorkstation](
		"query getCampusWorkstation($login: String!) {\n  student {\n    getWorkstationByLogin(login: $login) {\n      id\n      classroomId\n      hostName\n      workstationRow\n      workstationNumber\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCampusWorkstation](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetCampusWorkstation struct {
	Login string `json:"login"`
}


type Response_Data_GetCampusWorkstation struct {
	Response_Student_GetCampusWorkstation Response_Student_GetCampusWorkstation `json:"student"`
}

type Response_Student_GetCampusWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}

func (ctx *RequestContext) GetCampusWorkstation(variables Request_Variables_GetCampusWorkstation) (Response_Data_GetCampusWorkstation, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCampusWorkstation](
		"query getCampusWorkstation($login: String!) {   student {     getWorkstationByLogin(login: $login) {       id       classroomId       hostName       workstationRow       workstationNumber       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetCampusWorkstation](ctx, request)
}
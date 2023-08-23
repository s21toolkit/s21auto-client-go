package requests

import "s21client/gql"

type Variables_GetCampusWorkstation struct {
	Login string `json:"login"`
}


type Data_GetCampusWorkstation struct {
	Student_GetCampusWorkstation Student_GetCampusWorkstation `json:"student"`
}

type Student_GetCampusWorkstation struct {
	GetWorkstationByLogin interface{} `json:"getWorkstationByLogin"`
	Typename              string      `json:"__typename"`
}

func (ctx *RequestContext) GetCampusWorkstation(variables Variables_GetCampusWorkstation) (Data_GetCampusWorkstation, error) {
	request := gql.NewQueryRequest[Variables_GetCampusWorkstation](
		"query getCampusWorkstation($login: String!) {   student {     getWorkstationByLogin(login: $login) {       id       classroomId       hostName       workstationRow       workstationNumber       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCampusWorkstation](ctx, request)
}
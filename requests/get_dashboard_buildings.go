package requests

import "s21client/gql"

type Request_Variables_GetDashboardBuildings struct {
}


type Response_Data_GetDashboardBuildings struct {
	Response_Student_GetDashboardBuildings Response_Student_GetDashboardBuildings `json:"student"`
}

type Response_Student_GetDashboardBuildings struct {
	GetBuildings []Response_GetBuilding_GetDashboardBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type Response_GetBuilding_GetDashboardBuildings struct {
	ID         string      `json:"id"`
	Classrooms []Response_Classroom_GetDashboardBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Response_Classroom_GetDashboardBuildings struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetDashboardBuildings(variables Request_Variables_GetDashboardBuildings) (Response_Data_GetDashboardBuildings, error) {
	request := gql.NewQueryRequest[Request_Variables_GetDashboardBuildings](
		"query getDashboardBuildings {   student {     getBuildings {       id       classrooms {         id         number         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetDashboardBuildings](ctx, request)
}
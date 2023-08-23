package requests

import "s21client/gql"

type Variables_GetDashboardBuildings struct {
}


type Data_GetDashboardBuildings struct {
	Student_GetDashboardBuildings Student_GetDashboardBuildings `json:"student"`
}

type Student_GetDashboardBuildings struct {
	GetBuildings []GetBuilding_GetDashboardBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type GetBuilding_GetDashboardBuildings struct {
	ID         string      `json:"id"`
	Classrooms []Classroom_GetDashboardBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Classroom_GetDashboardBuildings struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetDashboardBuildings(variables Variables_GetDashboardBuildings) (Data_GetDashboardBuildings, error) {
	request := gql.NewQueryRequest[Variables_GetDashboardBuildings](
		"query getDashboardBuildings {   student {     getBuildings {       id       classrooms {         id         number         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetDashboardBuildings](ctx, request)
}
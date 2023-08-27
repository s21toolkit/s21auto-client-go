package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetDashboardBuildings struct {
}


type Data_GetDashboardBuildings struct {
	Data_Student_GetDashboardBuildings Data_Student_GetDashboardBuildings `json:"student"`
}

type Data_Student_GetDashboardBuildings struct {
	GetBuildings []Data_GetBuilding_GetDashboardBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type Data_GetBuilding_GetDashboardBuildings struct {
	ID         string      `json:"id"`
	Classrooms []Data_Classroom_GetDashboardBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Data_Classroom_GetDashboardBuildings struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetDashboardBuildings(variables Variables_GetDashboardBuildings) (Data_GetDashboardBuildings, error) {
	request := gql.NewQueryRequest[Variables_GetDashboardBuildings](
		"query getDashboardBuildings {\n  student {\n    getBuildings {\n      id\n      classrooms {\n        id\n        number\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetDashboardBuildings](ctx, request)
}
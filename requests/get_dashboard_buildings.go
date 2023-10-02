package requests

import "github.com/s21toolkit/s21client/gql"

type GetDashboardBuildings_Variables struct {
}


type GetDashboardBuildings_Data struct {
	Student GetDashboardBuildings_Data_Student `json:"student"`
}

type GetDashboardBuildings_Data_Student struct {
	GetBuildings []GetDashboardBuildings_Data_GetBuilding `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type GetDashboardBuildings_Data_GetBuilding struct {
	ID         string      `json:"id"`
	Classrooms []GetDashboardBuildings_Data_Classroom `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type GetDashboardBuildings_Data_Classroom struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetDashboardBuildings(variables GetDashboardBuildings_Variables) (GetDashboardBuildings_Data, error) {
	request := gql.NewQueryRequest[GetDashboardBuildings_Variables](
		"query getDashboardBuildings {\n  student {\n    getBuildings {\n      id\n      classrooms {\n        id\n        number\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetDashboardBuildings_Data](ctx, request)
}
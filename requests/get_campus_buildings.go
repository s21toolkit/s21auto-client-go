package requests

import "github.com/s21toolkit/s21client/gql"

type GetCampusBuildings_Variables struct {
}


type GetCampusBuildings_Data struct {
	Student GetCampusBuildings_Data_Student `json:"student"`
}

type GetCampusBuildings_Data_Student struct {
	GetBuildings []GetCampusBuildings_Data_GetBuilding `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type GetCampusBuildings_Data_GetBuilding struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Classrooms []GetCampusBuildings_Data_Classroom `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type GetCampusBuildings_Data_Classroom struct {
	ID                string        `json:"id"`
	Number            string        `json:"number"`
	Capacity          int64         `json:"capacity"`
	AvailableCapacity int64         `json:"availableCapacity"`
	Floor             int64         `json:"floor"`
	ClassroomPlan     GetCampusBuildings_Data_ClassroomPlan `json:"classroomPlan"`
	Specializations   []string      `json:"specializations"`
	Typename          string        `json:"__typename"`
}

type GetCampusBuildings_Data_ClassroomPlan struct {
	ClassroomPlanID string `json:"classroomPlanId"`
	PlanMeta        string `json:"planMeta"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusBuildings(variables GetCampusBuildings_Variables) (GetCampusBuildings_Data, error) {
	request := gql.NewQueryRequest[GetCampusBuildings_Variables](
		"query getCampusBuildings {\n  student {\n    getBuildings {\n      id\n      name\n      classrooms {\n        id\n        number\n        capacity\n        availableCapacity\n        floor\n        classroomPlan {\n          classroomPlanId\n          planMeta\n          __typename\n        }\n        specializations\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCampusBuildings_Data](ctx, request)
}
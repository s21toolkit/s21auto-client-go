package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCampusBuildings struct {
}


type Data_GetCampusBuildings struct {
	Data_Student_GetCampusBuildings Data_Student_GetCampusBuildings `json:"student"`
}

type Data_Student_GetCampusBuildings struct {
	GetBuildings []Data_GetBuilding_GetCampusBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type Data_GetBuilding_GetCampusBuildings struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Classrooms []Data_Classroom_GetCampusBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Data_Classroom_GetCampusBuildings struct {
	ID                string        `json:"id"`
	Number            string        `json:"number"`
	Capacity          int64         `json:"capacity"`
	AvailableCapacity int64         `json:"availableCapacity"`
	Floor             int64         `json:"floor"`
	Data_ClassroomPlan_GetCampusBuildings     Data_ClassroomPlan_GetCampusBuildings `json:"classroomPlan"`
	Specializations   []string      `json:"specializations"`
	Typename          string        `json:"__typename"`
}

type Data_ClassroomPlan_GetCampusBuildings struct {
	ClassroomPlanID string `json:"classroomPlanId"`
	PlanMeta        string `json:"planMeta"`
	Typename        string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusBuildings(variables Variables_GetCampusBuildings) (Data_GetCampusBuildings, error) {
	request := gql.NewQueryRequest[Variables_GetCampusBuildings](
		"query getCampusBuildings {\n  student {\n    getBuildings {\n      id\n      name\n      classrooms {\n        id\n        number\n        capacity\n        availableCapacity\n        floor\n        classroomPlan {\n          classroomPlanId\n          planMeta\n          __typename\n        }\n        specializations\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCampusBuildings](ctx, request)
}
package requests

import "s21client/gql"

type Variables_GetCampusBuildings struct {
}


type Data_GetCampusBuildings struct {
	Student_GetCampusBuildings Student_GetCampusBuildings `json:"student"`
}

type Student_GetCampusBuildings struct {
	GetBuildings []GetBuilding_GetCampusBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type GetBuilding_GetCampusBuildings struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Classrooms []Classroom_GetCampusBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Classroom_GetCampusBuildings struct {
	ID                string        `json:"id"`
	Number            string        `json:"number"`
	Capacity          int64         `json:"capacity"`
	AvailableCapacity int64         `json:"availableCapacity"`
	Floor             int64         `json:"floor"`
	ClassroomPlan_GetCampusBuildings     ClassroomPlan_GetCampusBuildings `json:"classroomPlan"`
	Specializations   []string      `json:"specializations"`
	Typename          string        `json:"__typename"`
}

type ClassroomPlan_GetCampusBuildings struct {
	ClassroomPlanID string `json:"classroomPlanId"`
	PlanMeta        string `json:"planMeta"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusBuildings(variables Variables_GetCampusBuildings) (Data_GetCampusBuildings, error) {
	request := gql.NewQueryRequest[Variables_GetCampusBuildings](
		"query getCampusBuildings {   student {     getBuildings {       id       name       classrooms {         id         number         capacity         availableCapacity         floor         classroomPlan {           classroomPlanId           planMeta           __typename         }         specializations         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCampusBuildings](ctx, request)
}
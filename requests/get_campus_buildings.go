package requests

import "s21client/gql"

type Request_Variables_GetCampusBuildings struct {
}


type Response_Data_GetCampusBuildings struct {
	Response_Student_GetCampusBuildings Response_Student_GetCampusBuildings `json:"student"`
}

type Response_Student_GetCampusBuildings struct {
	GetBuildings []Response_GetBuilding_GetCampusBuildings `json:"getBuildings"`
	Typename     string        `json:"__typename"`
}

type Response_GetBuilding_GetCampusBuildings struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Classrooms []Response_Classroom_GetCampusBuildings `json:"classrooms"`
	Typename   string      `json:"__typename"`
}

type Response_Classroom_GetCampusBuildings struct {
	ID                string        `json:"id"`
	Number            string        `json:"number"`
	Capacity          int64         `json:"capacity"`
	AvailableCapacity int64         `json:"availableCapacity"`
	Floor             int64         `json:"floor"`
	Response_ClassroomPlan_GetCampusBuildings     Response_ClassroomPlan_GetCampusBuildings `json:"classroomPlan"`
	Specializations   []string      `json:"specializations"`
	Typename          string        `json:"__typename"`
}

type Response_ClassroomPlan_GetCampusBuildings struct {
	ClassroomPlanID string `json:"classroomPlanId"`
	PlanMeta        string `json:"planMeta"`
	Typename        string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusBuildings(variables Request_Variables_GetCampusBuildings) (Response_Data_GetCampusBuildings, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCampusBuildings](
		"query getCampusBuildings {   student {     getBuildings {       id       name       classrooms {         id         number         capacity         availableCapacity         floor         classroomPlan {           classroomPlanId           planMeta           __typename         }         specializations         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetCampusBuildings](ctx, request)
}
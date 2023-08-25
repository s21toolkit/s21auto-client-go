package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetGraphCurrentType struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type Response_Data_GetGraphCurrentType struct {
	Response_Student_GetGraphCurrentType Response_Student_GetGraphCurrentType `json:"student"`
}

type Response_Student_GetGraphCurrentType struct {
	Response_GetStudentCurrentStageWithGraphType_GetGraphCurrentType Response_GetStudentCurrentStageWithGraphType_GetGraphCurrentType `json:"getStudentCurrentStageWithGraphType"`
	Typename                            string                              `json:"__typename"`
}

type Response_GetStudentCurrentStageWithGraphType_GetGraphCurrentType struct {
	ID        string `json:"id"`
	GraphType string `json:"graphType"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetGraphCurrentType(variables Request_Variables_GetGraphCurrentType) (Response_Data_GetGraphCurrentType, error) {
	request := gql.NewQueryRequest[Request_Variables_GetGraphCurrentType](
		"query getGraphCurrentType($userId: ID!, $schoolId: ID!) {   student {     getStudentCurrentStageWithGraphType(userId: $userId, schoolId: $schoolId) {       id       graphType       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetGraphCurrentType](ctx, request)
}
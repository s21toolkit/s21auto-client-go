package requests

import "s21client/gql"

type Variables_GetGraphCurrentType struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type Data_GetGraphCurrentType struct {
	Student_GetGraphCurrentType Student_GetGraphCurrentType `json:"student"`
}

type Student_GetGraphCurrentType struct {
	GetStudentCurrentStageWithGraphType_GetGraphCurrentType GetStudentCurrentStageWithGraphType_GetGraphCurrentType `json:"getStudentCurrentStageWithGraphType"`
	Typename                            string                              `json:"__typename"`
}

type GetStudentCurrentStageWithGraphType_GetGraphCurrentType struct {
	ID        string `json:"id"`
	GraphType string `json:"graphType"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetGraphCurrentType(variables Variables_GetGraphCurrentType) (Data_GetGraphCurrentType, error) {
	request := gql.NewQueryRequest[Variables_GetGraphCurrentType](
		"query getGraphCurrentType($userId: ID!, $schoolId: ID!) {   student {     getStudentCurrentStageWithGraphType(userId: $userId, schoolId: $schoolId) {       id       graphType       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetGraphCurrentType](ctx, request)
}
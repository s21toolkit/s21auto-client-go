package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetGraphCurrentType struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type Data_GetGraphCurrentType struct {
	Data_Student_GetGraphCurrentType Data_Student_GetGraphCurrentType `json:"student"`
}

type Data_Student_GetGraphCurrentType struct {
	Data_GetStudentCurrentStageWithGraphType_GetGraphCurrentType Data_GetStudentCurrentStageWithGraphType_GetGraphCurrentType `json:"getStudentCurrentStageWithGraphType"`
	Typename                            string                              `json:"__typename"`
}

type Data_GetStudentCurrentStageWithGraphType_GetGraphCurrentType struct {
	ID        string `json:"id"`
	GraphType string `json:"graphType"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetGraphCurrentType(variables Variables_GetGraphCurrentType) (Data_GetGraphCurrentType, error) {
	request := gql.NewQueryRequest[Variables_GetGraphCurrentType](
		"query getGraphCurrentType($userId: ID!, $schoolId: ID!) {\n  student {\n    getStudentCurrentStageWithGraphType(userId: $userId, schoolId: $schoolId) {\n      id\n      graphType\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetGraphCurrentType](ctx, request)
}
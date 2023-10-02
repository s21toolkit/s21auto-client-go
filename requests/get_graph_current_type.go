package requests

import "github.com/s21toolkit/s21client/gql"

type GetGraphCurrentType_Variables struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type GetGraphCurrentType_Data struct {
	Student GetGraphCurrentType_Data_Student `json:"student"`
}

type GetGraphCurrentType_Data_Student struct {
	GetStudentCurrentStageWithGraphType GetGraphCurrentType_Data_GetStudentCurrentStageWithGraphType `json:"getStudentCurrentStageWithGraphType"`
	Typename                            string                              `json:"__typename"`
}

type GetGraphCurrentType_Data_GetStudentCurrentStageWithGraphType struct {
	ID        string `json:"id"`
	GraphType string `json:"graphType"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetGraphCurrentType(variables GetGraphCurrentType_Variables) (GetGraphCurrentType_Data, error) {
	request := gql.NewQueryRequest[GetGraphCurrentType_Variables](
		"query getGraphCurrentType($userId: ID!, $schoolId: ID!) {\n  student {\n    getStudentCurrentStageWithGraphType(userId: $userId, schoolId: $schoolId) {\n      id\n      graphType\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetGraphCurrentType_Data](ctx, request)
}
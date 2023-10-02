package requests

import "github.com/s21toolkit/s21client/gql"

type GetStudentIsDeadlinesEnabled_Variables struct {
}


type GetStudentIsDeadlinesEnabled_Data struct {
	Student GetStudentIsDeadlinesEnabled_Data_Student `json:"student"`
}

type GetStudentIsDeadlinesEnabled_Data_Student struct {
	IsDeadlinesEnabled bool   `json:"isDeadlinesEnabled"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetStudentIsDeadlinesEnabled(variables GetStudentIsDeadlinesEnabled_Variables) (GetStudentIsDeadlinesEnabled_Data, error) {
	request := gql.NewQueryRequest[GetStudentIsDeadlinesEnabled_Variables](
		"query getStudentIsDeadlinesEnabled {\n  student {\n    isDeadlinesEnabled\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetStudentIsDeadlinesEnabled_Data](ctx, request)
}
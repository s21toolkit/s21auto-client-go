package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetStudentIsDeadlinesEnabled struct {
}


type Data_GetStudentIsDeadlinesEnabled struct {
	Data_Student_GetStudentIsDeadlinesEnabled Data_Student_GetStudentIsDeadlinesEnabled `json:"student"`
}

type Data_Student_GetStudentIsDeadlinesEnabled struct {
	IsDeadlinesEnabled bool   `json:"isDeadlinesEnabled"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetStudentIsDeadlinesEnabled(variables Variables_GetStudentIsDeadlinesEnabled) (Data_GetStudentIsDeadlinesEnabled, error) {
	request := gql.NewQueryRequest[Variables_GetStudentIsDeadlinesEnabled](
		"query getStudentIsDeadlinesEnabled {\n  student {\n    isDeadlinesEnabled\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetStudentIsDeadlinesEnabled](ctx, request)
}
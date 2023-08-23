package requests

import "s21client/gql"

type Variables_GetStudentIsDeadlinesEnabled struct {
}


type Data_GetStudentIsDeadlinesEnabled struct {
	Student_GetStudentIsDeadlinesEnabled Student_GetStudentIsDeadlinesEnabled `json:"student"`
}

type Student_GetStudentIsDeadlinesEnabled struct {
	IsDeadlinesEnabled bool   `json:"isDeadlinesEnabled"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetStudentIsDeadlinesEnabled(variables Variables_GetStudentIsDeadlinesEnabled) (Data_GetStudentIsDeadlinesEnabled, error) {
	request := gql.NewQueryRequest[Variables_GetStudentIsDeadlinesEnabled](
		"query getStudentIsDeadlinesEnabled {   student {     isDeadlinesEnabled     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetStudentIsDeadlinesEnabled](ctx, request)
}
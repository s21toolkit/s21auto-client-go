package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetStudentIsDeadlinesEnabled struct {
}


type Response_Data_GetStudentIsDeadlinesEnabled struct {
	Response_Student_GetStudentIsDeadlinesEnabled Response_Student_GetStudentIsDeadlinesEnabled `json:"student"`
}

type Response_Student_GetStudentIsDeadlinesEnabled struct {
	IsDeadlinesEnabled bool   `json:"isDeadlinesEnabled"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetStudentIsDeadlinesEnabled(variables Request_Variables_GetStudentIsDeadlinesEnabled) (Response_Data_GetStudentIsDeadlinesEnabled, error) {
	request := gql.NewQueryRequest[Request_Variables_GetStudentIsDeadlinesEnabled](
		"query getStudentIsDeadlinesEnabled {   student {     isDeadlinesEnabled     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetStudentIsDeadlinesEnabled](ctx, request)
}
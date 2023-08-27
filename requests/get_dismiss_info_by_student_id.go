package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetDismissInfoByStudentId struct {
	StudentID string `json:"studentId"`
}


type Data_GetDismissInfoByStudentId struct {
	Data_School21_GetDismissInfoByStudentId Data_School21_GetDismissInfoByStudentId `json:"school21"`
}

type Data_School21_GetDismissInfoByStudentId struct {
	GetDismissInfoByStudentID interface{} `json:"getDismissInfoByStudentId"`
	Typename                  string      `json:"__typename"`
}


func (ctx *RequestContext) GetDismissInfoByStudentId(variables Variables_GetDismissInfoByStudentId) (Data_GetDismissInfoByStudentId, error) {
	request := gql.NewQueryRequest[Variables_GetDismissInfoByStudentId](
		"query getDismissInfoByStudentId($studentId: UUID!) {\n  school21 {\n    getDismissInfoByStudentId(studentId: $studentId) {\n      dismissTypeId\n      dismissTs\n      lastStageGroupS21 {\n        waveId\n        waveName\n        eduForm\n        active\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetDismissInfoByStudentId](ctx, request)
}
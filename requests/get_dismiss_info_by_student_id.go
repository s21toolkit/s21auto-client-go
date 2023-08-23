package requests

import "s21client/gql"

type Variables_GetDismissInfoByStudentId struct {
	StudentID string `json:"studentId"`
}


type Data_GetDismissInfoByStudentId struct {
	School21_GetDismissInfoByStudentId School21_GetDismissInfoByStudentId `json:"school21"`
}

type School21_GetDismissInfoByStudentId struct {
	GetDismissInfoByStudentID interface{} `json:"getDismissInfoByStudentId"`
	Typename                  string      `json:"__typename"`
}

func (ctx *RequestContext) GetDismissInfoByStudentId(variables Variables_GetDismissInfoByStudentId) (Data_GetDismissInfoByStudentId, error) {
	request := gql.NewQueryRequest[Variables_GetDismissInfoByStudentId](
		"query getDismissInfoByStudentId($studentId: UUID!) {   school21 {     getDismissInfoByStudentId(studentId: $studentId) {       dismissTypeId       dismissTs       lastStageGroupS21 {         waveId         waveName         eduForm         active         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetDismissInfoByStudentId](ctx, request)
}
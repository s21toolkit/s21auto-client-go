package requests

import "s21client/gql"

type Request_Variables_GetDismissInfoByStudentId struct {
	StudentID string `json:"studentId"`
}


type Response_Data_GetDismissInfoByStudentId struct {
	Response_School21_GetDismissInfoByStudentId Response_School21_GetDismissInfoByStudentId `json:"school21"`
}

type Response_School21_GetDismissInfoByStudentId struct {
	GetDismissInfoByStudentID interface{} `json:"getDismissInfoByStudentId"`
	Typename                  string      `json:"__typename"`
}

func (ctx *RequestContext) GetDismissInfoByStudentId(variables Request_Variables_GetDismissInfoByStudentId) (Response_Data_GetDismissInfoByStudentId, error) {
	request := gql.NewQueryRequest[Request_Variables_GetDismissInfoByStudentId](
		"query getDismissInfoByStudentId($studentId: UUID!) {   school21 {     getDismissInfoByStudentId(studentId: $studentId) {       dismissTypeId       dismissTs       lastStageGroupS21 {         waveId         waveName         eduForm         active         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetDismissInfoByStudentId](ctx, request)
}
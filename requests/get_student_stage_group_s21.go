package requests

import "s21client/gql"

type Request_Variables_GetStudentStageGroupS21 struct {
	StudentID string `json:"studentId"`
}


type Response_Data_GetStudentStageGroupS21 struct {
	Response_Student_GetStudentStageGroupS21 Response_Student_GetStudentStageGroupS21 `json:"student"`
}

type Response_Student_GetStudentStageGroupS21 struct {
	Response_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 Response_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 `json:"getStageGroupS21PublicProfile"`
	Typename                      string                        `json:"__typename"`
}

type Response_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetStudentStageGroupS21(variables Request_Variables_GetStudentStageGroupS21) (Response_Data_GetStudentStageGroupS21, error) {
	request := gql.NewQueryRequest[Request_Variables_GetStudentStageGroupS21](
		"query getStudentStageGroupS21($studentId: UUID!) {   student {     getStageGroupS21PublicProfile(studentId: $studentId) {       waveId       waveName       eduForm       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetStudentStageGroupS21](ctx, request)
}
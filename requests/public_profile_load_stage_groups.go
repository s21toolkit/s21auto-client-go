package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileLoadStageGroups struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type Response_Data_PublicProfileLoadStageGroups struct {
	Response_School21_PublicProfileLoadStageGroups Response_School21_PublicProfileLoadStageGroups `json:"school21"`
}

type Response_School21_PublicProfileLoadStageGroups struct {
	Response_LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups []Response_LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups `json:"loadStudentStageGroupsS21PublicProfile"`
	Typename                               string                                   `json:"__typename"`
}

type Response_LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups struct {
	StageGroupStudentID string        `json:"stageGroupStudentId"`
	StudentID           string        `json:"studentId"`
	Response_StageGroupS21_PublicProfileLoadStageGroups       Response_StageGroupS21_PublicProfileLoadStageGroups `json:"stageGroupS21"`
	Response_SafeSchool_PublicProfileLoadStageGroups          Response_SafeSchool_PublicProfileLoadStageGroups    `json:"safeSchool"`
	Typename            string        `json:"__typename"`
}

type Response_SafeSchool_PublicProfileLoadStageGroups struct {
	FullName string `json:"fullName"`
	Typename string `json:"__typename"`
}

type Response_StageGroupS21_PublicProfileLoadStageGroups struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Active   bool   `json:"active"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileLoadStageGroups(variables Request_Variables_PublicProfileLoadStageGroups) (Response_Data_PublicProfileLoadStageGroups, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileLoadStageGroups](
		"query publicProfileLoadStageGroups($userId: UUID!, $schoolId: UUID!) {\n  school21 {\n    loadStudentStageGroupsS21PublicProfile(userId: $userId, schoolId: $schoolId) {\n      stageGroupStudentId\n      studentId\n      stageGroupS21 {\n        waveId\n        waveName\n        eduForm\n        active\n        __typename\n      }\n      safeSchool {\n        fullName\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileLoadStageGroups](ctx, request)
}
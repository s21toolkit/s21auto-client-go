package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileLoadStageGroups_Variables struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type PublicProfileLoadStageGroups_Data struct {
	School21 PublicProfileLoadStageGroups_Data_School21 `json:"school21"`
}

type PublicProfileLoadStageGroups_Data_School21 struct {
	LoadStudentStageGroupsS21PublicProfile []PublicProfileLoadStageGroups_Data_LoadStudentStageGroupsS21PublicProfile `json:"loadStudentStageGroupsS21PublicProfile"`
	Typename                               string                                   `json:"__typename"`
}

type PublicProfileLoadStageGroups_Data_LoadStudentStageGroupsS21PublicProfile struct {
	StageGroupStudentID string        `json:"stageGroupStudentId"`
	StudentID           string        `json:"studentId"`
	StageGroupS21       PublicProfileLoadStageGroups_Data_StageGroupS21 `json:"stageGroupS21"`
	SafeSchool          PublicProfileLoadStageGroups_Data_SafeSchool    `json:"safeSchool"`
	Typename            string        `json:"__typename"`
}

type PublicProfileLoadStageGroups_Data_SafeSchool struct {
	FullName string `json:"fullName"`
	Typename string `json:"__typename"`
}

type PublicProfileLoadStageGroups_Data_StageGroupS21 struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Active   bool   `json:"active"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileLoadStageGroups(variables PublicProfileLoadStageGroups_Variables) (PublicProfileLoadStageGroups_Data, error) {
	request := gql.NewQueryRequest[PublicProfileLoadStageGroups_Variables](
		"query publicProfileLoadStageGroups($userId: UUID!, $schoolId: UUID!) {\n  school21 {\n    loadStudentStageGroupsS21PublicProfile(userId: $userId, schoolId: $schoolId) {\n      stageGroupStudentId\n      studentId\n      stageGroupS21 {\n        waveId\n        waveName\n        eduForm\n        active\n        __typename\n      }\n      safeSchool {\n        fullName\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileLoadStageGroups_Data](ctx, request)
}
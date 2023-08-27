package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetStageInfo struct {
}


type Data_GetStageInfo struct {
	Data_User_GetStageInfo Data_User_GetStageInfo `json:"user"`
}

type Data_User_GetStageInfo struct {
	Data_GetCurrentUser_GetStageInfo          Data_GetCurrentUser_GetStageInfo            `json:"getCurrentUser"`
	Data_GetAllStagesTenantAware_GetStageInfo []Data_GetAllStagesTenantAware_GetStageInfo `json:"getAllStagesTenantAware"`
	Typename                string                    `json:"__typename"`
}

type Data_GetAllStagesTenantAware_GetStageInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type Data_GetCurrentUser_GetStageInfo struct {
	ID           string        `json:"id"`
	StudentRoles []Data_StudentRole_GetStageInfo `json:"studentRoles"`
	Typename     string        `json:"__typename"`
}

type Data_StudentRole_GetStageInfo struct {
	Status     string     `json:"status"`
	Data_School_GetStageInfo     Data_School_GetStageInfo     `json:"school"`
	Data_StageGroup_GetStageInfo Data_StageGroup_GetStageInfo `json:"stageGroup"`
	Typename   string     `json:"__typename"`
}

type Data_School_GetStageInfo struct {
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type Data_StageGroup_GetStageInfo struct {
	ClassSubjects []Data_ClassSubject_GetStageInfo `json:"classSubjects"`
	Name          string         `json:"name"`
	Data_Stage_GetStageInfo         int64          `json:"stage"`
	IsActive      bool           `json:"isActive"`
	Typename      string         `json:"__typename"`
}

type Data_ClassSubject_GetStageInfo struct {
	Data_Stage_GetStageInfo    Data_Stage_GetStageInfo  `json:"stage"`
	Typename string `json:"__typename"`
}

type Data_Stage_GetStageInfo struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetStageInfo(variables Variables_GetStageInfo) (Data_GetStageInfo, error) {
	request := gql.NewQueryRequest[Variables_GetStageInfo](
		"query getStageInfo {\n  user {\n    getCurrentUser {\n      id\n      studentRoles {\n        status\n        school {\n          organizationType\n          __typename\n        }\n        stageGroup {\n          classSubjects {\n            stage {\n              name\n              __typename\n            }\n            __typename\n          }\n          name\n          stage\n          isActive\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    getAllStagesTenantAware {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetStageInfo](ctx, request)
}
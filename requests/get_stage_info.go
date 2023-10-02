package requests

import "github.com/s21toolkit/s21client/gql"

type GetStageInfo_Variables struct {
}


type GetStageInfo_Data struct {
	User GetStageInfo_Data_User `json:"user"`
}

type GetStageInfo_Data_User struct {
	GetCurrentUser          GetStageInfo_Data_GetCurrentUser            `json:"getCurrentUser"`
	GetAllStagesTenantAware []GetStageInfo_Data_GetAllStagesTenantAware `json:"getAllStagesTenantAware"`
	Typename                string                    `json:"__typename"`
}

type GetStageInfo_Data_GetAllStagesTenantAware struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type GetStageInfo_Data_GetCurrentUser struct {
	ID           string        `json:"id"`
	StudentRoles []GetStageInfo_Data_StudentRole `json:"studentRoles"`
	Typename     string        `json:"__typename"`
}

type GetStageInfo_Data_StudentRole struct {
	Status     string     `json:"status"`
	School     GetStageInfo_Data_School     `json:"school"`
	StageGroup GetStageInfo_Data_StageGroup `json:"stageGroup"`
	Typename   string     `json:"__typename"`
}

type GetStageInfo_Data_School struct {
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type GetStageInfo_Data_StageGroup struct {
	ClassSubjects []GetStageInfo_Data_ClassSubject `json:"classSubjects"`
	Name          string         `json:"name"`
	Stage         int64          `json:"stage"`
	IsActive      bool           `json:"isActive"`
	Typename      string         `json:"__typename"`
}

type GetStageInfo_Data_ClassSubject struct {
	Stage    GetStageInfo_Data_Stage  `json:"stage"`
	Typename string `json:"__typename"`
}

type GetStageInfo_Data_Stage struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetStageInfo(variables GetStageInfo_Variables) (GetStageInfo_Data, error) {
	request := gql.NewQueryRequest[GetStageInfo_Variables](
		"query getStageInfo {\n  user {\n    getCurrentUser {\n      id\n      studentRoles {\n        status\n        school {\n          organizationType\n          __typename\n        }\n        stageGroup {\n          classSubjects {\n            stage {\n              name\n              __typename\n            }\n            __typename\n          }\n          name\n          stage\n          isActive\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    getAllStagesTenantAware {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetStageInfo_Data](ctx, request)
}
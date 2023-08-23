package requests

import "s21client/gql"

type Variables_GetStageInfo struct {
}


type Data_GetStageInfo struct {
	User_GetStageInfo User_GetStageInfo `json:"user"`
}

type User_GetStageInfo struct {
	GetCurrentUser_GetStageInfo          GetCurrentUser_GetStageInfo            `json:"getCurrentUser"`
	GetAllStagesTenantAware_GetStageInfo []GetAllStagesTenantAware_GetStageInfo `json:"getAllStagesTenantAware"`
	Typename                string                    `json:"__typename"`
}

type GetAllStagesTenantAware_GetStageInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type GetCurrentUser_GetStageInfo struct {
	ID           string        `json:"id"`
	StudentRoles []StudentRole_GetStageInfo `json:"studentRoles"`
	Typename     string        `json:"__typename"`
}

type StudentRole_GetStageInfo struct {
	Status     string     `json:"status"`
	School_GetStageInfo     School_GetStageInfo     `json:"school"`
	StageGroup_GetStageInfo StageGroup_GetStageInfo `json:"stageGroup"`
	Typename   string     `json:"__typename"`
}

type School_GetStageInfo struct {
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type StageGroup_GetStageInfo struct {
	ClassSubjects []ClassSubject_GetStageInfo `json:"classSubjects"`
	Name          string         `json:"name"`
	Stage_GetStageInfo         int64          `json:"stage"`
	IsActive      bool           `json:"isActive"`
	Typename      string         `json:"__typename"`
}

type ClassSubject_GetStageInfo struct {
	Stage_GetStageInfo    Stage_GetStageInfo  `json:"stage"`
	Typename string `json:"__typename"`
}

type Stage_GetStageInfo struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetStageInfo(variables Variables_GetStageInfo) (Data_GetStageInfo, error) {
	request := gql.NewQueryRequest[Variables_GetStageInfo](
		"query getStageInfo {   user {     getCurrentUser {       id       studentRoles {         status         school {           organizationType           __typename         }         stageGroup {           classSubjects {             stage {               name               __typename             }             __typename           }           name           stage           isActive           __typename         }         __typename       }       __typename     }     getAllStagesTenantAware {       id       name       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetStageInfo](ctx, request)
}
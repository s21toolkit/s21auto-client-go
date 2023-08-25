package requests

import "s21client/gql"

type Request_Variables_GetStageInfo struct {
}


type Response_Data_GetStageInfo struct {
	Response_User_GetStageInfo Response_User_GetStageInfo `json:"user"`
}

type Response_User_GetStageInfo struct {
	Response_GetCurrentUser_GetStageInfo          Response_GetCurrentUser_GetStageInfo            `json:"getCurrentUser"`
	Response_GetAllStagesTenantAware_GetStageInfo []Response_GetAllStagesTenantAware_GetStageInfo `json:"getAllStagesTenantAware"`
	Typename                string                    `json:"__typename"`
}

type Response_GetAllStagesTenantAware_GetStageInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type Response_GetCurrentUser_GetStageInfo struct {
	ID           string        `json:"id"`
	StudentRoles []Response_StudentRole_GetStageInfo `json:"studentRoles"`
	Typename     string        `json:"__typename"`
}

type Response_StudentRole_GetStageInfo struct {
	Status     string     `json:"status"`
	Response_School_GetStageInfo     Response_School_GetStageInfo     `json:"school"`
	Response_StageGroup_GetStageInfo Response_StageGroup_GetStageInfo `json:"stageGroup"`
	Typename   string     `json:"__typename"`
}

type Response_School_GetStageInfo struct {
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type Response_StageGroup_GetStageInfo struct {
	ClassSubjects []Response_ClassSubject_GetStageInfo `json:"classSubjects"`
	Name          string         `json:"name"`
	Response_Stage_GetStageInfo         int64          `json:"stage"`
	IsActive      bool           `json:"isActive"`
	Typename      string         `json:"__typename"`
}

type Response_ClassSubject_GetStageInfo struct {
	Response_Stage_GetStageInfo    Response_Stage_GetStageInfo  `json:"stage"`
	Typename string `json:"__typename"`
}

type Response_Stage_GetStageInfo struct {
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetStageInfo(variables Request_Variables_GetStageInfo) (Response_Data_GetStageInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_GetStageInfo](
		"query getStageInfo {   user {     getCurrentUser {       id       studentRoles {         status         school {           organizationType           __typename         }         stageGroup {           classSubjects {             stage {               name               __typename             }             __typename           }           name           stage           isActive           __typename         }         __typename       }       __typename     }     getAllStagesTenantAware {       id       name       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetStageInfo](ctx, request)
}
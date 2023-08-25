package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_UserRoleLoaderGetRoles struct {
}


type Response_Data_UserRoleLoaderGetRoles struct {
	Response_User_UserRoleLoaderGetRoles Response_User_UserRoleLoaderGetRoles `json:"user"`
}

type Response_User_UserRoleLoaderGetRoles struct {
	Response_GetCurrentUser_UserRoleLoaderGetRoles            Response_GetCurrentUser_UserRoleLoaderGetRoles             `json:"getCurrentUser"`
	GetCurrentUserSchoolRoles []Response_GetCurrentUserSchoolRole_UserRoleLoaderGetRoles `json:"getCurrentUserSchoolRoles"`
	Typename                  string                     `json:"__typename"`
}

type Response_GetCurrentUser_UserRoleLoaderGetRoles struct {
	FunctionalRoles       []Response_FunctionalRole_UserRoleLoaderGetRoles       `json:"functionalRoles"`
	ID                    string                 `json:"id"`
	StudentRoles          []Response_StudentRole_UserRoleLoaderGetRoles          `json:"studentRoles"`
	UserSchoolPermissions []Response_UserSchoolPermission_UserRoleLoaderGetRoles `json:"userSchoolPermissions"`
	SystemAdminRole       interface{}            `json:"systemAdminRole"`
	BusinessAdminRolesV2  []interface{}          `json:"businessAdminRolesV2"`
	Typename              string                 `json:"__typename"`
}

type Response_FunctionalRole_UserRoleLoaderGetRoles struct {
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

type Response_StudentRole_UserRoleLoaderGetRoles struct {
	ID       string `json:"id"`
	Response_School_UserRoleLoaderGetRoles   Response_School_UserRoleLoaderGetRoles `json:"school"`
	Status   string `json:"status"`
	Typename string `json:"__typename"`
}

type Response_School_UserRoleLoaderGetRoles struct {
	ID               string `json:"id"`
	ShortName        string `json:"shortName"`
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type Response_UserSchoolPermission_UserRoleLoaderGetRoles struct {
	SchoolID    string   `json:"schoolId"`
	Permissions []string `json:"permissions"`
	Typename    string   `json:"__typename"`
}

type Response_GetCurrentUserSchoolRole_UserRoleLoaderGetRoles struct {
	SchoolID string `json:"schoolId"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) UserRoleLoaderGetRoles(variables Request_Variables_UserRoleLoaderGetRoles) (Response_Data_UserRoleLoaderGetRoles, error) {
	request := gql.NewQueryRequest[Request_Variables_UserRoleLoaderGetRoles](
		"query userRoleLoaderGetRoles {   user {     getCurrentUser {       functionalRoles {         code         __typename       }       id       studentRoles {         id         school {           id           shortName           organizationType           __typename         }         status         __typename       }       userSchoolPermissions {         schoolId         permissions         __typename       }       systemAdminRole {         id         __typename       }       businessAdminRolesV2 {         id         school {           id           organizationType           __typename         }         orgUnitId         __typename       }       __typename     }     getCurrentUserSchoolRoles {       schoolId       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_UserRoleLoaderGetRoles](ctx, request)
}
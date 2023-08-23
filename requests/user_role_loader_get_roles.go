package requests

import "s21client/gql"

type Variables_UserRoleLoaderGetRoles struct {
}


type Data_UserRoleLoaderGetRoles struct {
	User_UserRoleLoaderGetRoles User_UserRoleLoaderGetRoles `json:"user"`
}

type User_UserRoleLoaderGetRoles struct {
	GetCurrentUser_UserRoleLoaderGetRoles            GetCurrentUser_UserRoleLoaderGetRoles             `json:"getCurrentUser"`
	GetCurrentUserSchoolRoles []GetCurrentUserSchoolRole_UserRoleLoaderGetRoles `json:"getCurrentUserSchoolRoles"`
	Typename                  string                     `json:"__typename"`
}

type GetCurrentUser_UserRoleLoaderGetRoles struct {
	FunctionalRoles       []FunctionalRole_UserRoleLoaderGetRoles       `json:"functionalRoles"`
	ID                    string                 `json:"id"`
	StudentRoles          []StudentRole_UserRoleLoaderGetRoles          `json:"studentRoles"`
	UserSchoolPermissions []UserSchoolPermission_UserRoleLoaderGetRoles `json:"userSchoolPermissions"`
	SystemAdminRole       interface{}            `json:"systemAdminRole"`
	BusinessAdminRolesV2  []interface{}          `json:"businessAdminRolesV2"`
	Typename              string                 `json:"__typename"`
}

type FunctionalRole_UserRoleLoaderGetRoles struct {
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

type StudentRole_UserRoleLoaderGetRoles struct {
	ID       string `json:"id"`
	School_UserRoleLoaderGetRoles   School_UserRoleLoaderGetRoles `json:"school"`
	Status   string `json:"status"`
	Typename string `json:"__typename"`
}

type School_UserRoleLoaderGetRoles struct {
	ID               string `json:"id"`
	ShortName        string `json:"shortName"`
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type UserSchoolPermission_UserRoleLoaderGetRoles struct {
	SchoolID    string   `json:"schoolId"`
	Permissions []string `json:"permissions"`
	Typename    string   `json:"__typename"`
}

type GetCurrentUserSchoolRole_UserRoleLoaderGetRoles struct {
	SchoolID string `json:"schoolId"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) UserRoleLoaderGetRoles(variables Variables_UserRoleLoaderGetRoles) (Data_UserRoleLoaderGetRoles, error) {
	request := gql.NewQueryRequest[Variables_UserRoleLoaderGetRoles](
		"query userRoleLoaderGetRoles {   user {     getCurrentUser {       functionalRoles {         code         __typename       }       id       studentRoles {         id         school {           id           shortName           organizationType           __typename         }         status         __typename       }       userSchoolPermissions {         schoolId         permissions         __typename       }       systemAdminRole {         id         __typename       }       businessAdminRolesV2 {         id         school {           id           organizationType           __typename         }         orgUnitId         __typename       }       __typename     }     getCurrentUserSchoolRoles {       schoolId       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_UserRoleLoaderGetRoles](ctx, request)
}
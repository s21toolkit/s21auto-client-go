package requests

import "github.com/s21toolkit/s21client/gql"

type UserRoleLoaderGetRoles_Variables struct {
}


type UserRoleLoaderGetRoles_Data struct {
	User UserRoleLoaderGetRoles_Data_User `json:"user"`
}

type UserRoleLoaderGetRoles_Data_User struct {
	GetCurrentUser            UserRoleLoaderGetRoles_Data_GetCurrentUser             `json:"getCurrentUser"`
	GetCurrentUserSchoolRoles []UserRoleLoaderGetRoles_Data_GetCurrentUserSchoolRole `json:"getCurrentUserSchoolRoles"`
	Typename                  string                     `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_GetCurrentUser struct {
	FunctionalRoles       []UserRoleLoaderGetRoles_Data_FunctionalRole       `json:"functionalRoles"`
	ID                    string                 `json:"id"`
	StudentRoles          []UserRoleLoaderGetRoles_Data_StudentRole          `json:"studentRoles"`
	UserSchoolPermissions []UserRoleLoaderGetRoles_Data_UserSchoolPermission `json:"userSchoolPermissions"`
	SystemAdminRole       interface{}            `json:"systemAdminRole"`
	BusinessAdminRolesV2  []interface{}          `json:"businessAdminRolesV2"`
	Typename              string                 `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_FunctionalRole struct {
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_StudentRole struct {
	ID       string `json:"id"`
	School   UserRoleLoaderGetRoles_Data_School `json:"school"`
	Status   string `json:"status"`
	Typename string `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_School struct {
	ID               string `json:"id"`
	ShortName        string `json:"shortName"`
	OrganizationType string `json:"organizationType"`
	Typename         string `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_UserSchoolPermission struct {
	SchoolID    string   `json:"schoolId"`
	Permissions []string `json:"permissions"`
	Typename    string   `json:"__typename"`
}

type UserRoleLoaderGetRoles_Data_GetCurrentUserSchoolRole struct {
	SchoolID string `json:"schoolId"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) UserRoleLoaderGetRoles(variables UserRoleLoaderGetRoles_Variables) (UserRoleLoaderGetRoles_Data, error) {
	request := gql.NewQueryRequest[UserRoleLoaderGetRoles_Variables](
		"query userRoleLoaderGetRoles {\n  user {\n    getCurrentUser {\n      functionalRoles {\n        code\n        __typename\n      }\n      id\n      studentRoles {\n        id\n        school {\n          id\n          shortName\n          organizationType\n          __typename\n        }\n        status\n        __typename\n      }\n      userSchoolPermissions {\n        schoolId\n        permissions\n        __typename\n      }\n      systemAdminRole {\n        id\n        __typename\n      }\n      businessAdminRolesV2 {\n        id\n        school {\n          id\n          organizationType\n          __typename\n        }\n        orgUnitId\n        __typename\n      }\n      __typename\n    }\n    getCurrentUserSchoolRoles {\n      schoolId\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[UserRoleLoaderGetRoles_Data](ctx, request)
}
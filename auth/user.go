package auth

import (
	"context"
	"s21client/gql"

	"github.com/go-resty/resty/v2"
)

type userDataRoleResponse struct {
	Data userDataRole `json:"data"`
}

type userDataRole struct {
	User User `json:"user"`
}

type User struct {
	Data  UserData   `json:"getCurrentUser"`
	Roles []UserRole `json:"getCurrentUserSchoolRoles"`
}

type UserData struct {
	FunctionalRoles    []FunctionalRole       `json:"functionalRoles"`
	ID                 string                 `json:"id"`
	StudentRoles       []StudentRole          `json:"studentRoles"`
	SchoolPermissions  []UserSchoolPermission `json:"userSchoolPermissions"`
	SystemAdminRole    interface{}            `json:"systemAdminRole"`
	BusinessAdminRoles []interface{}          `json:"businessAdminRolesV2"`
}

type FunctionalRole struct {
	Code string `json:"code"`
}

type StudentRole struct {
	ID     string `json:"id"`
	School School `json:"school"`
	Status string `json:"status"`
}

type School struct {
	ID               string `json:"id"`
	ShortName        string `json:"shortName"`
	OrganizationType string `json:"organizationType"`
}

type UserSchoolPermission struct {
	SchoolID    string   `json:"schoolId"`
	Permissions []string `json:"permissions"`
}

type UserRole struct {
	SchoolID string `json:"schoolId"`
}

var userRoleLoaderGetRolesRequest = gql.NewQueryRequest(
	"query userRoleLoaderGetRoles { user { getCurrentUser { functionalRoles { code __typename } id studentRoles { id school { id shortName organizationType __typename } status __typename } userSchoolPermissions { schoolId permissions __typename } systemAdminRole { id __typename } businessAdminRolesV2 { id school { id organizationType __typename } orgUnitId __typename } __typename } getCurrentUserSchoolRoles { schoolId __typename } __typename } } ",
	gql.None{},
)

var s21ApiUrl = "https://edu.21-school.ru/services/graphql"

func RequestUserData(token Token, ctx context.Context) (user User, err error) {
	client := resty.New()

	res, err := client.R().SetContext(ctx).SetAuthToken(token.AccessToken).SetBody(userRoleLoaderGetRolesRequest).Post(s21ApiUrl)

	if err != nil {
		return
	}

	userRoleData, err := gql.ExtractResponseData[userDataRole](res)

	if err != nil {
		return
	}

	user = userRoleData.User

	return
}

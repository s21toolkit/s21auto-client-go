package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUserRestrictionsInfo struct {
}


type Data_GetUserRestrictionsInfo struct {
	Data_School21_GetUserRestrictionsInfo Data_School21_GetUserRestrictionsInfo `json:"school21"`
}

type Data_School21_GetUserRestrictionsInfo struct {
	GetUserRestrictions []interface{} `json:"getUserRestrictions"`
	Typename            string        `json:"__typename"`
}


func (ctx *RequestContext) GetUserRestrictionsInfo(variables Variables_GetUserRestrictionsInfo) (Data_GetUserRestrictionsInfo, error) {
	request := gql.NewQueryRequest[Variables_GetUserRestrictionsInfo](
		"query getUserRestrictionsInfo {\n  school21 {\n    getUserRestrictions {\n      restrictionId\n      restrictionType\n      userId\n      schoolId\n      isActive\n      createdTs\n      updatedTs\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUserRestrictionsInfo](ctx, request)
}
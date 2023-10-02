package requests

import "github.com/s21toolkit/s21client/gql"

type GetUserRestrictionsInfo_Variables struct {
}


type GetUserRestrictionsInfo_Data struct {
	School21 GetUserRestrictionsInfo_Data_School21 `json:"school21"`
}

type GetUserRestrictionsInfo_Data_School21 struct {
	GetUserRestrictions []interface{} `json:"getUserRestrictions"`
	Typename            string        `json:"__typename"`
}


func (ctx *RequestContext) GetUserRestrictionsInfo(variables GetUserRestrictionsInfo_Variables) (GetUserRestrictionsInfo_Data, error) {
	request := gql.NewQueryRequest[GetUserRestrictionsInfo_Variables](
		"query getUserRestrictionsInfo {\n  school21 {\n    getUserRestrictions {\n      restrictionId\n      restrictionType\n      userId\n      schoolId\n      isActive\n      createdTs\n      updatedTs\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetUserRestrictionsInfo_Data](ctx, request)
}
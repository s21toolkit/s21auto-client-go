package requests

import "s21client/gql"

type Variables_GetUserRestrictionsInfo struct {
}


type Data_GetUserRestrictionsInfo struct {
	School21_GetUserRestrictionsInfo School21_GetUserRestrictionsInfo `json:"school21"`
}

type School21_GetUserRestrictionsInfo struct {
	GetUserRestrictions []interface{} `json:"getUserRestrictions"`
	Typename            string        `json:"__typename"`
}

func (ctx *RequestContext) GetUserRestrictionsInfo(variables Variables_GetUserRestrictionsInfo) (Data_GetUserRestrictionsInfo, error) {
	request := gql.NewQueryRequest[Variables_GetUserRestrictionsInfo](
		"query getUserRestrictionsInfo {   school21 {     getUserRestrictions {       restrictionId       restrictionType       userId       schoolId       isActive       createdTs       updatedTs       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetUserRestrictionsInfo](ctx, request)
}
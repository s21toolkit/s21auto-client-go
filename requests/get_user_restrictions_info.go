package requests

import "s21client/gql"

type Request_Variables_GetUserRestrictionsInfo struct {
}


type Response_Data_GetUserRestrictionsInfo struct {
	Response_School21_GetUserRestrictionsInfo Response_School21_GetUserRestrictionsInfo `json:"school21"`
}

type Response_School21_GetUserRestrictionsInfo struct {
	GetUserRestrictions []interface{} `json:"getUserRestrictions"`
	Typename            string        `json:"__typename"`
}

func (ctx *RequestContext) GetUserRestrictionsInfo(variables Request_Variables_GetUserRestrictionsInfo) (Response_Data_GetUserRestrictionsInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_GetUserRestrictionsInfo](
		"query getUserRestrictionsInfo {   school21 {     getUserRestrictions {       restrictionId       restrictionType       userId       schoolId       isActive       createdTs       updatedTs       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetUserRestrictionsInfo](ctx, request)
}
package requests

import "s21client/gql"

type Variables_GetUserFeatureFlags struct {
	EntityIDList []string `json:"entityIdList"`
}


type Data_GetUserFeatureFlags struct {
	User_GetUserFeatureFlags User_GetUserFeatureFlags `json:"user"`
}

type User_GetUserFeatureFlags struct {
	GetAllBackendConfigurations []GetAllBackendConfiguration_GetUserFeatureFlags `json:"getAllBackendConfigurations"`
	Typename                    string                       `json:"__typename"`
}

type GetAllBackendConfiguration_GetUserFeatureFlags struct {
	PropertyCode string `json:"propertyCode"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) GetUserFeatureFlags(variables Variables_GetUserFeatureFlags) (Data_GetUserFeatureFlags, error) {
	request := gql.NewQueryRequest[Variables_GetUserFeatureFlags](
		"query getUserFeatureFlags($entityIdList: [ID!]!) {   user {     getAllBackendConfigurations(entityIdList: $entityIdList) {       propertyCode       value       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetUserFeatureFlags](ctx, request)
}
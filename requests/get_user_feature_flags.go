package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUserFeatureFlags struct {
	EntityID     *string  `json:"entityId,omitempty"`
	EntityIDList []string `json:"entityIdList,omitempty"`
}


type Data_GetUserFeatureFlags struct {
	Data_User_GetUserFeatureFlags Data_User_GetUserFeatureFlags `json:"user"`
}

type Data_User_GetUserFeatureFlags struct {
	GetAllBackendConfigurations []Data_GetAllBackendConfiguration_GetUserFeatureFlags `json:"getAllBackendConfigurations"`
	Typename                    string                       `json:"__typename"`
}

type Data_GetAllBackendConfiguration_GetUserFeatureFlags struct {
	PropertyCode string `json:"propertyCode"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) GetUserFeatureFlags(variables Variables_GetUserFeatureFlags) (Data_GetUserFeatureFlags, error) {
	request := gql.NewQueryRequest[Variables_GetUserFeatureFlags](
		"query getUserFeatureFlags($entityId: String!) {\n  user {\n    getAllBackendConfigurations: getAllBackendConfigurationsV2(entityId: $entityId) {\n      propertyCode\n      value\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUserFeatureFlags](ctx, request)
}
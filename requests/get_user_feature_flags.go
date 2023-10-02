package requests

import "github.com/s21toolkit/s21client/gql"

type GetUserFeatureFlags_Variables struct {
	EntityID     *string  `json:"entityId,omitempty"`
	EntityIDList []string `json:"entityIdList,omitempty"`
}


type GetUserFeatureFlags_Data struct {
	User GetUserFeatureFlags_Data_User `json:"user"`
}

type GetUserFeatureFlags_Data_User struct {
	GetAllBackendConfigurations []GetUserFeatureFlags_Data_GetAllBackendConfiguration `json:"getAllBackendConfigurations"`
	Typename                    string                       `json:"__typename"`
}

type GetUserFeatureFlags_Data_GetAllBackendConfiguration struct {
	PropertyCode string `json:"propertyCode"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) GetUserFeatureFlags(variables GetUserFeatureFlags_Variables) (GetUserFeatureFlags_Data, error) {
	request := gql.NewQueryRequest[GetUserFeatureFlags_Variables](
		"query getUserFeatureFlags($entityId: String!) {\n  user {\n    getAllBackendConfigurations: getAllBackendConfigurationsV2(entityId: $entityId) {\n      propertyCode\n      value\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetUserFeatureFlags_Data](ctx, request)
}
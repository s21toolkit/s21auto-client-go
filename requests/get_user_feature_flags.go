package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetUserFeatureFlags struct {
	EntityIDList []string `json:"entityIdList,omitempty"`
	EntityID     *string  `json:"entityId,omitempty"`
}


type Response_Data_GetUserFeatureFlags struct {
	Response_User_GetUserFeatureFlags Response_User_GetUserFeatureFlags `json:"user"`
}

type Response_User_GetUserFeatureFlags struct {
	GetAllBackendConfigurations []Response_GetAllBackendConfiguration_GetUserFeatureFlags `json:"getAllBackendConfigurations"`
	Typename                    string                       `json:"__typename"`
}

type Response_GetAllBackendConfiguration_GetUserFeatureFlags struct {
	PropertyCode string `json:"propertyCode"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) GetUserFeatureFlags(variables Request_Variables_GetUserFeatureFlags) (Response_Data_GetUserFeatureFlags, error) {
	request := gql.NewQueryRequest[Request_Variables_GetUserFeatureFlags](
		"query getUserFeatureFlags($entityIdList: [ID!]!) {\n  user {\n    getAllBackendConfigurations(entityIdList: $entityIdList) {\n      propertyCode\n      value\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetUserFeatureFlags](ctx, request)
}
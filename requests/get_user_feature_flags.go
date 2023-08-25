package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetUserFeatureFlags struct {
	EntityIDList []string `json:"entityIdList"`
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
		"query getUserFeatureFlags($entityIdList: [ID!]!) {   user {     getAllBackendConfigurations(entityIdList: $entityIdList) {       propertyCode       value       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetUserFeatureFlags](ctx, request)
}
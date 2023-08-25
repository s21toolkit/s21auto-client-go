package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetSaleProgressPercentages struct {
}


type Response_Data_GetSaleProgressPercentages struct {
	Response_School21_GetSaleProgressPercentages Response_School21_GetSaleProgressPercentages `json:"school21"`
}

type Response_School21_GetSaleProgressPercentages struct {
	GetSaleProgressPercentages []Response_GetSaleProgressPercentage_GetSaleProgressPercentages `json:"getSaleProgressPercentages"`
	Typename                   string                      `json:"__typename"`
}

type Response_GetSaleProgressPercentage_GetSaleProgressPercentages struct {
	RpType             string `json:"rpType"`
	ProgressPercentage int64  `json:"progressPercentage"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetSaleProgressPercentages(variables Request_Variables_GetSaleProgressPercentages) (Response_Data_GetSaleProgressPercentages, error) {
	request := gql.NewQueryRequest[Request_Variables_GetSaleProgressPercentages](
		"query getSaleProgressPercentages {\n  school21 {\n    getSaleProgressPercentages {\n      ...RpSaleInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment RpSaleInfo on RpSaleProgress {\n  rpType\n  progressPercentage\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetSaleProgressPercentages](ctx, request)
}
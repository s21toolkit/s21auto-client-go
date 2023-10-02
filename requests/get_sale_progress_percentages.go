package requests

import "github.com/s21toolkit/s21client/gql"

type GetSaleProgressPercentages_Variables struct {
}


type GetSaleProgressPercentages_Data struct {
	School21 GetSaleProgressPercentages_Data_School21 `json:"school21"`
}

type GetSaleProgressPercentages_Data_School21 struct {
	GetSaleProgressPercentages []GetSaleProgressPercentages_Data_GetSaleProgressPercentage `json:"getSaleProgressPercentages"`
	Typename                   string                      `json:"__typename"`
}

type GetSaleProgressPercentages_Data_GetSaleProgressPercentage struct {
	RpType             string `json:"rpType"`
	ProgressPercentage int64  `json:"progressPercentage"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetSaleProgressPercentages(variables GetSaleProgressPercentages_Variables) (GetSaleProgressPercentages_Data, error) {
	request := gql.NewQueryRequest[GetSaleProgressPercentages_Variables](
		"query getSaleProgressPercentages {\n  school21 {\n    getSaleProgressPercentages {\n      ...RpSaleInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment RpSaleInfo on RpSaleProgress {\n  rpType\n  progressPercentage\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetSaleProgressPercentages_Data](ctx, request)
}
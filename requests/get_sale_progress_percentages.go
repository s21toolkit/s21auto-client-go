package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetSaleProgressPercentages struct {
}


type Data_GetSaleProgressPercentages struct {
	Data_School21_GetSaleProgressPercentages Data_School21_GetSaleProgressPercentages `json:"school21"`
}

type Data_School21_GetSaleProgressPercentages struct {
	GetSaleProgressPercentages []Data_GetSaleProgressPercentage_GetSaleProgressPercentages `json:"getSaleProgressPercentages"`
	Typename                   string                      `json:"__typename"`
}

type Data_GetSaleProgressPercentage_GetSaleProgressPercentages struct {
	RpType             string `json:"rpType"`
	ProgressPercentage int64  `json:"progressPercentage"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetSaleProgressPercentages(variables Variables_GetSaleProgressPercentages) (Data_GetSaleProgressPercentages, error) {
	request := gql.NewQueryRequest[Variables_GetSaleProgressPercentages](
		"query getSaleProgressPercentages {\n  school21 {\n    getSaleProgressPercentages {\n      ...RpSaleInfo\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment RpSaleInfo on RpSaleProgress {\n  rpType\n  progressPercentage\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetSaleProgressPercentages](ctx, request)
}
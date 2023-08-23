package requests

import "s21client/gql"

type Variables_GetSaleProgressPercentages struct {
}


type Data_GetSaleProgressPercentages struct {
	School21_GetSaleProgressPercentages School21_GetSaleProgressPercentages `json:"school21"`
}

type School21_GetSaleProgressPercentages struct {
	GetSaleProgressPercentages []GetSaleProgressPercentage_GetSaleProgressPercentages `json:"getSaleProgressPercentages"`
	Typename                   string                      `json:"__typename"`
}

type GetSaleProgressPercentage_GetSaleProgressPercentages struct {
	RpType             string `json:"rpType"`
	ProgressPercentage int64  `json:"progressPercentage"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetSaleProgressPercentages(variables Variables_GetSaleProgressPercentages) (Data_GetSaleProgressPercentages, error) {
	request := gql.NewQueryRequest[Variables_GetSaleProgressPercentages](
		"query getSaleProgressPercentages {   school21 {     getSaleProgressPercentages {       ...RpSaleInfo       __typename     }     __typename   } }  fragment RpSaleInfo on RpSaleProgress {   rpType   progressPercentage   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetSaleProgressPercentages](ctx, request)
}
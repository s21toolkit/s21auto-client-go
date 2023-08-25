package requests

import "s21client/gql"

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
		"query getSaleProgressPercentages {   school21 {     getSaleProgressPercentages {       ...RpSaleInfo       __typename     }     __typename   } }  fragment RpSaleInfo on RpSaleProgress {   rpType   progressPercentage   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetSaleProgressPercentages](ctx, request)
}
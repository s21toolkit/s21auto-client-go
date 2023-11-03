package requests

import "github.com/s21toolkit/s21client/gql"

type GetIsHonorRatingNeeded_Variables struct {
}


type GetIsHonorRatingNeeded_Data struct {
	HonorRating GetIsHonorRatingNeeded_Data_HonorRating `json:"honorRating"`
}

type GetIsHonorRatingNeeded_Data_HonorRating struct {
	IsHonorRatingNeeded bool   `json:"isHonorRatingNeeded"`
	Typename            string `json:"__typename"`
}


func (ctx *RequestContext) GetIsHonorRatingNeeded(variables GetIsHonorRatingNeeded_Variables) (GetIsHonorRatingNeeded_Data, error) {
	request := gql.NewQueryRequest[GetIsHonorRatingNeeded_Variables](
		"query getIsHonorRatingNeeded {\n  honorRating {\n    isHonorRatingNeeded\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetIsHonorRatingNeeded_Data](ctx, request)
}
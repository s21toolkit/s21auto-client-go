package requests

import "s21client/gql"

type Variables_GetIsHonorRatingNeeded struct {
}


type Data_GetIsHonorRatingNeeded struct {
	Student_GetIsHonorRatingNeeded Student_GetIsHonorRatingNeeded `json:"student"`
}

type Student_GetIsHonorRatingNeeded struct {
	IsHonorRatingNeeded bool   `json:"isHonorRatingNeeded"`
	Typename            string `json:"__typename"`
}

func (ctx *RequestContext) GetIsHonorRatingNeeded(variables Variables_GetIsHonorRatingNeeded) (Data_GetIsHonorRatingNeeded, error) {
	request := gql.NewQueryRequest[Variables_GetIsHonorRatingNeeded](
		"query getIsHonorRatingNeeded {   student {     isHonorRatingNeeded     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetIsHonorRatingNeeded](ctx, request)
}
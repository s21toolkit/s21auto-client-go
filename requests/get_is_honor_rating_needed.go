package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetIsHonorRatingNeeded struct {
}


type Data_GetIsHonorRatingNeeded struct {
	Data_Student_GetIsHonorRatingNeeded Data_Student_GetIsHonorRatingNeeded `json:"student"`
}

type Data_Student_GetIsHonorRatingNeeded struct {
	IsHonorRatingNeeded bool   `json:"isHonorRatingNeeded"`
	Typename            string `json:"__typename"`
}


func (ctx *RequestContext) GetIsHonorRatingNeeded(variables Variables_GetIsHonorRatingNeeded) (Data_GetIsHonorRatingNeeded, error) {
	request := gql.NewQueryRequest[Variables_GetIsHonorRatingNeeded](
		"query getIsHonorRatingNeeded {\n  student {\n    isHonorRatingNeeded\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetIsHonorRatingNeeded](ctx, request)
}
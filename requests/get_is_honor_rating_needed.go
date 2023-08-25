package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetIsHonorRatingNeeded struct {
}


type Response_Data_GetIsHonorRatingNeeded struct {
	Response_Student_GetIsHonorRatingNeeded Response_Student_GetIsHonorRatingNeeded `json:"student"`
}

type Response_Student_GetIsHonorRatingNeeded struct {
	IsHonorRatingNeeded bool   `json:"isHonorRatingNeeded"`
	Typename            string `json:"__typename"`
}

func (ctx *RequestContext) GetIsHonorRatingNeeded(variables Request_Variables_GetIsHonorRatingNeeded) (Response_Data_GetIsHonorRatingNeeded, error) {
	request := gql.NewQueryRequest[Request_Variables_GetIsHonorRatingNeeded](
		"query getIsHonorRatingNeeded {\n  student {\n    isHonorRatingNeeded\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetIsHonorRatingNeeded](ctx, request)
}
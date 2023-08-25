package requests

import "s21client/gql"

type Variables_GetSecondRoundCodeReviewProjects struct {
	Paging_GetSecondRoundCodeReviewProjects Paging_GetSecondRoundCodeReviewProjects `json:"paging"`
}

type Paging_GetSecondRoundCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetSecondRoundCodeReviewProjects struct {
	Student_GetSecondRoundCodeReviewProjects Student_GetSecondRoundCodeReviewProjects `json:"student"`
}

type Student_GetSecondRoundCodeReviewProjects struct {
	GetSecondRoundCodeReviewProjects []interface{} `json:"getSecondRoundCodeReviewProjects"`
	Typename                         string        `json:"__typename"`
}

func (ctx *RequestContext) GetSecondRoundCodeReviewProjects(variables Variables_GetSecondRoundCodeReviewProjects) (Data_GetSecondRoundCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetSecondRoundCodeReviewProjects](
		"query getSecondRoundCodeReviewProjects($paging: PagingInput!) {   student {     getSecondRoundCodeReviewProjects(paging: $paging) {       ...CodeReviewProject       __typename     }     __typename   } }  fragment CodeReviewProject on CodeReview {   goalId   goalTitle   studentGoalId   studentCodeReviewStatus   goalExecutionType   studentTaskAdditionalAttributesModel {     codeReviewCost     codeReviewDuration     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetSecondRoundCodeReviewProjects](ctx, request)
}
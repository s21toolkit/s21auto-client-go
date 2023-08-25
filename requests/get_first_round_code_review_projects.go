package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetFirstRoundCodeReviewProjects struct {
	Paging_GetFirstRoundCodeReviewProjects Paging_GetFirstRoundCodeReviewProjects `json:"paging"`
}

type Paging_GetFirstRoundCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type Data_GetFirstRoundCodeReviewProjects struct {
	Student_GetFirstRoundCodeReviewProjects Student_GetFirstRoundCodeReviewProjects `json:"student"`
}

type Student_GetFirstRoundCodeReviewProjects struct {
	GetFirstRoundCodeReviewProjects []interface{} `json:"getFirstRoundCodeReviewProjects"`
	Typename                        string        `json:"__typename"`
}

func (ctx *RequestContext) GetFirstRoundCodeReviewProjects(variables Variables_GetFirstRoundCodeReviewProjects) (Data_GetFirstRoundCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetFirstRoundCodeReviewProjects](
		"query getFirstRoundCodeReviewProjects($paging: PagingInput!) {   student {     getFirstRoundCodeReviewProjects(paging: $paging) {       ...CodeReviewProject       __typename     }     __typename   } }  fragment CodeReviewProject on CodeReview {   goalId   goalTitle   studentGoalId   studentCodeReviewStatus   goalExecutionType   studentTaskAdditionalAttributesModel {     codeReviewCost     codeReviewDuration     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetFirstRoundCodeReviewProjects](ctx, request)
}

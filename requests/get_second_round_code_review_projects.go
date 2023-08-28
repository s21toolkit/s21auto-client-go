package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetSecondRoundCodeReviewProjects struct {
	Variables_Paging_GetSecondRoundCodeReviewProjects Variables_Paging_GetSecondRoundCodeReviewProjects `json:"paging"`
}

type Variables_Paging_GetSecondRoundCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetSecondRoundCodeReviewProjects struct {
	Data_Student_GetSecondRoundCodeReviewProjects Data_Student_GetSecondRoundCodeReviewProjects `json:"student"`
}

type Data_Student_GetSecondRoundCodeReviewProjects struct {
	GetSecondRoundCodeReviewProjects []interface{} `json:"getSecondRoundCodeReviewProjects"`
	Typename                         string        `json:"__typename"`
}


func (ctx *RequestContext) GetSecondRoundCodeReviewProjects(variables Variables_GetSecondRoundCodeReviewProjects) (Data_GetSecondRoundCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetSecondRoundCodeReviewProjects](
		"query getSecondRoundCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getSecondRoundCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetSecondRoundCodeReviewProjects](ctx, request)
}
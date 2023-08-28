package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetFirstRoundCodeReviewProjects struct {
	Variables_Paging_GetFirstRoundCodeReviewProjects Variables_Paging_GetFirstRoundCodeReviewProjects `json:"paging"`
}

type Variables_Paging_GetFirstRoundCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetFirstRoundCodeReviewProjects struct {
	Data_Student_GetFirstRoundCodeReviewProjects Data_Student_GetFirstRoundCodeReviewProjects `json:"student"`
}

type Data_Student_GetFirstRoundCodeReviewProjects struct {
	GetFirstRoundCodeReviewProjects []interface{} `json:"getFirstRoundCodeReviewProjects"`
	Typename                        string        `json:"__typename"`
}


func (ctx *RequestContext) GetFirstRoundCodeReviewProjects(variables Variables_GetFirstRoundCodeReviewProjects) (Data_GetFirstRoundCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetFirstRoundCodeReviewProjects](
		"query getFirstRoundCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getFirstRoundCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetFirstRoundCodeReviewProjects](ctx, request)
}
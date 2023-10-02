package requests

import "github.com/s21toolkit/s21client/gql"

type GetFirstRoundCodeReviewProjects_Variables struct {
	Paging GetFirstRoundCodeReviewProjects_Variables_Paging `json:"paging"`
}

type GetFirstRoundCodeReviewProjects_Variables_Paging struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetFirstRoundCodeReviewProjects_Data struct {
	Student GetFirstRoundCodeReviewProjects_Data_Student `json:"student"`
}

type GetFirstRoundCodeReviewProjects_Data_Student struct {
	GetFirstRoundCodeReviewProjects []interface{} `json:"getFirstRoundCodeReviewProjects"`
	Typename                        string        `json:"__typename"`
}


func (ctx *RequestContext) GetFirstRoundCodeReviewProjects(variables GetFirstRoundCodeReviewProjects_Variables) (GetFirstRoundCodeReviewProjects_Data, error) {
	request := gql.NewQueryRequest[GetFirstRoundCodeReviewProjects_Variables](
		"query getFirstRoundCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getFirstRoundCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetFirstRoundCodeReviewProjects_Data](ctx, request)
}
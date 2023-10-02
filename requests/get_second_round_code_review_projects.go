package requests

import "github.com/s21toolkit/s21client/gql"

type GetSecondRoundCodeReviewProjects_Variables struct {
	Paging GetSecondRoundCodeReviewProjects_Variables_Paging `json:"paging"`
}

type GetSecondRoundCodeReviewProjects_Variables_Paging struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetSecondRoundCodeReviewProjects_Data struct {
	Student GetSecondRoundCodeReviewProjects_Data_Student `json:"student"`
}

type GetSecondRoundCodeReviewProjects_Data_Student struct {
	GetSecondRoundCodeReviewProjects []interface{} `json:"getSecondRoundCodeReviewProjects"`
	Typename                         string        `json:"__typename"`
}


func (ctx *RequestContext) GetSecondRoundCodeReviewProjects(variables GetSecondRoundCodeReviewProjects_Variables) (GetSecondRoundCodeReviewProjects_Data, error) {
	request := gql.NewQueryRequest[GetSecondRoundCodeReviewProjects_Variables](
		"query getSecondRoundCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getSecondRoundCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetSecondRoundCodeReviewProjects_Data](ctx, request)
}
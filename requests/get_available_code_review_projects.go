package requests

import "github.com/s21toolkit/s21client/gql"

type GetAvailableCodeReviewProjects_Variables struct {
	Paging GetAvailableCodeReviewProjects_Variables_Paging `json:"paging"`
}

type GetAvailableCodeReviewProjects_Variables_Paging struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetAvailableCodeReviewProjects_Data struct {
	Student GetAvailableCodeReviewProjects_Data_Student `json:"student"`
}

type GetAvailableCodeReviewProjects_Data_Student struct {
	GetAvailableCodeReviewProjects []GetAvailableCodeReviewProjects_Data_GetAvailableCodeReviewProject `json:"getAvailableCodeReviewProjects"`
	Typename                       string                          `json:"__typename"`
}

type GetAvailableCodeReviewProjects_Data_GetAvailableCodeReviewProject struct {
	GoalID                               string                               `json:"goalId"`
	GoalTitle                            string                               `json:"goalTitle"`
	StudentGoalID                        string                               `json:"studentGoalId"`
	StudentCodeReviewStatus              string                               `json:"studentCodeReviewStatus"`
	GoalExecutionType                    string                               `json:"goalExecutionType"`
	StudentTaskAdditionalAttributesModel GetAvailableCodeReviewProjects_Data_StudentTaskAdditionalAttributesModel `json:"studentTaskAdditionalAttributesModel"`
	Typename                             string                               `json:"__typename"`
}

type GetAvailableCodeReviewProjects_Data_StudentTaskAdditionalAttributesModel struct {
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetAvailableCodeReviewProjects(variables GetAvailableCodeReviewProjects_Variables) (GetAvailableCodeReviewProjects_Data, error) {
	request := gql.NewQueryRequest[GetAvailableCodeReviewProjects_Variables](
		"query getAvailableCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getAvailableCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetAvailableCodeReviewProjects_Data](ctx, request)
}
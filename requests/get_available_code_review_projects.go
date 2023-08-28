package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetAvailableCodeReviewProjects struct {
	Variables_Paging_GetAvailableCodeReviewProjects Variables_Paging_GetAvailableCodeReviewProjects `json:"paging"`
}

type Variables_Paging_GetAvailableCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetAvailableCodeReviewProjects struct {
	Data_Student_GetAvailableCodeReviewProjects Data_Student_GetAvailableCodeReviewProjects `json:"student"`
}

type Data_Student_GetAvailableCodeReviewProjects struct {
	GetAvailableCodeReviewProjects []Data_GetAvailableCodeReviewProject_GetAvailableCodeReviewProjects `json:"getAvailableCodeReviewProjects"`
	Typename                       string                          `json:"__typename"`
}

type Data_GetAvailableCodeReviewProject_GetAvailableCodeReviewProjects struct {
	GoalID                               string                               `json:"goalId"`
	GoalTitle                            string                               `json:"goalTitle"`
	StudentGoalID                        string                               `json:"studentGoalId"`
	StudentCodeReviewStatus              string                               `json:"studentCodeReviewStatus"`
	GoalExecutionType                    string                               `json:"goalExecutionType"`
	Data_StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects Data_StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects `json:"studentTaskAdditionalAttributesModel"`
	Typename                             string                               `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects struct {
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetAvailableCodeReviewProjects(variables Variables_GetAvailableCodeReviewProjects) (Data_GetAvailableCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetAvailableCodeReviewProjects](
		"query getAvailableCodeReviewProjects($paging: PagingInput!) {\n  student {\n    getAvailableCodeReviewProjects(paging: $paging) {\n      ...CodeReviewProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CodeReviewProject on CodeReview {\n  goalId\n  goalTitle\n  studentGoalId\n  studentCodeReviewStatus\n  goalExecutionType\n  studentTaskAdditionalAttributesModel {\n    codeReviewCost\n    codeReviewDuration\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetAvailableCodeReviewProjects](ctx, request)
}
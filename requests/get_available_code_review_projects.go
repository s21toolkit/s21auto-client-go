package requests

import "s21client/gql"

type Variables_GetAvailableCodeReviewProjects struct {
	Paging_GetAvailableCodeReviewProjects Paging_GetAvailableCodeReviewProjects `json:"paging"`
}

type Paging_GetAvailableCodeReviewProjects struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetAvailableCodeReviewProjects struct {
	Student_GetAvailableCodeReviewProjects Student_GetAvailableCodeReviewProjects `json:"student"`
}

type Student_GetAvailableCodeReviewProjects struct {
	GetAvailableCodeReviewProjects []GetAvailableCodeReviewProject_GetAvailableCodeReviewProjects `json:"getAvailableCodeReviewProjects"`
	Typename                       string                          `json:"__typename"`
}

type GetAvailableCodeReviewProject_GetAvailableCodeReviewProjects struct {
	GoalID                               string                               `json:"goalId"`
	GoalTitle                            string                               `json:"goalTitle"`
	StudentGoalID                        string                               `json:"studentGoalId"`
	StudentCodeReviewStatus              string                               `json:"studentCodeReviewStatus"`
	GoalExecutionType                    string                               `json:"goalExecutionType"`
	StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects `json:"studentTaskAdditionalAttributesModel"`
	Typename                             string                               `json:"__typename"`
}

type StudentTaskAdditionalAttributesModel_GetAvailableCodeReviewProjects struct {
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CodeReviewDuration int64  `json:"codeReviewDuration"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetAvailableCodeReviewProjects(variables Variables_GetAvailableCodeReviewProjects) (Data_GetAvailableCodeReviewProjects, error) {
	request := gql.NewQueryRequest[Variables_GetAvailableCodeReviewProjects](
		"query getAvailableCodeReviewProjects($paging: PagingInput!) {   student {     getAvailableCodeReviewProjects(paging: $paging) {       ...CodeReviewProject       __typename     }     __typename   } }  fragment CodeReviewProject on CodeReview {   goalId   goalTitle   studentGoalId   studentCodeReviewStatus   goalExecutionType   studentTaskAdditionalAttributesModel {     codeReviewCost     codeReviewDuration     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetAvailableCodeReviewProjects](ctx, request)
}
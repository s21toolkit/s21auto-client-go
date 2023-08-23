package requests

import "s21client/gql"

type Variables_PublicProfileGetProjects struct {
	StudentID    string `json:"studentId"`
	StageGroupID string `json:"stageGroupId"`
}


type Data_PublicProfileGetProjects struct {
	School21_PublicProfileGetProjects School21_PublicProfileGetProjects `json:"school21"`
}

type School21_PublicProfileGetProjects struct {
	GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects []GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects `json:"getStudentProjectsForPublicProfileByStageGroup"`
	Typename                                       string                                           `json:"__typename"`
}

type GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects struct {
	GroupName             string      `json:"groupName"`
	Name                  string      `json:"name"`
	Experience            int64       `json:"experience"`
	FinalPercentage       *int64      `json:"finalPercentage"`
	GoalID                int64       `json:"goalId"`
	GoalStatus            string      `json:"goalStatus"`
	AmountAnswers         interface{} `json:"amountAnswers"`
	AmountReviewedAnswers interface{} `json:"amountReviewedAnswers"`
	Typename              string      `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetProjects(variables Variables_PublicProfileGetProjects) (Data_PublicProfileGetProjects, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetProjects](
		"query publicProfileGetProjects($studentId: UUID!, $stageGroupId: ID!) {   school21 {     getStudentProjectsForPublicProfileByStageGroup(       studentId: $studentId       stageGroupId: $stageGroupId     ) {       groupName       name       experience       finalPercentage       goalId       goalStatus       amountAnswers       amountReviewedAnswers       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetProjects](ctx, request)
}
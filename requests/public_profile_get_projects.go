package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetProjects_Variables struct {
	StudentID    string `json:"studentId"`
	StageGroupID string `json:"stageGroupId"`
}


type PublicProfileGetProjects_Data struct {
	School21 PublicProfileGetProjects_Data_School21 `json:"school21"`
}

type PublicProfileGetProjects_Data_School21 struct {
	GetStudentProjectsForPublicProfileByStageGroup []PublicProfileGetProjects_Data_GetStudentProjectsForPublicProfileByStageGroup `json:"getStudentProjectsForPublicProfileByStageGroup"`
	Typename                                       string                                           `json:"__typename"`
}

type PublicProfileGetProjects_Data_GetStudentProjectsForPublicProfileByStageGroup struct {
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


func (ctx *RequestContext) PublicProfileGetProjects(variables PublicProfileGetProjects_Variables) (PublicProfileGetProjects_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetProjects_Variables](
		"query publicProfileGetProjects($studentId: UUID!, $stageGroupId: ID!) {\n  school21 {\n    getStudentProjectsForPublicProfileByStageGroup(\n      studentId: $studentId\n      stageGroupId: $stageGroupId\n    ) {\n      groupName\n      name\n      experience\n      finalPercentage\n      goalId\n      goalStatus\n      amountAnswers\n      amountReviewedAnswers\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetProjects_Data](ctx, request)
}
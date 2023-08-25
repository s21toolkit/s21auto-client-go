package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetProjects struct {
	StudentID    string `json:"studentId"`
	StageGroupID string `json:"stageGroupId"`
}


type Response_Data_PublicProfileGetProjects struct {
	Response_School21_PublicProfileGetProjects Response_School21_PublicProfileGetProjects `json:"school21"`
}

type Response_School21_PublicProfileGetProjects struct {
	Response_GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects []Response_GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects `json:"getStudentProjectsForPublicProfileByStageGroup"`
	Typename                                       string                                           `json:"__typename"`
}

type Response_GetStudentProjectsForPublicProfileByStageGroup_PublicProfileGetProjects struct {
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

func (ctx *RequestContext) PublicProfileGetProjects(variables Request_Variables_PublicProfileGetProjects) (Response_Data_PublicProfileGetProjects, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetProjects](
		"query publicProfileGetProjects($studentId: UUID!, $stageGroupId: ID!) {\n  school21 {\n    getStudentProjectsForPublicProfileByStageGroup(\n      studentId: $studentId\n      stageGroupId: $stageGroupId\n    ) {\n      groupName\n      name\n      experience\n      finalPercentage\n      goalId\n      goalStatus\n      amountAnswers\n      amountReviewedAnswers\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetProjects](ctx, request)
}
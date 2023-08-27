package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetProjectConsistencyInfo struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectConsistencyInfo struct {
	Data_School21_GetProjectConsistencyInfo Data_School21_GetProjectConsistencyInfo `json:"school21"`
}

type Data_School21_GetProjectConsistencyInfo struct {
	Data_LoadGoalConsistencyInfo_GetProjectConsistencyInfo Data_LoadGoalConsistencyInfo_GetProjectConsistencyInfo `json:"loadGoalConsistencyInfo"`
	Typename                string                  `json:"__typename"`
}

type Data_LoadGoalConsistencyInfo_GetProjectConsistencyInfo struct {
	GoalID       string `json:"goalId"`
	IsConsistent bool   `json:"isConsistent"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectConsistencyInfo(variables Variables_GetProjectConsistencyInfo) (Data_GetProjectConsistencyInfo, error) {
	request := gql.NewQueryRequest[Variables_GetProjectConsistencyInfo](
		"query getProjectConsistencyInfo($goalId: ID!) {\n  school21 {\n    loadGoalConsistencyInfo(goalId: $goalId) {\n      goalId\n      isConsistent\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetProjectConsistencyInfo](ctx, request)
}
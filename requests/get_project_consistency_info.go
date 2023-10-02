package requests

import "github.com/s21toolkit/s21client/gql"

type GetProjectConsistencyInfo_Variables struct {
	GoalID string `json:"goalId"`
}


type GetProjectConsistencyInfo_Data struct {
	School21 GetProjectConsistencyInfo_Data_School21 `json:"school21"`
}

type GetProjectConsistencyInfo_Data_School21 struct {
	LoadGoalConsistencyInfo GetProjectConsistencyInfo_Data_LoadGoalConsistencyInfo `json:"loadGoalConsistencyInfo"`
	Typename                string                  `json:"__typename"`
}

type GetProjectConsistencyInfo_Data_LoadGoalConsistencyInfo struct {
	GoalID       string `json:"goalId"`
	IsConsistent bool   `json:"isConsistent"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectConsistencyInfo(variables GetProjectConsistencyInfo_Variables) (GetProjectConsistencyInfo_Data, error) {
	request := gql.NewQueryRequest[GetProjectConsistencyInfo_Variables](
		"query getProjectConsistencyInfo($goalId: ID!) {\n  school21 {\n    loadGoalConsistencyInfo(goalId: $goalId) {\n      goalId\n      isConsistent\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetProjectConsistencyInfo_Data](ctx, request)
}
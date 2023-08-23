package requests

import "s21client/gql"

type Variables_GetProjectConsistencyInfo struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectConsistencyInfo struct {
	School21_GetProjectConsistencyInfo School21_GetProjectConsistencyInfo `json:"school21"`
}

type School21_GetProjectConsistencyInfo struct {
	LoadGoalConsistencyInfo_GetProjectConsistencyInfo LoadGoalConsistencyInfo_GetProjectConsistencyInfo `json:"loadGoalConsistencyInfo"`
	Typename                string                  `json:"__typename"`
}

type LoadGoalConsistencyInfo_GetProjectConsistencyInfo struct {
	GoalID       string `json:"goalId"`
	IsConsistent bool   `json:"isConsistent"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectConsistencyInfo(variables Variables_GetProjectConsistencyInfo) (Data_GetProjectConsistencyInfo, error) {
	request := gql.NewQueryRequest[Variables_GetProjectConsistencyInfo](
		"query getProjectConsistencyInfo($goalId: ID!) {   school21 {     loadGoalConsistencyInfo(goalId: $goalId) {       goalId       isConsistent       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetProjectConsistencyInfo](ctx, request)
}
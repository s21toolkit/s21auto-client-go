package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetProjectConsistencyInfo struct {
	GoalID string `json:"goalId"`
}


type Response_Data_GetProjectConsistencyInfo struct {
	Response_School21_GetProjectConsistencyInfo Response_School21_GetProjectConsistencyInfo `json:"school21"`
}

type Response_School21_GetProjectConsistencyInfo struct {
	Response_LoadGoalConsistencyInfo_GetProjectConsistencyInfo Response_LoadGoalConsistencyInfo_GetProjectConsistencyInfo `json:"loadGoalConsistencyInfo"`
	Typename                string                  `json:"__typename"`
}

type Response_LoadGoalConsistencyInfo_GetProjectConsistencyInfo struct {
	GoalID       string `json:"goalId"`
	IsConsistent bool   `json:"isConsistent"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectConsistencyInfo(variables Request_Variables_GetProjectConsistencyInfo) (Response_Data_GetProjectConsistencyInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_GetProjectConsistencyInfo](
		"query getProjectConsistencyInfo($goalId: ID!) {   school21 {     loadGoalConsistencyInfo(goalId: $goalId) {       goalId       isConsistent       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetProjectConsistencyInfo](ctx, request)
}
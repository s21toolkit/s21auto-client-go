package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetPenaltyList struct {
	Request_Page_GetPenaltyList     Request_Page_GetPenaltyList     `json:"page"`
	Statuses []string `json:"statuses"`
}

type Request_Page_GetPenaltyList struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Response_Data_GetPenaltyList struct {
	Response_School21_GetPenaltyList Response_School21_GetPenaltyList `json:"school21"`
}

type Response_School21_GetPenaltyList struct {
	GetMyPenalties    []interface{}      `json:"getMyPenalties"`
	GetPenaltyReasons []Response_GetPenaltyReason_GetPenaltyList `json:"getPenaltyReasons"`
	CountMyPenalties  int64              `json:"countMyPenalties"`
	Typename          string             `json:"__typename"`
}

type Response_GetPenaltyReason_GetPenaltyList struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltyList(variables Request_Variables_GetPenaltyList) (Response_Data_GetPenaltyList, error) {
	request := gql.NewQueryRequest[Request_Variables_GetPenaltyList](
		"query getPenaltyList($statuses: [String]!, $from: DateTime, $to: DateTime, $sorting: SortingField, $page: PagingInput!) {\n  school21 {\n    getMyPenalties(\n      statuses: $statuses\n      from: $from\n      to: $to\n      sorting: $sorting\n      page: $page\n    ) {\n      ...Penalty\n      __typename\n    }\n    getPenaltyReasons {\n      id\n      name\n      __typename\n    }\n    countMyPenalties(statuses: $statuses)\n    __typename\n  }\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetPenaltyList](ctx, request)
}
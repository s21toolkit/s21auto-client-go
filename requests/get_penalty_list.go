package requests

import "github.com/s21toolkit/s21client/gql"

type GetPenaltyList_Variables struct {
	Page     GetPenaltyList_Variables_Page     `json:"page"`
	Statuses []string `json:"statuses"`
}

type GetPenaltyList_Variables_Page struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetPenaltyList_Data struct {
	School21 GetPenaltyList_Data_School21 `json:"school21"`
}

type GetPenaltyList_Data_School21 struct {
	GetMyPenalties    []interface{}      `json:"getMyPenalties"`
	GetPenaltyReasons []GetPenaltyList_Data_GetPenaltyReason `json:"getPenaltyReasons"`
	CountMyPenalties  int64              `json:"countMyPenalties"`
	Typename          string             `json:"__typename"`
}

type GetPenaltyList_Data_GetPenaltyReason struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetPenaltyList(variables GetPenaltyList_Variables) (GetPenaltyList_Data, error) {
	request := gql.NewQueryRequest[GetPenaltyList_Variables](
		"query getPenaltyList($statuses: [String]!, $from: DateTime, $to: DateTime, $sorting: SortingField, $page: PagingInput!) {\n  school21 {\n    getMyPenalties(\n      statuses: $statuses\n      from: $from\n      to: $to\n      sorting: $sorting\n      page: $page\n    ) {\n      ...Penalty\n      __typename\n    }\n    getPenaltyReasons {\n      id\n      name\n      __typename\n    }\n    countMyPenalties(statuses: $statuses)\n    __typename\n  }\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetPenaltyList_Data](ctx, request)
}
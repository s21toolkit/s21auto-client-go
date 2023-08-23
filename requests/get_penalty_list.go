package requests

import "s21client/gql"

type Variables_GetPenaltyList struct {
	Page_GetPenaltyList     Page_GetPenaltyList     `json:"page"`
	Statuses []string `json:"statuses"`
}

type Page_GetPenaltyList struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetPenaltyList struct {
	School21_GetPenaltyList School21_GetPenaltyList `json:"school21"`
}

type School21_GetPenaltyList struct {
	GetMyPenalties    []interface{}      `json:"getMyPenalties"`
	GetPenaltyReasons []GetPenaltyReason_GetPenaltyList `json:"getPenaltyReasons"`
	CountMyPenalties  int64              `json:"countMyPenalties"`
	Typename          string             `json:"__typename"`
}

type GetPenaltyReason_GetPenaltyList struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetPenaltyList(variables Variables_GetPenaltyList) (Data_GetPenaltyList, error) {
	request := gql.NewQueryRequest[Variables_GetPenaltyList](
		"query getPenaltyList($statuses: [String]!, $from: DateTime, $to: DateTime, $sorting: SortingField, $page: PagingInput!) {   school21 {     getMyPenalties(       statuses: $statuses       from: $from       to: $to       sorting: $sorting       page: $page     ) {       ...Penalty       __typename     }     getPenaltyReasons {       id       name       __typename     }     countMyPenalties(statuses: $statuses)     __typename   } }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetPenaltyList](ctx, request)
}
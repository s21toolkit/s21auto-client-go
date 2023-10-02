package requests

import "github.com/s21toolkit/s21client/gql"

type GetUserNotificationsCount_Variables struct {
	WasReadIncluded bool `json:"wasReadIncluded"`
}


type GetUserNotificationsCount_Data struct {
	Student GetUserNotificationsCount_Data_Student `json:"student"`
}

type GetUserNotificationsCount_Data_Student struct {
	GetS21NotificationsCount int64  `json:"getS21NotificationsCount"`
	Typename                 string `json:"__typename"`
}


func (ctx *RequestContext) GetUserNotificationsCount(variables GetUserNotificationsCount_Variables) (GetUserNotificationsCount_Data, error) {
	request := gql.NewQueryRequest[GetUserNotificationsCount_Variables](
		"query getUserNotificationsCount($wasReadIncluded: Boolean) {\n  student {\n    getS21NotificationsCount(wasReadIncluded: $wasReadIncluded)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetUserNotificationsCount_Data](ctx, request)
}
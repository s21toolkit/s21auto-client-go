package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUserNotificationsCount struct {
	WasReadIncluded bool `json:"wasReadIncluded"`
}


type Data_GetUserNotificationsCount struct {
	Data_Student_GetUserNotificationsCount Data_Student_GetUserNotificationsCount `json:"student"`
}

type Data_Student_GetUserNotificationsCount struct {
	GetS21NotificationsCount int64  `json:"getS21NotificationsCount"`
	Typename                 string `json:"__typename"`
}


func (ctx *RequestContext) GetUserNotificationsCount(variables Variables_GetUserNotificationsCount) (Data_GetUserNotificationsCount, error) {
	request := gql.NewQueryRequest[Variables_GetUserNotificationsCount](
		"query getUserNotificationsCount($wasReadIncluded: Boolean) {\n  student {\n    getS21NotificationsCount(wasReadIncluded: $wasReadIncluded)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUserNotificationsCount](ctx, request)
}
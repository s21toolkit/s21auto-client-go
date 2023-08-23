package requests

import "s21client/gql"

type Variables_GetUserNotificationsCount struct {
	WasReadIncluded bool `json:"wasReadIncluded"`
}


type Data_GetUserNotificationsCount struct {
	Student_GetUserNotificationsCount Student_GetUserNotificationsCount `json:"student"`
}

type Student_GetUserNotificationsCount struct {
	GetS21NotificationsCount int64  `json:"getS21NotificationsCount"`
	Typename                 string `json:"__typename"`
}

func (ctx *RequestContext) GetUserNotificationsCount(variables Variables_GetUserNotificationsCount) (Data_GetUserNotificationsCount, error) {
	request := gql.NewQueryRequest[Variables_GetUserNotificationsCount](
		"query getUserNotificationsCount($wasReadIncluded: Boolean) {   student {     getS21NotificationsCount(wasReadIncluded: $wasReadIncluded)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetUserNotificationsCount](ctx, request)
}
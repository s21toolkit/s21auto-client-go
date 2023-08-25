package requests

import "s21client/gql"

type Request_Variables_GetUserNotificationsCount struct {
	WasReadIncluded bool `json:"wasReadIncluded"`
}


type Response_Data_GetUserNotificationsCount struct {
	Response_Student_GetUserNotificationsCount Response_Student_GetUserNotificationsCount `json:"student"`
}

type Response_Student_GetUserNotificationsCount struct {
	GetS21NotificationsCount int64  `json:"getS21NotificationsCount"`
	Typename                 string `json:"__typename"`
}

func (ctx *RequestContext) GetUserNotificationsCount(variables Request_Variables_GetUserNotificationsCount) (Response_Data_GetUserNotificationsCount, error) {
	request := gql.NewQueryRequest[Request_Variables_GetUserNotificationsCount](
		"query getUserNotificationsCount($wasReadIncluded: Boolean) {   student {     getS21NotificationsCount(wasReadIncluded: $wasReadIncluded)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetUserNotificationsCount](ctx, request)
}
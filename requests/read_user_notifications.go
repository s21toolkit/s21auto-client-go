package requests

import "s21client/gql"

type Variables_ReadUserNotifications struct {
	NotificationIDS []string `json:"notificationIds"`
}


type Data_ReadUserNotifications struct {
	Student_ReadUserNotifications Student_ReadUserNotifications `json:"student"`
}

type Student_ReadUserNotifications struct {
	ReadNotifications []string `json:"readNotifications"`
	Typename          string   `json:"__typename"`
}

func (ctx *RequestContext) ReadUserNotifications(variables Variables_ReadUserNotifications) (Data_ReadUserNotifications, error) {
	request := gql.NewQueryRequest[Variables_ReadUserNotifications](
		"mutation readUserNotifications($notificationIds: [ID!]!) {   student {     readNotifications(notificationIds: $notificationIds)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_ReadUserNotifications](ctx, request)
}
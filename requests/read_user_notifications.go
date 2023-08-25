package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_ReadUserNotifications struct {
	NotificationIDS []string `json:"notificationIds"`
}


type Response_Data_ReadUserNotifications struct {
	Response_Student_ReadUserNotifications Response_Student_ReadUserNotifications `json:"student"`
}

type Response_Student_ReadUserNotifications struct {
	ReadNotifications []string `json:"readNotifications"`
	Typename          string   `json:"__typename"`
}

func (ctx *RequestContext) ReadUserNotifications(variables Request_Variables_ReadUserNotifications) (Response_Data_ReadUserNotifications, error) {
	request := gql.NewQueryRequest[Request_Variables_ReadUserNotifications](
		"mutation readUserNotifications($notificationIds: [ID!]!) {   student {     readNotifications(notificationIds: $notificationIds)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_ReadUserNotifications](ctx, request)
}
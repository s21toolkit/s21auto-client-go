package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_ReadUserNotifications struct {
	NotificationIDS []string `json:"notificationIds"`
}


type Data_ReadUserNotifications struct {
	Data_Student_ReadUserNotifications Data_Student_ReadUserNotifications `json:"student"`
}

type Data_Student_ReadUserNotifications struct {
	ReadNotifications []string `json:"readNotifications"`
	Typename          string   `json:"__typename"`
}


func (ctx *RequestContext) ReadUserNotifications(variables Variables_ReadUserNotifications) (Data_ReadUserNotifications, error) {
	request := gql.NewQueryRequest[Variables_ReadUserNotifications](
		"mutation readUserNotifications($notificationIds: [ID!]!) {\n  student {\n    readNotifications(notificationIds: $notificationIds)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_ReadUserNotifications](ctx, request)
}
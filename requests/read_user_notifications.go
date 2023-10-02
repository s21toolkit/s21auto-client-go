package requests

import "github.com/s21toolkit/s21client/gql"

type ReadUserNotifications_Variables struct {
	NotificationIDS []string `json:"notificationIds"`
}


type ReadUserNotifications_Data struct {
	Student ReadUserNotifications_Data_Student `json:"student"`
}

type ReadUserNotifications_Data_Student struct {
	ReadNotifications []string `json:"readNotifications"`
	Typename          string   `json:"__typename"`
}


func (ctx *RequestContext) ReadUserNotifications(variables ReadUserNotifications_Variables) (ReadUserNotifications_Data, error) {
	request := gql.NewQueryRequest[ReadUserNotifications_Variables](
		"mutation readUserNotifications($notificationIds: [ID!]!) {\n  student {\n    readNotifications(notificationIds: $notificationIds)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[ReadUserNotifications_Data](ctx, request)
}
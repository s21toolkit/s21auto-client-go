package requests

import "github.com/s21toolkit/s21client/gql"

type GetMySuggestedActivities_Variables struct {
	Page GetMySuggestedActivities_Variables_Page `json:"page"`
}

type GetMySuggestedActivities_Variables_Page struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetMySuggestedActivities_Data struct {
	School21 GetMySuggestedActivities_Data_School21 `json:"school21"`
}

type GetMySuggestedActivities_Data_School21 struct {
	GetMySuggestedActivities []interface{} `json:"getMySuggestedActivities"`
	Typename                 string        `json:"__typename"`
}


func (ctx *RequestContext) GetMySuggestedActivities(variables GetMySuggestedActivities_Variables) (GetMySuggestedActivities_Data, error) {
	request := gql.NewQueryRequest[GetMySuggestedActivities_Variables](
		"query getMySuggestedActivities($page: PagingInput!, $statuses: [ParticipantEventStatus]) {\n  school21 {\n    getMySuggestedActivities(page: $page, statuses: $statuses) {\n      id\n      start\n      end\n      eventType\n      description\n      eventCode\n      activity {\n        organizers {\n          id\n          login\n          __typename\n        }\n        eventCreator {\n          id\n          login\n          __typename\n        }\n        comments {\n          type\n          createTs\n          comment\n          __typename\n        }\n        averageFeedbackRating\n        isVisible\n        activityType\n        status\n        activityEventId\n        eventId\n        name\n        description\n        location\n        currentStudentsCount\n        maxStudentCount\n        isRegistered\n        isInWaitList\n        isWaitListActive\n        stopRegisterDate\n        beginDate\n        endDate\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetMySuggestedActivities_Data](ctx, request)
}
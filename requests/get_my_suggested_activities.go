package requests

import "s21client/gql"

type Request_Variables_GetMySuggestedActivities struct {
	Request_Page_GetMySuggestedActivities Request_Page_GetMySuggestedActivities `json:"page"`
}

type Request_Page_GetMySuggestedActivities struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Response_Data_GetMySuggestedActivities struct {
	Response_School21_GetMySuggestedActivities Response_School21_GetMySuggestedActivities `json:"school21"`
}

type Response_School21_GetMySuggestedActivities struct {
	GetMySuggestedActivities []interface{} `json:"getMySuggestedActivities"`
	Typename                 string        `json:"__typename"`
}

func (ctx *RequestContext) GetMySuggestedActivities(variables Request_Variables_GetMySuggestedActivities) (Response_Data_GetMySuggestedActivities, error) {
	request := gql.NewQueryRequest[Request_Variables_GetMySuggestedActivities](
		"query getMySuggestedActivities($page: PagingInput!, $statuses: [ParticipantEventStatus]) {   school21 {     getMySuggestedActivities(page: $page, statuses: $statuses) {       id       start       end       eventType       description       eventCode       activity {         organizers {           id           login           __typename         }         eventCreator {           id           login           __typename         }         comments {           type           createTs           comment           __typename         }         averageFeedbackRating         isVisible         activityType         status         activityEventId         eventId         name         description         location         currentStudentsCount         maxStudentCount         isRegistered         isInWaitList         isWaitListActive         stopRegisterDate         beginDate         endDate         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetMySuggestedActivities](ctx, request)
}
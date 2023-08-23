package requests

import "s21client/gql"

type Variables_CalendarGetMyBookings struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_CalendarGetMyBookings struct {
	Student_CalendarGetMyBookings Student_CalendarGetMyBookings `json:"student"`
}

type Student_CalendarGetMyBookings struct {
	GetMyCalendarBookings []interface{} `json:"getMyCalendarBookings"`
	Typename              string        `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetMyBookings(variables Variables_CalendarGetMyBookings) (Data_CalendarGetMyBookings, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetMyBookings](
		"query calendarGetMyBookings($from: DateTime!, $to: DateTime!) {   student {     getMyCalendarBookings(from: $from, to: $to) {       ...CalendarReviewBooking       __typename     }     __typename   } }  fragment CalendarReviewBooking on CalendarBooking {   id   answerId   eventSlotId   task {     id     goalId     goalName     studentTaskAdditionalAttributes {       cookiesCount       __typename     }     assignmentType     __typename   }   eventSlot {     id     start     end     event {       eventUserRole       __typename     }     __typename   }   verifierUser {     ...CalendarReviewUser     __typename   }   verifiableStudent {     id     user {       ...CalendarReviewUser       __typename     }     __typename   }   bookingStatus   team {     ...ProjectTeamMembers     __typename   }   isOnline   vcLinkUrl   __typename }  fragment CalendarReviewUser on User {   id   login   __typename }  fragment ProjectTeamMembers on ProjectTeamMembers {   id   teamLead {     ...ProjectTeamMember     __typename   }   members {     ...ProjectTeamMember     __typename   }   invitedUsers {     ...ProjectTeamMember     __typename   }   teamName   teamStatus   minTeamMemberCount   maxTeamMemberCount   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_CalendarGetMyBookings](ctx, request)
}
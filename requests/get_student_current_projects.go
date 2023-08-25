package requests

import "s21client/gql"

type Request_Variables_GetStudentCurrentProjects struct {
	UserID string `json:"userId"`
}


type Response_Data_GetStudentCurrentProjects struct {
	Response_Student_GetStudentCurrentProjects Response_Student_GetStudentCurrentProjects `json:"student"`
}

type Response_Student_GetStudentCurrentProjects struct {
	GetStudentCurrentProjects []Response_GetStudentCurrentProject_GetStudentCurrentProjects `json:"getStudentCurrentProjects"`
	Typename                  string                     `json:"__typename"`
}

type Response_GetStudentCurrentProject_GetStudentCurrentProjects struct {
	GoalID                         int64       `json:"goalId"`
	Name                           string      `json:"name"`
	Description                    string      `json:"description"`
	Experience                     int64       `json:"experience"`
	DateTime                       interface{} `json:"dateTime"`
	FinalPercentage                interface{} `json:"finalPercentage"`
	Laboriousness                  int64       `json:"laboriousness"`
	ExecutionType                  string      `json:"executionType"`
	GoalStatus                     string      `json:"goalStatus"`
	CourseType                     interface{} `json:"courseType"`
	DisplayedCourseStatus          interface{} `json:"displayedCourseStatus"`
	AmountAnswers                  *int64      `json:"amountAnswers"`
	AmountMembers                  *int64      `json:"amountMembers"`
	AmountJoinedMembers            *int64      `json:"amountJoinedMembers"`
	AmountReviewedAnswers          *int64      `json:"amountReviewedAnswers"`
	AmountCodeReviewMembers        interface{} `json:"amountCodeReviewMembers"`
	AmountCurrentCodeReviewMembers interface{} `json:"amountCurrentCodeReviewMembers"`
	GroupName                      string      `json:"groupName"`
	LocalCourseID                  interface{} `json:"localCourseId"`
	Typename                       string      `json:"__typename"`
}

func (ctx *RequestContext) GetStudentCurrentProjects(variables Request_Variables_GetStudentCurrentProjects) (Response_Data_GetStudentCurrentProjects, error) {
	request := gql.NewQueryRequest[Request_Variables_GetStudentCurrentProjects](
		"query getStudentCurrentProjects($userId: ID!) {   student {     getStudentCurrentProjects(userId: $userId) {       ...StudentProjectItem       __typename     }     __typename   } }  fragment StudentProjectItem on StudentItem {   goalId   name   description   experience   dateTime   finalPercentage   laboriousness   executionType   goalStatus   courseType   displayedCourseStatus   amountAnswers   amountMembers   amountJoinedMembers   amountReviewedAnswers   amountCodeReviewMembers   amountCurrentCodeReviewMembers   groupName   localCourseId   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetStudentCurrentProjects](ctx, request)
}
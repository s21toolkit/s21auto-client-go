package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetProjectAttemptEvaluationsInfo struct {
	GoalID    string `json:"goalId"`
	StudentID string `json:"studentId"`
}


type Response_Data_GetProjectAttemptEvaluationsInfo struct {
	Response_School21_GetProjectAttemptEvaluationsInfo Response_School21_GetProjectAttemptEvaluationsInfo `json:"school21"`
}

type Response_School21_GetProjectAttemptEvaluationsInfo struct {
	Response_GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo []Response_GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo `json:"getProjectAttemptEvaluationsInfo_V1"`
	Typename                           string                               `json:"__typename"`
}

type Response_GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo struct {
	StudentAnswerID      *string        `json:"studentAnswerId"`
	StudentGoalAttemptID string         `json:"studentGoalAttemptId"`
	Response_AttemptResult_GetProjectAttemptEvaluationsInfo        *Response_AttemptResult_GetProjectAttemptEvaluationsInfo `json:"attemptResult"`
	Team                 interface{}    `json:"team"`
	Response_P2P_GetProjectAttemptEvaluationsInfo                  []Response_P2P_GetProjectAttemptEvaluationsInfo          `json:"p2p"`
	Response_Auto_GetProjectAttemptEvaluationsInfo                 Response_Auto_GetProjectAttemptEvaluationsInfo           `json:"auto"`
	Response_CodeReview_GetProjectAttemptEvaluationsInfo           Response_CodeReview_GetProjectAttemptEvaluationsInfo     `json:"codeReview"`
	Typename             string         `json:"__typename"`
	AttemptStatus        *string        `json:"attemptStatus,omitempty"`
}

type Response_AttemptResult_GetProjectAttemptEvaluationsInfo struct {
	FinalPointProject      int64  `json:"finalPointProject"`
	FinalPercentageProject int64  `json:"finalPercentageProject"`
	ResultModuleCompletion string `json:"resultModuleCompletion"`
	ResultDate             string `json:"resultDate"`
	Typename               string `json:"__typename"`
}

type Response_Auto_GetProjectAttemptEvaluationsInfo struct {
	Status             string      `json:"status"`
	ReceivedPercentage int64       `json:"receivedPercentage"`
	EndTimeCheck       interface{} `json:"endTimeCheck"`
	ResultInfo         interface{} `json:"resultInfo"`
	Typename           string      `json:"__typename"`
}

type Response_CodeReview_GetProjectAttemptEvaluationsInfo struct {
	AverageMark        *string       `json:"averageMark"`
	StudentCodeReviews []interface{} `json:"studentCodeReviews"`
	Typename           string        `json:"__typename"`
}

type Response_P2P_GetProjectAttemptEvaluationsInfo struct {
	Status    string     `json:"status"`
	Response_Checklist_GetProjectAttemptEvaluationsInfo *Response_Checklist_GetProjectAttemptEvaluationsInfo `json:"checklist"`
	Typename  string     `json:"__typename"`
}

type Response_Checklist_GetProjectAttemptEvaluationsInfo struct {
	ID                 string          `json:"id"`
	ChecklistID        string          `json:"checklistId"`
	EndTimeCheck       string          `json:"endTimeCheck"`
	StartTimeCheck     string          `json:"startTimeCheck"`
	Response_Reviewer_GetProjectAttemptEvaluationsInfo           Response_Reviewer_GetProjectAttemptEvaluationsInfo        `json:"reviewer"`
	Response_ReviewFeedback_GetProjectAttemptEvaluationsInfo     *Response_ReviewFeedback_GetProjectAttemptEvaluationsInfo `json:"reviewFeedback"`
	Comment            string          `json:"comment"`
	ReceivedPoint      int64           `json:"receivedPoint"`
	ReceivedPercentage int64           `json:"receivedPercentage"`
	QuickAction        interface{}     `json:"quickAction"`
	CheckType          string          `json:"checkType"`
	Video              interface{}     `json:"video"`
	Typename           string          `json:"__typename"`
}

type Response_ReviewFeedback_GetProjectAttemptEvaluationsInfo struct {
	ID                           string                        `json:"id"`
	Comment                      string                        `json:"comment"`
	Response_FilledChecklist_GetProjectAttemptEvaluationsInfo              Response_FilledChecklist_GetProjectAttemptEvaluationsInfo               `json:"filledChecklist"`
	ReviewFeedbackCategoryValues []Response_ReviewFeedbackCategoryValue_GetProjectAttemptEvaluationsInfo `json:"reviewFeedbackCategoryValues"`
	Typename                     string                        `json:"__typename"`
}

type Response_FilledChecklist_GetProjectAttemptEvaluationsInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Response_ReviewFeedbackCategoryValue_GetProjectAttemptEvaluationsInfo struct {
	FeedbackCategory string `json:"feedbackCategory"`
	FeedbackValue    string `json:"feedbackValue"`
	ID               string `json:"id"`
	Typename         string `json:"__typename"`
}

type Response_Reviewer_GetProjectAttemptEvaluationsInfo struct {
	AvatarURL          string        `json:"avatarUrl"`
	Login              string        `json:"login"`
	BusinessAdminRoles []interface{} `json:"businessAdminRoles"`
	Typename           string        `json:"__typename"`
}

func (ctx *RequestContext) GetProjectAttemptEvaluationsInfo(variables Request_Variables_GetProjectAttemptEvaluationsInfo) (Response_Data_GetProjectAttemptEvaluationsInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_GetProjectAttemptEvaluationsInfo](
		"query getProjectAttemptEvaluationsInfo($goalId: ID!, $studentId: UUID!) {\n  school21 {\n    getProjectAttemptEvaluationsInfo_V1(goalId: $goalId, studentId: $studentId) {\n      ...ProjectAttemptEvaluations_V1\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectAttemptEvaluations_V1 on ProjectAttemptEvaluationsInfo_V1 {\n  studentAnswerId\n  studentGoalAttemptId\n  attemptResult {\n    ...AtemptResult\n    __typename\n  }\n  team {\n    ...AttemptTeamWithMembers\n    __typename\n  }\n  p2p {\n    ...P2PEvaluation\n    __typename\n  }\n  auto {\n    status\n    receivedPercentage\n    endTimeCheck\n    resultInfo\n    __typename\n  }\n  codeReview {\n    averageMark\n    studentCodeReviews {\n      user {\n        avatarUrl\n        login\n        __typename\n      }\n      finalMark\n      markTime\n      reviewerCommentsCount\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment AtemptResult on StudentGoalAttempt {\n  finalPointProject\n  finalPercentageProject\n  resultModuleCompletion\n  resultDate\n  __typename\n}\n\nfragment AttemptTeamWithMembers on TeamWithMembers {\n  team {\n    id\n    name\n    __typename\n  }\n  members {\n    ...TeamMemberWithRole\n    __typename\n  }\n  __typename\n}\n\nfragment TeamMemberWithRole on TeamMember {\n  role\n  user {\n    ...ProjectTeamMember\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n\nfragment P2PEvaluation on P2PEvaluationInfo {\n  status\n  checklist {\n    ...Checklist\n    __typename\n  }\n  __typename\n}\n\nfragment Checklist on FilledChecklist {\n  id\n  checklistId\n  endTimeCheck\n  startTimeCheck\n  reviewer {\n    avatarUrl\n    login\n    businessAdminRoles {\n      id\n      school {\n        id\n        organizationType\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  reviewFeedback {\n    ...EvaluationFeedback\n    __typename\n  }\n  comment\n  receivedPoint\n  receivedPercentage\n  quickAction\n  checkType\n  video {\n    ...P2PReviewVideo\n    __typename\n  }\n  __typename\n}\n\nfragment EvaluationFeedback on ReviewFeedback {\n  id\n  comment\n  filledChecklist {\n    id\n    __typename\n  }\n  reviewFeedbackCategoryValues {\n    feedbackCategory\n    feedbackValue\n    id\n    __typename\n  }\n  __typename\n}\n\nfragment P2PReviewVideo on OnlineReviewVideo {\n  link\n  status\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetProjectAttemptEvaluationsInfo](ctx, request)
}
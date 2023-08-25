package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetProjectInfo struct {
	GoalID string `json:"goalId"`
}


type Response_Data_GetProjectInfo struct {
	Response_Student_GetProjectInfo Response_Student_GetProjectInfo `json:"student"`
}

type Response_Student_GetProjectInfo struct {
	Response_GetModuleByID_GetProjectInfo                Response_GetModuleByID_GetProjectInfo                `json:"getModuleById"`
	Response_GetModuleCoverInformation_GetProjectInfo    Response_GetModuleCoverInformation_GetProjectInfo    `json:"getModuleCoverInformation"`
	Response_GetP2PChecksInfo_GetProjectInfo             Response_GetP2PChecksInfo_GetProjectInfo             `json:"getP2PChecksInfo"`
	Response_GetStudentCodeReviewByGoalID_GetProjectInfo Response_GetStudentCodeReviewByGoalID_GetProjectInfo `json:"getStudentCodeReviewByGoalId"`
	Typename                     string                       `json:"__typename"`
}

type Response_GetModuleByID_GetProjectInfo struct {
	ID                                string        `json:"id"`
	ModuleTitle                       string        `json:"moduleTitle"`
	FinalPercentage                   *int64        `json:"finalPercentage"`
	FinalPoint                        *int64        `json:"finalPoint"`
	GoalExecutionType                 string        `json:"goalExecutionType"`
	DisplayedGoalStatus               string        `json:"displayedGoalStatus"`
	AccessBeforeStartProgress         bool          `json:"accessBeforeStartProgress"`
	ResultModuleCompletion            *string       `json:"resultModuleCompletion"`
	FinishedExecutionDateByScheduler  interface{}   `json:"finishedExecutionDateByScheduler"`
	DurationFromStageSubjectGroupPlan int64         `json:"durationFromStageSubjectGroupPlan"`
	CurrentAttemptNumber              int64         `json:"currentAttemptNumber"`
	IsDeadlineFree                    bool          `json:"isDeadlineFree"`
	IsRetryAvailable                  bool          `json:"isRetryAvailable"`
	LocalCourseID                     interface{}   `json:"localCourseId"`
	Response_TeamSettings_GetProjectInfo                      *Response_TeamSettings_GetProjectInfo `json:"teamSettings"`
	Response_StudyModule_GetProjectInfo                       Response_StudyModule_GetProjectInfo   `json:"studyModule"`
	Response_CurrentTask_GetProjectInfo                       Response_CurrentTask_GetProjectInfo   `json:"currentTask"`
	Typename                          string        `json:"__typename"`
}

type Response_CurrentTask_GetProjectInfo struct {
	ID           string          `json:"id"`
	TaskID       string          `json:"taskId"`
	Task         Response_CurrentTaskTask_GetProjectInfo `json:"task"`
	Response_LastAnswer_GetProjectInfo   *Response_LastAnswer_GetProjectInfo     `json:"lastAnswer"`
	Response_TeamSettings_GetProjectInfo *Response_TeamSettings_GetProjectInfo   `json:"teamSettings"`
	Typename     string          `json:"__typename"`
}

type Response_LastAnswer_GetProjectInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Response_CurrentTaskTask_GetProjectInfo struct {
	ID                              string                          `json:"id"`
	AssignmentType                  string                          `json:"assignmentType"`
	Response_StudentTaskAdditionalAttributes_GetProjectInfo Response_StudentTaskAdditionalAttributes_GetProjectInfo `json:"studentTaskAdditionalAttributes"`
	CheckTypes                      []string                        `json:"checkTypes"`
	Typename                        string                          `json:"__typename"`
}

type Response_StudentTaskAdditionalAttributes_GetProjectInfo struct {
	CookiesCount       int64  `json:"cookiesCount"`
	MaxCodeReviewCount int64  `json:"maxCodeReviewCount"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CiCDMode           string `json:"ciCdMode"`
	Typename           string `json:"__typename"`
}

type Response_TeamSettings_GetProjectInfo struct {
	TeamCreateOption    string `json:"teamCreateOption"`
	MinAmountMember     int64  `json:"minAmountMember"`
	MaxAmountMember     int64  `json:"maxAmountMember"`
	EnableSurrenderTeam bool   `json:"enableSurrenderTeam"`
	Typename            string `json:"__typename"`
}

type Response_StudyModule_GetProjectInfo struct {
	ID            string        `json:"id"`
	Idea          string        `json:"idea"`
	Duration      int64         `json:"duration"`
	GoalPoint     int64         `json:"goalPoint"`
	Response_RetrySettings_GetProjectInfo Response_RetrySettings_GetProjectInfo `json:"retrySettings"`
	Levels        []Response_Level_GetProjectInfo       `json:"levels"`
	Typename      string        `json:"__typename"`
}

type Response_Level_GetProjectInfo struct {
	ID           string        `json:"id"`
	GoalElements []Response_GoalElement_GetProjectInfo `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type Response_GoalElement_GetProjectInfo struct {
	ID       string        `json:"id"`
	Tasks    []Response_TaskElement_GetProjectInfo `json:"tasks"`
	Typename string        `json:"__typename"`
}

type Response_TaskElement_GetProjectInfo struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Typename string `json:"__typename"`
}

type Response_RetrySettings_GetProjectInfo struct {
	MaxModuleAttempts   int64  `json:"maxModuleAttempts"`
	IsUnlimitedAttempts bool   `json:"isUnlimitedAttempts"`
	Typename            string `json:"__typename"`
}

type Response_GetModuleCoverInformation_GetProjectInfo struct {
	IsOwnStudentTimeline bool              `json:"isOwnStudentTimeline"`
	SoftSkills           []Response_SoftSkill_GetProjectInfo       `json:"softSkills"`
	Response_Timeline_GetProjectInfo             []Response_Timeline_GetProjectInfo        `json:"timeline"`
	Response_ProjectStatistics_GetProjectInfo    Response_ProjectStatistics_GetProjectInfo `json:"projectStatistics"`
	Typename             string            `json:"__typename"`
}

type Response_ProjectStatistics_GetProjectInfo struct {
	RegisteredStudents        int64                   `json:"registeredStudents"`
	InProgressStudents        int64                   `json:"inProgressStudents"`
	EvaluationStudents        int64                   `json:"evaluationStudents"`
	FinishedStudents          int64                   `json:"finishedStudents"`
	AcceptedStudents          int64                   `json:"acceptedStudents"`
	FailedStudents            int64                   `json:"failedStudents"`
	RetriedStudentsPercentage int64                   `json:"retriedStudentsPercentage"`
	Response_GroupProjectStatistics_GetProjectInfo    *Response_GroupProjectStatistics_GetProjectInfo `json:"groupProjectStatistics"`
	Typename                  string                  `json:"__typename"`
}

type Response_GroupProjectStatistics_GetProjectInfo struct {
	InProgressTeams int64  `json:"inProgressTeams"`
	EvaluationTeams int64  `json:"evaluationTeams"`
	FinishedTeams   int64  `json:"finishedTeams"`
	AcceptedTeams   int64  `json:"acceptedTeams"`
	FailedTeams     int64  `json:"failedTeams"`
	Typename        string `json:"__typename"`
}

type Response_SoftSkill_GetProjectInfo struct {
	SoftSkillID       int64   `json:"softSkillId"`
	SoftSkillName     string  `json:"softSkillName"`
	TotalPower        int64   `json:"totalPower"`
	MaxPower          int64   `json:"maxPower"`
	CurrentUserPower  int64   `json:"currentUserPower"`
	AchievedUserPower *int64  `json:"achievedUserPower"`
	TeamRole          *string `json:"teamRole"`
	Typename          string  `json:"__typename"`
}

type Response_Timeline_GetProjectInfo struct {
	Type     string      `json:"type"`
	Status   string      `json:"status"`
	Start    interface{} `json:"start"`
	End      interface{} `json:"end"`
	Children []Response_Child_GetProjectInfo     `json:"children"`
	Typename string      `json:"__typename"`
}

type Response_Child_GetProjectInfo struct {
	Type        string  `json:"type"`
	ElementType string  `json:"elementType"`
	Status      string  `json:"status"`
	Start       *string `json:"start"`
	End         *string `json:"end"`
	Order       int64   `json:"order"`
	Typename    string  `json:"__typename"`
}

type Response_GetP2PChecksInfo_GetProjectInfo struct {
	CookiesCount         int64              `json:"cookiesCount"`
	PeriodOfVerification int64              `json:"periodOfVerification"`
	Response_ProjectReviewsInfo_GetProjectInfo   Response_ProjectReviewsInfo_GetProjectInfo `json:"projectReviewsInfo"`
	Typename             string             `json:"__typename"`
}

type Response_ProjectReviewsInfo_GetProjectInfo struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type Response_GetStudentCodeReviewByGoalID_GetProjectInfo struct {
	CountRound1     int64       `json:"countRound1"`
	CountRound2     int64       `json:"countRound2"`
	CodeReviewsInfo interface{} `json:"codeReviewsInfo"`
	Typename        string      `json:"__typename"`
}

func (ctx *RequestContext) GetProjectInfo(variables Request_Variables_GetProjectInfo) (Response_Data_GetProjectInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_GetProjectInfo](
		"query getProjectInfo($goalId: ID!) {   student {     getModuleById(goalId: $goalId) {       ...ProjectInfo       __typename     }     getModuleCoverInformation(goalId: $goalId) {       ...ModuleCoverInfo       __typename     }     getP2PChecksInfo(goalId: $goalId) {       ...P2PInfo       __typename     }     getStudentCodeReviewByGoalId(goalId: $goalId) {       ...StudentsCodeReview       __typename     }     __typename   } }  fragment ProjectInfo on StudentModule {   id   moduleTitle   finalPercentage   finalPoint   goalExecutionType   displayedGoalStatus   accessBeforeStartProgress   resultModuleCompletion   finishedExecutionDateByScheduler   durationFromStageSubjectGroupPlan   currentAttemptNumber   isDeadlineFree   isRetryAvailable   localCourseId   teamSettings {     ...teamSettingsInfo     __typename   }   studyModule {     id     idea     duration     goalPoint     retrySettings {       ...RetrySettings       __typename     }     levels {       id       goalElements {         id         tasks {           id           taskId           __typename         }         __typename       }       __typename     }     __typename   }   currentTask {     ...CurrentInternshipTaskInfo     __typename   }   __typename }  fragment teamSettingsInfo on TeamSettings {   teamCreateOption   minAmountMember   maxAmountMember   enableSurrenderTeam   __typename }  fragment RetrySettings on ModuleAttemptsSettings {   maxModuleAttempts   isUnlimitedAttempts   __typename }  fragment CurrentInternshipTaskInfo on StudentTask {   id   taskId   task {     id     assignmentType     studentTaskAdditionalAttributes {       cookiesCount       maxCodeReviewCount       codeReviewCost       ciCdMode       __typename     }     checkTypes     __typename   }   lastAnswer {     id     __typename   }   teamSettings {     ...teamSettingsInfo     __typename   }   __typename }  fragment ModuleCoverInfo on ModuleCoverInformation {   isOwnStudentTimeline   softSkills {     softSkillId     softSkillName     totalPower     maxPower     currentUserPower     achievedUserPower     teamRole     __typename   }   timeline {     ...TimelineItem     __typename   }   projectStatistics {     ...ProjectStatistics     __typename   }   __typename }  fragment TimelineItem on ProjectTimelineItem {   type   status   start   end   children {     ...TimelineItemChildren     __typename   }   __typename }  fragment TimelineItemChildren on ProjectTimelineItem {   type   elementType   status   start   end   order   __typename }  fragment ProjectStatistics on ProjectStatistics {   registeredStudents   inProgressStudents   evaluationStudents   finishedStudents   acceptedStudents   failedStudents   retriedStudentsPercentage   groupProjectStatistics {     ...GroupProjectStatistics     __typename   }   __typename }  fragment GroupProjectStatistics on GroupProjectStatistics {   inProgressTeams   evaluationTeams   finishedTeams   acceptedTeams   failedTeams   __typename }  fragment P2PInfo on P2PChecksInfo {   cookiesCount   periodOfVerification   projectReviewsInfo {     ...ProjectReviewsInfo     __typename   }   __typename }  fragment ProjectReviewsInfo on ProjectReviewsInfo {   reviewByStudentCount   relevantReviewByStudentsCount   reviewByInspectionStaffCount   relevantReviewByInspectionStaffCount   __typename }  fragment StudentsCodeReview on StudentCodeReviewsWithCountRound {   countRound1   countRound2   codeReviewsInfo {     maxCodeReviewCount     codeReviewDuration     codeReviewCost     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetProjectInfo](ctx, request)
}
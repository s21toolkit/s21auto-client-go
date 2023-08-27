package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetProjectInfo struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectInfo struct {
	Data_Student_GetProjectInfo Data_Student_GetProjectInfo `json:"student"`
}

type Data_Student_GetProjectInfo struct {
	Data_GetModuleByID_GetProjectInfo                Data_GetModuleByID_GetProjectInfo                `json:"getModuleById"`
	Data_GetModuleCoverInformation_GetProjectInfo    Data_GetModuleCoverInformation_GetProjectInfo    `json:"getModuleCoverInformation"`
	Data_GetP2PChecksInfo_GetProjectInfo             Data_GetP2PChecksInfo_GetProjectInfo             `json:"getP2PChecksInfo"`
	Data_GetStudentCodeReviewByGoalID_GetProjectInfo Data_GetStudentCodeReviewByGoalID_GetProjectInfo `json:"getStudentCodeReviewByGoalId"`
	Typename                     string                       `json:"__typename"`
}

type Data_GetModuleByID_GetProjectInfo struct {
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
	CourseBaseParameters              interface{}   `json:"courseBaseParameters"`
	Data_TeamSettings_GetProjectInfo                      *Data_TeamSettings_GetProjectInfo `json:"teamSettings"`
	Data_StudyModule_GetProjectInfo                       Data_StudyModule_GetProjectInfo   `json:"studyModule"`
	Data_CurrentTask_GetProjectInfo                       Data_CurrentTask_GetProjectInfo   `json:"currentTask"`
	Typename                          string        `json:"__typename"`
}

type Data_CurrentTask_GetProjectInfo struct {
	ID           string          `json:"id"`
	TaskID       string          `json:"taskId"`
	Task         Data_CurrentTaskTask_GetProjectInfo `json:"task"`
	Data_LastAnswer_GetProjectInfo   *Data_LastAnswer_GetProjectInfo     `json:"lastAnswer"`
	Data_TeamSettings_GetProjectInfo *Data_TeamSettings_GetProjectInfo   `json:"teamSettings"`
	Typename     string          `json:"__typename"`
}

type Data_LastAnswer_GetProjectInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Data_CurrentTaskTask_GetProjectInfo struct {
	ID                              string                          `json:"id"`
	AssignmentType                  string                          `json:"assignmentType"`
	Data_StudentTaskAdditionalAttributes_GetProjectInfo Data_StudentTaskAdditionalAttributes_GetProjectInfo `json:"studentTaskAdditionalAttributes"`
	CheckTypes                      []string                        `json:"checkTypes"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_GetProjectInfo struct {
	CookiesCount       int64  `json:"cookiesCount"`
	MaxCodeReviewCount int64  `json:"maxCodeReviewCount"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CiCDMode           string `json:"ciCdMode"`
	Typename           string `json:"__typename"`
}

type Data_TeamSettings_GetProjectInfo struct {
	TeamCreateOption    string `json:"teamCreateOption"`
	MinAmountMember     int64  `json:"minAmountMember"`
	MaxAmountMember     int64  `json:"maxAmountMember"`
	EnableSurrenderTeam bool   `json:"enableSurrenderTeam"`
	Typename            string `json:"__typename"`
}

type Data_StudyModule_GetProjectInfo struct {
	ID            string        `json:"id"`
	Idea          string        `json:"idea"`
	Duration      int64         `json:"duration"`
	GoalPoint     int64         `json:"goalPoint"`
	Data_RetrySettings_GetProjectInfo Data_RetrySettings_GetProjectInfo `json:"retrySettings"`
	Levels        []Data_Level_GetProjectInfo       `json:"levels"`
	Typename      string        `json:"__typename"`
}

type Data_Level_GetProjectInfo struct {
	ID           string        `json:"id"`
	GoalElements []Data_GoalElement_GetProjectInfo `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type Data_GoalElement_GetProjectInfo struct {
	ID       string        `json:"id"`
	Tasks    []Data_TaskElement_GetProjectInfo `json:"tasks"`
	Typename string        `json:"__typename"`
}

type Data_TaskElement_GetProjectInfo struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Typename string `json:"__typename"`
}

type Data_RetrySettings_GetProjectInfo struct {
	MaxModuleAttempts   int64  `json:"maxModuleAttempts"`
	IsUnlimitedAttempts bool   `json:"isUnlimitedAttempts"`
	Typename            string `json:"__typename"`
}

type Data_GetModuleCoverInformation_GetProjectInfo struct {
	IsOwnStudentTimeline bool              `json:"isOwnStudentTimeline"`
	SoftSkills           []Data_SoftSkill_GetProjectInfo       `json:"softSkills"`
	Data_Timeline_GetProjectInfo             []Data_Timeline_GetProjectInfo        `json:"timeline"`
	Data_ProjectStatistics_GetProjectInfo    Data_ProjectStatistics_GetProjectInfo `json:"projectStatistics"`
	Typename             string            `json:"__typename"`
}

type Data_ProjectStatistics_GetProjectInfo struct {
	RegisteredStudents        int64                   `json:"registeredStudents"`
	InProgressStudents        int64                   `json:"inProgressStudents"`
	EvaluationStudents        int64                   `json:"evaluationStudents"`
	FinishedStudents          int64                   `json:"finishedStudents"`
	AcceptedStudents          int64                   `json:"acceptedStudents"`
	FailedStudents            int64                   `json:"failedStudents"`
	RetriedStudentsPercentage int64                   `json:"retriedStudentsPercentage"`
	Data_GroupProjectStatistics_GetProjectInfo    *Data_GroupProjectStatistics_GetProjectInfo `json:"groupProjectStatistics"`
	Typename                  string                  `json:"__typename"`
}

type Data_GroupProjectStatistics_GetProjectInfo struct {
	InProgressTeams int64  `json:"inProgressTeams"`
	EvaluationTeams int64  `json:"evaluationTeams"`
	FinishedTeams   int64  `json:"finishedTeams"`
	AcceptedTeams   int64  `json:"acceptedTeams"`
	FailedTeams     int64  `json:"failedTeams"`
	Typename        string `json:"__typename"`
}

type Data_SoftSkill_GetProjectInfo struct {
	SoftSkillID       int64   `json:"softSkillId"`
	SoftSkillName     string  `json:"softSkillName"`
	TotalPower        int64   `json:"totalPower"`
	MaxPower          int64   `json:"maxPower"`
	CurrentUserPower  int64   `json:"currentUserPower"`
	AchievedUserPower *int64  `json:"achievedUserPower"`
	TeamRole          *string `json:"teamRole"`
	Typename          string  `json:"__typename"`
}

type Data_Timeline_GetProjectInfo struct {
	Type     string      `json:"type"`
	Status   string      `json:"status"`
	Start    interface{} `json:"start"`
	End      interface{} `json:"end"`
	Children []Data_Child_GetProjectInfo     `json:"children"`
	Typename string      `json:"__typename"`
}

type Data_Child_GetProjectInfo struct {
	Type        string  `json:"type"`
	ElementType string  `json:"elementType"`
	Status      string  `json:"status"`
	Start       *string `json:"start"`
	End         *string `json:"end"`
	Order       int64   `json:"order"`
	Typename    string  `json:"__typename"`
}

type Data_GetP2PChecksInfo_GetProjectInfo struct {
	CookiesCount         int64              `json:"cookiesCount"`
	PeriodOfVerification int64              `json:"periodOfVerification"`
	Data_ProjectReviewsInfo_GetProjectInfo   Data_ProjectReviewsInfo_GetProjectInfo `json:"projectReviewsInfo"`
	Typename             string             `json:"__typename"`
}

type Data_ProjectReviewsInfo_GetProjectInfo struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type Data_GetStudentCodeReviewByGoalID_GetProjectInfo struct {
	CountRound1     int64       `json:"countRound1"`
	CountRound2     int64       `json:"countRound2"`
	CodeReviewsInfo interface{} `json:"codeReviewsInfo"`
	Typename        string      `json:"__typename"`
}


func (ctx *RequestContext) GetProjectInfo(variables Variables_GetProjectInfo) (Data_GetProjectInfo, error) {
	request := gql.NewQueryRequest[Variables_GetProjectInfo](
		"query getProjectInfo($goalId: ID!) {\n  student {\n    getModuleById(goalId: $goalId) {\n      ...ProjectInfo\n      __typename\n    }\n    getModuleCoverInformation(goalId: $goalId) {\n      ...ModuleCoverInfo\n      __typename\n    }\n    getP2PChecksInfo(goalId: $goalId) {\n      ...P2PInfo\n      __typename\n    }\n    getStudentCodeReviewByGoalId(goalId: $goalId) {\n      ...StudentsCodeReview\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectInfo on StudentModule {\n  id\n  moduleTitle\n  finalPercentage\n  finalPoint\n  goalExecutionType\n  displayedGoalStatus\n  accessBeforeStartProgress\n  resultModuleCompletion\n  finishedExecutionDateByScheduler\n  durationFromStageSubjectGroupPlan\n  currentAttemptNumber\n  isDeadlineFree\n  isRetryAvailable\n  localCourseId\n  courseBaseParameters {\n    isGradedCourse\n    __typename\n  }\n  teamSettings {\n    ...teamSettingsInfo\n    __typename\n  }\n  studyModule {\n    id\n    idea\n    duration\n    goalPoint\n    retrySettings {\n      ...RetrySettings\n      __typename\n    }\n    levels {\n      id\n      goalElements {\n        id\n        tasks {\n          id\n          taskId\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  currentTask {\n    ...CurrentInternshipTaskInfo\n    __typename\n  }\n  __typename\n}\n\nfragment teamSettingsInfo on TeamSettings {\n  teamCreateOption\n  minAmountMember\n  maxAmountMember\n  enableSurrenderTeam\n  __typename\n}\n\nfragment RetrySettings on ModuleAttemptsSettings {\n  maxModuleAttempts\n  isUnlimitedAttempts\n  __typename\n}\n\nfragment CurrentInternshipTaskInfo on StudentTask {\n  id\n  taskId\n  task {\n    id\n    assignmentType\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      maxCodeReviewCount\n      codeReviewCost\n      ciCdMode\n      __typename\n    }\n    checkTypes\n    __typename\n  }\n  lastAnswer {\n    id\n    __typename\n  }\n  teamSettings {\n    ...teamSettingsInfo\n    __typename\n  }\n  __typename\n}\n\nfragment ModuleCoverInfo on ModuleCoverInformation {\n  isOwnStudentTimeline\n  softSkills {\n    softSkillId\n    softSkillName\n    totalPower\n    maxPower\n    currentUserPower\n    achievedUserPower\n    teamRole\n    __typename\n  }\n  timeline {\n    ...TimelineItem\n    __typename\n  }\n  projectStatistics {\n    ...ProjectStatistics\n    __typename\n  }\n  __typename\n}\n\nfragment TimelineItem on ProjectTimelineItem {\n  type\n  status\n  start\n  end\n  children {\n    ...TimelineItemChildren\n    __typename\n  }\n  __typename\n}\n\nfragment TimelineItemChildren on ProjectTimelineItem {\n  type\n  elementType\n  status\n  start\n  end\n  order\n  __typename\n}\n\nfragment ProjectStatistics on ProjectStatistics {\n  registeredStudents\n  inProgressStudents\n  evaluationStudents\n  finishedStudents\n  acceptedStudents\n  failedStudents\n  retriedStudentsPercentage\n  groupProjectStatistics {\n    ...GroupProjectStatistics\n    __typename\n  }\n  __typename\n}\n\nfragment GroupProjectStatistics on GroupProjectStatistics {\n  inProgressTeams\n  evaluationTeams\n  finishedTeams\n  acceptedTeams\n  failedTeams\n  __typename\n}\n\nfragment P2PInfo on P2PChecksInfo {\n  cookiesCount\n  periodOfVerification\n  projectReviewsInfo {\n    ...ProjectReviewsInfo\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectReviewsInfo on ProjectReviewsInfo {\n  reviewByStudentCount\n  relevantReviewByStudentsCount\n  reviewByInspectionStaffCount\n  relevantReviewByInspectionStaffCount\n  __typename\n}\n\nfragment StudentsCodeReview on StudentCodeReviewsWithCountRound {\n  countRound1\n  countRound2\n  codeReviewsInfo {\n    maxCodeReviewCount\n    codeReviewDuration\n    codeReviewCost\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetProjectInfo](ctx, request)
}
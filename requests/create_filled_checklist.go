package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CreateFilledChecklist struct {
	StudentAnswerID string `json:"studentAnswerId"`
}


type Data_CreateFilledChecklist struct {
	Data_Student_CreateFilledChecklist Data_Student_CreateFilledChecklist `json:"student"`
}

type Data_Student_CreateFilledChecklist struct {
	Data_CreateFilledChecklist_CreateFilledChecklist Data_CreateFilledChecklist_CreateFilledChecklist `json:"createFilledChecklist"`
	Typename              string                `json:"__typename"`
}

type Data_CreateFilledChecklist_CreateFilledChecklist struct {
	ID                      string                  `json:"id"`
	Data_GitlabStudentProjectURL_CreateFilledChecklist Data_GitlabStudentProjectURL_CreateFilledChecklist `json:"gitlabStudentProjectUrl"`
	Data_Checklist_CreateFilledChecklist               Data_Checklist_CreateFilledChecklist               `json:"checklist"`
	Data_ModuleInfoP2P_CreateFilledChecklist           Data_ModuleInfoP2P_CreateFilledChecklist           `json:"moduleInfoP2P"`
	Data_ProgressCheckInfo_CreateFilledChecklist       Data_ProgressCheckInfo_CreateFilledChecklist       `json:"progressCheckInfo"`
	Data_VerifiableUsers_CreateFilledChecklist         Data_VerifiableUsers_CreateFilledChecklist         `json:"verifiableUsers"`
	Video                   interface{}             `json:"video"`
	Typename                string                  `json:"__typename"`
}

type Data_Checklist_CreateFilledChecklist struct {
	Language     string        `json:"language"`
	Introduction string        `json:"introduction"`
	Guidelines   string        `json:"guidelines"`
	Data_SectionList_CreateFilledChecklist  []Data_SectionList_CreateFilledChecklist `json:"sectionList"`
	QuickActions []string      `json:"quickActions"`
	Typename     string        `json:"__typename"`
}

type Data_SectionList_CreateFilledChecklist struct {
	ChecklistSectionID string         `json:"checklistSectionId"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	KindQuestionID     string         `json:"kindQuestionId"`
	Data_QuestionList_CreateFilledChecklist       []Data_QuestionList_CreateFilledChecklist `json:"questionList"`
	Typename           string         `json:"__typename"`
}

type Data_QuestionList_CreateFilledChecklist struct {
	SectionQuestionID   string              `json:"sectionQuestionId"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	Data_TaskAssessmentScale_CreateFilledChecklist Data_TaskAssessmentScale_CreateFilledChecklist `json:"taskAssessmentScale"`
	Typename            string              `json:"__typename"`
}

type Data_TaskAssessmentScale_CreateFilledChecklist struct {
	CriterionScaleID string        `json:"criterionScaleId"`
	Type             string        `json:"type"`
	Description      string        `json:"description"`
	ScaleWeights     []Data_ScaleWeight_CreateFilledChecklist `json:"scaleWeights"`
	Typename         string        `json:"__typename"`
}

type Data_ScaleWeight_CreateFilledChecklist struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Typename string `json:"__typename"`
}

type Data_GitlabStudentProjectURL_CreateFilledChecklist struct {
	SSHLink   string `json:"sshLink"`
	HTTPSLink string `json:"httpsLink"`
	Typename  string `json:"__typename"`
}

type Data_ModuleInfoP2P_CreateFilledChecklist struct {
	ModuleName           string `json:"moduleName"`
	ExecutionType        string `json:"executionType"`
	PeriodOfVerification int64  `json:"periodOfVerification"`
	Typename             string `json:"__typename"`
}

type Data_ProgressCheckInfo_CreateFilledChecklist struct {
	ReviewUserCount         int64  `json:"reviewUserCount"`
	ReviewUserCountExecuted int64  `json:"reviewUserCountExecuted"`
	Typename                string `json:"__typename"`
}

type Data_VerifiableUsers_CreateFilledChecklist struct {
	Data_TeamWithMembers_CreateFilledChecklist Data_TeamWithMembers_CreateFilledChecklist `json:"teamWithMembers"`
	Data_User_CreateFilledChecklist            interface{}     `json:"user"`
	Typename        string          `json:"__typename"`
}

type Data_TeamWithMembers_CreateFilledChecklist struct {
	Data_Team_CreateFilledChecklist     Data_Team_CreateFilledChecklist     `json:"team"`
	Members  []Data_Member_CreateFilledChecklist `json:"members"`
	Typename string   `json:"__typename"`
}

type Data_Member_CreateFilledChecklist struct {
	Data_User_CreateFilledChecklist     Data_User_CreateFilledChecklist   `json:"user"`
	Role     string `json:"role"`
	Typename string `json:"__typename"`
}

type Data_User_CreateFilledChecklist struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	Data_UserExperience_CreateFilledChecklist Data_UserExperience_CreateFilledChecklist `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_CreateFilledChecklist struct {
	Data_Level_CreateFilledChecklist            Data_Level_CreateFilledChecklist  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Data_Level_CreateFilledChecklist struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Data_Range_CreateFilledChecklist     Data_Range_CreateFilledChecklist  `json:"range"`
	Typename  string `json:"__typename"`
}

type Data_Range_CreateFilledChecklist struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Data_Team_CreateFilledChecklist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CreateFilledChecklist(variables Variables_CreateFilledChecklist) (Data_CreateFilledChecklist, error) {
	request := gql.NewQueryRequest[Variables_CreateFilledChecklist](
		"mutation createFilledChecklist($studentAnswerId: ID!) {\n  student {\n    createFilledChecklist(studentAnswerId: $studentAnswerId) {\n      id\n      gitlabStudentProjectUrl {\n        sshLink\n        httpsLink\n        __typename\n      }\n      checklist {\n        ...FormChecklist\n        __typename\n      }\n      moduleInfoP2P {\n        ...FilledChecklistModuleInfo\n        __typename\n      }\n      progressCheckInfo {\n        reviewUserCount\n        reviewUserCountExecuted\n        __typename\n      }\n      verifiableUsers {\n        teamWithMembers {\n          team {\n            id\n            name\n            __typename\n          }\n          members {\n            ...TeamMember\n            __typename\n          }\n          __typename\n        }\n        user {\n          ...TeamMemberUser\n          __typename\n        }\n        __typename\n      }\n      video {\n        filledChecklistId\n        link\n        status\n        statusDetails\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment FormChecklist on Checklist {\n  language\n  introduction\n  guidelines\n  sectionList {\n    ...FormChecklistSection\n    __typename\n  }\n  quickActions\n  __typename\n}\n\nfragment FormChecklistSection on ChecklistSection {\n  checklistSectionId\n  name\n  description\n  kindQuestionId\n  questionList {\n    ...FormChecklistQuestion\n    __typename\n  }\n  __typename\n}\n\nfragment FormChecklistQuestion on SectionQuestion {\n  sectionQuestionId\n  name\n  description\n  taskAssessmentScale {\n    criterionScaleId\n    type\n    description\n    scaleWeights {\n      key\n      value\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment FilledChecklistModuleInfo on ModuleInfoP2P {\n  moduleName\n  executionType\n  periodOfVerification\n  __typename\n}\n\nfragment TeamMember on TeamMember {\n  user {\n    ...TeamMemberUser\n    __typename\n  }\n  role\n  __typename\n}\n\nfragment TeamMemberUser on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      levelCode\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CreateFilledChecklist](ctx, request)
}
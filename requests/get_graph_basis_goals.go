package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetGraphBasisGoals struct {
	StudentID string `json:"studentId"`
}


type Response_Data_GetGraphBasisGoals struct {
	Response_Student_GetGraphBasisGoals Response_Student_GetGraphBasisGoals `json:"student"`
}

type Response_Student_GetGraphBasisGoals struct {
	Response_GetBasisGraph_GetGraphBasisGoals Response_GetBasisGraph_GetGraphBasisGoals `json:"getBasisGraph"`
	Typename      string        `json:"__typename"`
}

type Response_GetBasisGraph_GetGraphBasisGoals struct {
	IsIntensiveGraphAvailable bool        `json:"isIntensiveGraphAvailable"`
	GraphNodes                []Response_GraphNode_GetGraphBasisGoals `json:"graphNodes"`
	Typename                  string      `json:"__typename"`
}

type Response_GraphNode_GetGraphBasisGoals struct {
	GraphNodeID     string           `json:"graphNodeId"`
	NodeCode        string           `json:"nodeCode"`
	StudyDirections []Response_StudyDirection_GetGraphBasisGoals `json:"studyDirections"`
	EntityType      string           `json:"entityType"`
	EntityID        string           `json:"entityId"`
	Goal            *Response_Course_GetGraphBasisGoals          `json:"goal"`
	Response_Course_GetGraphBasisGoals          *Response_Course_GetGraphBasisGoals          `json:"course"`
	ParentNodeCodes []string         `json:"parentNodeCodes"`
	Typename        string           `json:"__typename"`
}

type Response_Course_GetGraphBasisGoals struct {
	CourseType         *string     `json:"courseType,omitempty"`
	ProjectState       string      `json:"projectState"`
	ProjectName        string      `json:"projectName"`
	ProjectDescription string      `json:"projectDescription"`
	ProjectPoints      int64       `json:"projectPoints"`
	ProjectDate        interface{} `json:"projectDate"`
	Duration           int64       `json:"duration"`
	LocalCourseID      *int64      `json:"localCourseId,omitempty"`
	Typename           string      `json:"__typename"`
	GoalExecutionType  *string     `json:"goalExecutionType,omitempty"`
	IsMandatory        *bool       `json:"isMandatory,omitempty"`
}

type Response_StudyDirection_GetGraphBasisGoals struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	TextColor string `json:"textColor"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetGraphBasisGoals(variables Request_Variables_GetGraphBasisGoals) (Response_Data_GetGraphBasisGoals, error) {
	request := gql.NewQueryRequest[Request_Variables_GetGraphBasisGoals](
		"query getGraphBasisGoals($studentId: UUID!) {\n  student {\n    getBasisGraph(studentId: $studentId) {\n      isIntensiveGraphAvailable\n      graphNodes {\n        ...BasisGraphNode\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment BasisGraphNode on GraphNode {\n  graphNodeId\n  nodeCode\n  studyDirections {\n    id\n    name\n    color\n    textColor\n    __typename\n  }\n  entityType\n  entityId\n  goal {\n    goalExecutionType\n    projectState\n    projectName\n    projectDescription\n    projectPoints\n    projectDate\n    duration\n    isMandatory\n    __typename\n  }\n  course {\n    courseType\n    projectState\n    projectName\n    projectDescription\n    projectPoints\n    projectDate\n    duration\n    localCourseId\n    __typename\n  }\n  parentNodeCodes\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetGraphBasisGoals](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetTasks struct {
	IDS []string `json:"ids"`
}


type Data_GetTasks struct {
	Data_Student_GetTasks Data_Student_GetTasks `json:"student"`
}

type Data_Student_GetTasks struct {
	GetTasksByIDS []Data_GetTasksByID_GetTasks `json:"getTasksByIds"`
	Typename      string         `json:"__typename"`
}

type Data_GetTasksByID_GetTasks struct {
	ID       string `json:"id"`
	Data_Task_GetTasks     Data_Task_GetTasks   `json:"task"`
	Typename string `json:"__typename"`
}

type Data_Task_GetTasks struct {
	ID       string  `json:"id"`
	Data_Content_GetTasks  Data_Content_GetTasks `json:"content"`
	Typename string  `json:"__typename"`
}

type Data_Content_GetTasks struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetTasks(variables Variables_GetTasks) (Data_GetTasks, error) {
	request := gql.NewQueryRequest[Variables_GetTasks](
		"query getTasks($ids: [ID!]!) {\n  student {\n    getTasksByIds(ids: $ids) {\n      ...StudentTaskInProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentTaskInProject on StudentTask {\n  id\n  task {\n    id\n    content {\n      id\n      body\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetTasks](ctx, request)
}
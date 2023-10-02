package requests

import "github.com/s21toolkit/s21client/gql"

type GetTasks_Variables struct {
	IDS []string `json:"ids"`
}


type GetTasks_Data struct {
	Student GetTasks_Data_Student `json:"student"`
}

type GetTasks_Data_Student struct {
	GetTasksByIDS []GetTasks_Data_GetTasksByID `json:"getTasksByIds"`
	Typename      string         `json:"__typename"`
}

type GetTasks_Data_GetTasksByID struct {
	ID       string `json:"id"`
	Task     GetTasks_Data_Task   `json:"task"`
	Typename string `json:"__typename"`
}

type GetTasks_Data_Task struct {
	ID       string  `json:"id"`
	Content  GetTasks_Data_Content `json:"content"`
	Typename string  `json:"__typename"`
}

type GetTasks_Data_Content struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetTasks(variables GetTasks_Variables) (GetTasks_Data, error) {
	request := gql.NewQueryRequest[GetTasks_Variables](
		"query getTasks($ids: [ID!]!) {\n  student {\n    getTasksByIds(ids: $ids) {\n      ...StudentTaskInProject\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentTaskInProject on StudentTask {\n  id\n  task {\n    id\n    content {\n      id\n      body\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetTasks_Data](ctx, request)
}
package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetAsapWidgets struct {
}


type Data_GetAsapWidgets struct {
	Data_Student_GetAsapWidgets Data_Student_GetAsapWidgets `json:"student"`
}

type Data_Student_GetAsapWidgets struct {
	Data_GetAsapWidgetList_GetAsapWidgets Data_GetAsapWidgetList_GetAsapWidgets `json:"getAsapWidgetList"`
	Typename          string            `json:"__typename"`
}

type Data_GetAsapWidgetList_GetAsapWidgets struct {
	Data_WidgetList_GetAsapWidgets []Data_WidgetList_GetAsapWidgets `json:"widgetList"`
	Typename   string       `json:"__typename"`
}

type Data_WidgetList_GetAsapWidgets struct {
	ShortImg       string      `json:"shortImg"`
	ShortTitle     string      `json:"shortTitle"`
	ShortURL       interface{} `json:"shortUrl"`
	StartDate      string      `json:"startDate"`
	FinishDate     string      `json:"finishDate"`
	ShowFinishDate bool        `json:"showFinishDate"`
	FullTitle      string      `json:"fullTitle"`
	Text           string      `json:"text"`
	FullImgURL     string      `json:"fullImgUrl"`
	ADTTypeID      int64       `json:"adtTypeId"`
	ADTWidgetID    int64       `json:"adtWidgetId"`
	Typename       string      `json:"__typename"`
}


func (ctx *RequestContext) GetAsapWidgets(variables Variables_GetAsapWidgets) (Data_GetAsapWidgets, error) {
	request := gql.NewQueryRequest[Variables_GetAsapWidgets](
		"query getAsapWidgets {\n  student {\n    getAsapWidgetList {\n      widgetList {\n        ...AsapWidget\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment AsapWidget on AsapWidgetInfo {\n  shortImg\n  shortTitle\n  shortUrl\n  startDate\n  finishDate\n  showFinishDate\n  fullTitle\n  text\n  fullImgUrl\n  adtTypeId\n  adtWidgetId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetAsapWidgets](ctx, request)
}
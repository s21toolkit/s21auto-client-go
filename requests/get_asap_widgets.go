package requests

import "github.com/s21toolkit/s21client/gql"

type GetAsapWidgets_Variables struct {
}


type GetAsapWidgets_Data struct {
	Student GetAsapWidgets_Data_Student `json:"student"`
}

type GetAsapWidgets_Data_Student struct {
	GetAsapWidgetList GetAsapWidgets_Data_GetAsapWidgetList `json:"getAsapWidgetList"`
	Typename          string            `json:"__typename"`
}

type GetAsapWidgets_Data_GetAsapWidgetList struct {
	WidgetList []GetAsapWidgets_Data_WidgetList `json:"widgetList"`
	Typename   string       `json:"__typename"`
}

type GetAsapWidgets_Data_WidgetList struct {
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


func (ctx *RequestContext) GetAsapWidgets(variables GetAsapWidgets_Variables) (GetAsapWidgets_Data, error) {
	request := gql.NewQueryRequest[GetAsapWidgets_Variables](
		"query getAsapWidgets {\n  student {\n    getAsapWidgetList {\n      widgetList {\n        ...AsapWidget\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment AsapWidget on AsapWidgetInfo {\n  shortImg\n  shortTitle\n  shortUrl\n  startDate\n  finishDate\n  showFinishDate\n  fullTitle\n  text\n  fullImgUrl\n  adtTypeId\n  adtWidgetId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetAsapWidgets_Data](ctx, request)
}
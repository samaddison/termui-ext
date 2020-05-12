package termui_ext

import (
	"encoding/json"
)

type ParagraphRenderer struct {
	Paragraph Paragraph
}

type paragraphData struct {
	Title string
	Text  string
}

func processParagraphWidgetData(data *WidgetData) (*paragraphData, error) {
	var paragraphData = paragraphData{}
	err := json.Unmarshal(data.Json, &paragraphData)
	if err != nil {
		return nil, err
	} else {
		return &paragraphData, nil
	}
}

func (renderer ParagraphRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer ParagraphRenderer) Render(data *WidgetData) error {
	widgetData, err := processParagraphWidgetData(data)
	if err != nil {
		return err
	}
	renderer.Paragraph.Title = widgetData.Title
	renderer.Paragraph.Text = widgetData.Text
	return nil
}

func (renderer ParagraphRenderer) PostRender(data *WidgetData) error {
	return nil
}

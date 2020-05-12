package termui_ext

import (
	"encoding/json"
)

type ListRenderer struct {
	List List
}

type listData struct {
	Rows []string
}

func processListWidgetData(data *WidgetData) (*listData, error) {
	var listData = listData{}
	err := json.Unmarshal(data.Json, &listData)
	if err != nil {
		return nil, err
	} else {
		return &listData, nil
	}
}

func (renderer ListRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer ListRenderer) Render(data *WidgetData) error {
	listWidgetData, err := processListWidgetData(data)
	if err != nil {
		return err
	}
	renderer.List.Rows = listWidgetData.Rows
	return nil
}

func (renderer ListRenderer) PostRender(data *WidgetData) error {
	return nil
}

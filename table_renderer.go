package termui_ext

import (
	"encoding/json"
)

type TableRenderer struct {
	Table Table
}

type tableData struct {
	Rows [][]string
}

func processTableWidgetData(data *WidgetData) (*tableData, error) {
	var tableData = tableData{}
	err := json.Unmarshal(data.Json, &tableData)
	if err != nil {
		return nil, err
	} else {
		return &tableData, nil
	}
}

func (renderer TableRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer TableRenderer) Render(data *WidgetData) error {
	listWidgetData, err := processTableWidgetData(data)
	if err != nil {
		return err
	}
	renderer.Table.Rows = listWidgetData.Rows
	return nil
}

func (renderer TableRenderer) PostRender(data *WidgetData) error {
	return nil
}

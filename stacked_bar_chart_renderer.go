package termui_ext

import (
	"encoding/json"
)

type StackedBarChartRenderer struct {
	StackedBarChart StackedBarChart
}

type stackedBarChartData struct {
	Data   [][]float64
	Labels []string
}

func processStackedBarChartWidgetData(data *WidgetData) (*stackedBarChartData, error) {
	var chartData = stackedBarChartData{
		Data: make([][]float64, 0),
	}
	err := json.Unmarshal(data.Json, &chartData)
	if err != nil {
		return nil, err
	} else {
		return &chartData, nil
	}
}

func (renderer StackedBarChartRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer StackedBarChartRenderer) Render(data *WidgetData) error {
	chartData, err := processStackedBarChartWidgetData(data)
	if err != nil {
		return err
	}
	renderer.StackedBarChart.Data = chartData.Data
	renderer.StackedBarChart.Labels = chartData.Labels
	return nil
}

func (renderer StackedBarChartRenderer) PostRender(data *WidgetData) error {
	return nil
}

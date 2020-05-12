package termui_ext

import (
	"encoding/json"
)

type BarChartRenderer struct {
	BarChart BarChart
}

type barChartData struct {
	Data   []float64
	Labels []string
}

func processBarChartWidgetData(data *WidgetData) (*barChartData, error) {
	var chartData = barChartData{
		Data:   make([]float64, 0),
		Labels: make([]string, 0),
	}
	err := json.Unmarshal(data.Json, &chartData)
	if err != nil {
		return nil, err
	} else {
		return &chartData, nil
	}
}

func (renderer BarChartRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer BarChartRenderer) Render(data *WidgetData) error {
	chartData, err := processBarChartWidgetData(data)
	if err != nil {
		return err
	}
	renderer.BarChart.Labels = chartData.Labels
	renderer.BarChart.Data = chartData.Data
	return nil
}

func (renderer BarChartRenderer) PostRender(data *WidgetData) error {
	return nil
}

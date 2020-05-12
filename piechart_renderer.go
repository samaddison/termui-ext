package termui_ext

import (
	"encoding/json"
)

type PieChartRenderer struct {
	PieChart PieChart
}

type pieChartData struct {
	Data   []float64
	Labels []string
}

func processPieChartWidgetData(data *WidgetData) (*pieChartData, error) {
	var chartData = pieChartData{
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

func (renderer PieChartRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer PieChartRenderer) Render(data *WidgetData) error {
	chartData, err := processPieChartWidgetData(data)
	if err != nil {
		return err
	}
	renderer.PieChart.Data = chartData.Data
	return nil
}

func (renderer PieChartRenderer) PostRender(data *WidgetData) error {
	return nil
}

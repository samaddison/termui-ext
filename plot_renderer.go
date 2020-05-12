package termui_ext

import (
	"encoding/json"
)

type PlotRenderer struct {
	Plot Plot
}

type plotData struct {
	Data [][]float64
}

func processPlotWidgetData(data *WidgetData) (*plotData, error) {
	var chartData = plotData{
		Data: make([][]float64, 0),
	}
	err := json.Unmarshal(data.Json, &chartData)
	if err != nil {
		return nil, err
	} else {
		return &chartData, nil
	}
}

func (renderer PlotRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer PlotRenderer) Render(data *WidgetData) error {
	chartData, err := processPlotWidgetData(data)
	if err != nil {
		return err
	}
	renderer.Plot.Data = chartData.Data
	return nil
}

func (renderer PlotRenderer) PostRender(data *WidgetData) error {
	return nil
}

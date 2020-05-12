package termui_ext

import (
	"encoding/json"
	"github.com/gizak/termui/v3/widgets"
)

type SparklineGroupRenderer struct {
	SparklineGroup SparklineGroup
}

type sparklineGroupData struct {
	Data [][]float64
}

func processSparklineGroupWidgetData(data *WidgetData) (*sparklineGroupData, error) {
	var chartData = sparklineGroupData{
		Data: make([][]float64, 0),
	}
	err := json.Unmarshal(data.Json, &chartData)
	if err != nil {
		return nil, err
	} else {
		return &chartData, nil
	}
}

func (renderer SparklineGroupRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer SparklineGroupRenderer) Render(data *WidgetData) error {
	chartData, err := processSparklineGroupWidgetData(data)
	if err != nil {
		return err
	}
	sparklines := make([]*widgets.Sparkline, 0)
	for _, v := range chartData.Data {
		sparkLine := widgets.NewSparkline()
		sparkLine.Data = v
		sparklines = append(sparklines, sparkLine)
	}
	renderer.SparklineGroup.Sparklines = sparklines
	return nil
}

func (renderer SparklineGroupRenderer) PostRender(data *WidgetData) error {
	return nil
}

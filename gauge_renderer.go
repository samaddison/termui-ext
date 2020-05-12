package termui_ext

import (
	"encoding/json"
)

type GaugeRenderer struct {
	Gauge Gauge
}

type gaugeData struct {
	Percent int
}

func processWidgetData(data *WidgetData) (*gaugeData, error) {
	var gaugeData = gaugeData{}
	err := json.Unmarshal(data.Json, &gaugeData)
	if err != nil {
		return nil, err
	} else {
		return &gaugeData, nil
	}
}

func (renderer GaugeRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer GaugeRenderer) Render(data *WidgetData) error {
	gaugeData, err := processWidgetData(data)
	if err != nil {
		return err
	}
	renderer.Gauge.Percent = gaugeData.Percent
	return nil
}

func (renderer GaugeRenderer) PostRender(data *WidgetData) error {
	return nil
}

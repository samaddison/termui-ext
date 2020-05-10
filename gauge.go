package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Gauge struct {
	*widgets.Gauge
	BaseWidget
}

func NewGaugeWithRenderer(provider DataProvider, renderer GaugeRenderer) *Gauge {
	widgetGauge := widgets.NewGauge()

	gauge := Gauge{
		Gauge: widgetGauge,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetGauge,
		},
	}
	renderer.Gauge = gauge
	gauge.BaseWidget.DataRenderer = renderer

	return &gauge
}

func NewGauge(provider DataProvider) *Gauge {
	dataRenderer := GaugeRenderer{}
	gauge := NewGaugeWithRenderer(provider, dataRenderer)
	return gauge
}

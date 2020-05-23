package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Gauge struct {
	*widgets.Gauge
	BaseWidget
}

func NewGaugeWithRenderer(provider DataProvider, renderer GaugeRenderer) *Gauge {
	widgetGauge := widgets.NewGauge()
	return NewGaugeWithExisting(widgetGauge, provider, renderer)
}

func NewGaugeWithExisting(gaugeWidget *widgets.Gauge, provider DataProvider, renderer GaugeRenderer) *Gauge {
	gauge := Gauge{
		Gauge: gaugeWidget,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     gaugeWidget,
		},
	}
	renderer.Gauge = gauge
	gauge.BaseWidget.DataRenderer = renderer

	return &gauge
}

func NewGaugeDefaultRenderer(provider DataProvider) *Gauge {
	dataRenderer := GaugeRenderer{}
	gauge := NewGaugeWithRenderer(provider, dataRenderer)
	return gauge
}

func NewGauge2(gaugeWidget *widgets.Gauge, provider DataProvider) *Gauge {
	dataRenderer := GaugeRenderer{}
	gauge := NewGaugeWithExisting(gaugeWidget, provider, dataRenderer)
	return gauge
}

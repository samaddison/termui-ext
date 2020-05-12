package termui_ext

import "github.com/gizak/termui/v3/widgets"

type BarChart struct {
	*widgets.BarChart
	BaseWidget
}

func NewBarChartWithRenderer(provider DataProvider, renderer BarChartRenderer) *BarChart {
	widgetBarChart := widgets.NewBarChart()

	barChart := BarChart{
		BarChart: widgetBarChart,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetBarChart,
		},
	}
	renderer.BarChart = barChart
	barChart.BaseWidget.DataRenderer = renderer

	return &barChart
}

func NewBarChart(provider DataProvider) *BarChart {
	dataRenderer := BarChartRenderer{}
	barChart := NewBarChartWithRenderer(provider, dataRenderer)
	return barChart
}

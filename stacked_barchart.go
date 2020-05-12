package termui_ext

import "github.com/gizak/termui/v3/widgets"

type StackedBarChart struct {
	*widgets.StackedBarChart
	BaseWidget
}

func NewStackedBarChartWithRenderer(provider DataProvider, renderer StackedBarChartRenderer) *StackedBarChart {
	widgetPlot := widgets.NewStackedBarChart()

	stackedBarChart := StackedBarChart{
		StackedBarChart: widgetPlot,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetPlot,
		},
	}
	renderer.StackedBarChart = stackedBarChart
	stackedBarChart.BaseWidget.DataRenderer = renderer

	return &stackedBarChart
}

func NewStackedBarChart(provider DataProvider) *StackedBarChart {
	dataRenderer := StackedBarChartRenderer{}
	chart := NewStackedBarChartWithRenderer(provider, dataRenderer)
	return chart
}

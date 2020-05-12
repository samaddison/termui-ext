package termui_ext

import "github.com/gizak/termui/v3/widgets"

type PieChart struct {
	*widgets.PieChart
	BaseWidget
}

func NewPieChartWithRenderer(provider DataProvider, renderer PieChartRenderer) *PieChart {
	widgetChart := widgets.NewPieChart()

	pieChart := PieChart{
		PieChart: widgetChart,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetChart,
		},
	}
	renderer.PieChart = pieChart
	pieChart.BaseWidget.DataRenderer = renderer

	return &pieChart
}

func NewPieChart(provider DataProvider) *PieChart {
	dataRenderer := PieChartRenderer{}
	pieChart := NewPieChartWithRenderer(provider, dataRenderer)
	return pieChart
}

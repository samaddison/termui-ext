package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Plot struct {
	*widgets.Plot
	BaseWidget
}

func NewPlotWithRenderer(provider DataProvider, renderer PlotRenderer) *Plot {
	widgetPlot := widgets.NewPlot()

	plot := Plot{
		Plot: widgetPlot,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetPlot,
		},
	}
	renderer.Plot = plot
	plot.BaseWidget.DataRenderer = renderer

	return &plot
}

func NewPlot(provider DataProvider) *Plot {
	dataRenderer := PlotRenderer{}
	chart := NewPlotWithRenderer(provider, dataRenderer)
	return chart
}

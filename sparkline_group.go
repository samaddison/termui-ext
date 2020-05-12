package termui_ext

import "github.com/gizak/termui/v3/widgets"

type SparklineGroup struct {
	*widgets.SparklineGroup
	BaseWidget
}

func NewSparklineGroupWithRenderer(provider DataProvider, renderer SparklineGroupRenderer) *SparklineGroup {
	widgetSparklineGroup := widgets.NewSparklineGroup()

	sparklineGroup := SparklineGroup{
		SparklineGroup: widgetSparklineGroup,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetSparklineGroup,
		},
	}
	renderer.SparklineGroup = sparklineGroup
	sparklineGroup.BaseWidget.DataRenderer = renderer

	return &sparklineGroup
}

func NewSparklineGroup(provider DataProvider) *SparklineGroup {
	dataRenderer := SparklineGroupRenderer{}
	chart := NewSparklineGroupWithRenderer(provider, dataRenderer)
	return chart
}

package termui_ext

import "github.com/gizak/termui/v3/widgets"

type List struct {
	*widgets.List
	BaseWidget
}

func NewListWithRenderer(provider DataProvider, renderer ListRenderer) *List {
	widgetList := widgets.NewList()

	list := List{
		List: widgetList,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetList,
		},
	}

	renderer.List = list
	list.BaseWidget.DataRenderer = renderer

	return &list
}

func NewList(provider DataProvider) *List {
	dataRenderer := ListRenderer{}
	gauge := NewListWithRenderer(provider, dataRenderer)
	return gauge
}

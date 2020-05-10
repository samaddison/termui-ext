package termui_ext

type DataRenderer interface {
	PreRender(*WidgetData) error
	Render(*WidgetData) error
	PostRender(*WidgetData) error
}

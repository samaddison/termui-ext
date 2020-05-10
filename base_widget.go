package termui_ext

import (
	ui "github.com/gizak/termui/v3"
	"time"
)

type BaseWidget struct {
	DataProvider
	DataRenderer
	Drawable ui.Drawable
}

func (widget *BaseWidget) Refresh(d time.Duration) {
	doEvery(d, widget.internalRefresh)
}

func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func (widget *BaseWidget) internalRefresh() {
	err := widget.PreRetrieve()
	if err != nil {
		return
	}
	widgetData, err := widget.Retrieve()
	err = widget.PostRetrieve(widgetData, err)
	if err != nil {
		return
	}
	err = widget.PreRender(widgetData)
	if err != nil {
		return
	}
	err = widget.Render(widgetData)
	if err != nil {
		return
	}
	err = widget.PostRender(widgetData)
	if err != nil {
		return
	}
	ui.Render(widget.Drawable)
}

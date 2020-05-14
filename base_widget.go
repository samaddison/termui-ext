package termui_ext

import (
	ui "github.com/gizak/termui/v3"
	"log"
	"time"
)

type BaseWidget struct {
	DataProvider
	DataRenderer
	Drawable ui.Drawable
	quit     chan struct{}
}

func (widget *BaseWidget) Refresh(d time.Duration) {
	doEvery(widget, d, widget.internalRefresh)
}

func (widget *BaseWidget) GoRefresh(d time.Duration) {
	go doEvery(widget, d, widget.internalRefresh)
}

func (widget *BaseWidget) OneRefresh() {
	widget.internalRefresh()
}

func doEvery(widget *BaseWidget, d time.Duration, f func()) {
	widget.quit = make(chan struct{})
	for _ = range time.Tick(d) {
		select {
		case <-widget.quit:
			return
		default:
			f()
		}
	}
}

func (widget *BaseWidget) Shutdown() {
	close(widget.quit)
}

func (widget *BaseWidget) internalRefresh() {
	err := widget.PreRetrieve()
	if err != nil {
		log.Print("Error: " + err.Error())
		return
	}
	widgetData, err := widget.Retrieve()
	err = widget.PostRetrieve(widgetData, err)
	if err != nil {
		log.Print("Error: " + err.Error())
		return
	}
	err = widget.PreRender(widgetData)
	if err != nil {
		log.Print("Error: " + err.Error())
		return
	}
	err = widget.Render(widgetData)
	if err != nil {
		log.Print("Error: " + err.Error())
		return
	}
	err = widget.PostRender(widgetData)
	if err != nil {
		log.Print("Error: " + err.Error())
		return
	}
	ui.Render(widget.Drawable)
}

package main

import (
	ui "github.com/gizak/termui/v3"
	"log"
	termui_ext "termui-ext"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "./docs/image_input.json"}

	imageWidget := termui_ext.NewImage(dataProvider)
	imageWidget.Title = "This is the title"
	imageWidget.SetRect(0, 0, 100, 50)

	ui.Render(imageWidget)

	go imageWidget.Refresh(3 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

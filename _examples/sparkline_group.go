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

	dataProvider := termui_ext.File{Path: "./docs/sparkline_input.json"}

	sparklineGroup := termui_ext.NewSparklineGroup(dataProvider)
	sparklineGroup.Title = "Sparklines"
	sparklineGroup.SetRect(5, 5, 100, 25)

	// The following call to Render cannot be done at this point because
	// there are no sparklines associated to the group and the termui library doesn't handle that very well.
	//ui.Render(sparklineGroup)

	go sparklineGroup.Refresh(3 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

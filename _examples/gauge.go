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

	dataProvider := termui_ext.File{Path: "./docs/gauge_input.json"}

	gauge := termui_ext.NewGauge(dataProvider)
	gauge.Title = "Slim Gauge"
	gauge.SetRect(0, 14, 50, 17)
	gauge.BarColor = ui.ColorRed
	gauge.BorderStyle.Fg = ui.ColorWhite
	gauge.TitleStyle.Fg = ui.ColorCyan

	ui.Render(gauge)

	go gauge.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

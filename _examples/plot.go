package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	termui_ext "termui-ext"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "./docs/plot_input.json"}

	plot := termui_ext.NewPlot(dataProvider)
	plot.Title = "Pie Chart"
	plot.SetRect(5, 5, 100, 25)
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.Marker = widgets.MarkerBraille
	plot.PlotType = widgets.ScatterPlot

	ui.Render(plot)

	go plot.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

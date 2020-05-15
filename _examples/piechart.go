package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"log"
	"math"
	termui_ext "termui-ext"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "./docs/piechart_input.json"}

	pieChart := termui_ext.NewPieChart(dataProvider)
	pieChart.Title = "Pie Chart"
	pieChart.SetRect(5, 5, 100, 25)
	pieChart.AngleOffset = -.5 * math.Pi
	pieChart.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

	ui.Render(pieChart)

	go pieChart.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	"log"
	"math"
	termui_ext "termui-ext"
	"time"
)

func main() {
	//mainBarChart()
	mainTabs()
}

func mainTabs() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress h or l to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 35, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	tabpane := widgets.NewTabPane("pierwszy", "drugi", "trzeci", "żółw", "four", "five")
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
		case 1:
			ui.Render(bc)
		}
	}

	ui.Render(header, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}
}

func mainStackedBarchart() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/stacked_barchart_input.json"}

	bc := termui_ext.NewStackedBarChart(dataProvider)
	bc.Title = "Stacked Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	go bc.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func mainSparkline() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/sparkline_input.json"}

	sparklineGroup := termui_ext.NewSparklineGroup(dataProvider)
	sparklineGroup.Title = "Pie Chart"
	sparklineGroup.SetRect(5, 5, 100, 25)

	// The following call to Render cannot be done at this point because
	// there are no sparklines associated to the group and the library doesn't handle that very well.
	//ui.Render(sparklineGroup)

	go sparklineGroup.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func mainPlot() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/plot_input.json"}

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

func mainPieChart() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/piechart_input.json"}

	bc := termui_ext.NewPieChart(dataProvider)
	bc.Title = "Pie Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.AngleOffset = -.5 * math.Pi
	bc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

	ui.Render(bc)

	go bc.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func mainParagraph() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/paragraph_input.json"}

	bc := termui_ext.NewParagraph(dataProvider)
	bc.Title = "Paragraph Title"
	bc.SetRect(5, 5, 100, 25)

	bc.BorderStyle.Fg = ui.ColorYellow

	ui.Render(bc)

	go bc.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}

}

func mainBarchart() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/barchart_input.json"}

	bc := termui_ext.NewBarChart(dataProvider)
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	go bc.Refresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func mainGauge() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/gauge_input.json"}

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

func main2() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	ui.Render(c)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

}

func main3() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sinFloat64 := (func() []float64 {
		n := 400
		data := make([]float64, n)
		for i := range data {
			data[i] = 1 + math.Sin(float64(i)/5)
		}
		return data
	})()

	sl := widgets.NewSparkline()
	sl.Data = sinFloat64[:100]
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "SparklineGroup"

	lc := widgets.NewPlot()
	lc.Title = "braille-mode Line Chart"
	lc.Data = append(lc.Data, sinFloat64)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorYellow

	gs := make([]*widgets.Gauge, 3)
	for i := range gs {
		gs[i] = widgets.NewGauge()
		gs[i].Percent = i * 10
		gs[i].BarColor = ui.ColorRed
	}

	ls := widgets.NewList()
	ls.Rows = []string{
		"[1] Downloading File 1",
		"",
		"",
		"",
		"[2] Downloading File 2",
		"",
		"",
		"",
		"[3] Uploading File 3",
	}
	ls.Border = false

	p := widgets.NewParagraph()
	p.Text = "<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget"
	p.Title = "Demonstration"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, slg),
			ui.NewCol(1.0/2, lc),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, ls),
			ui.NewCol(1.0/4,
				ui.NewRow(.9/3, gs[0]),
				ui.NewRow(.9/3, gs[1]),
				ui.NewRow(1.2/3, gs[2]),
			),
			ui.NewCol(1.0/2, p),
		),
	)

	ui.Render(grid)

	tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if tickerCount == 100 {
				return
			}
			for _, g := range gs {
				g.Percent = (g.Percent + 3) % 100
			}
			slg.Sparklines[0].Data = sinFloat64[tickerCount : tickerCount+100]
			lc.Data[0] = sinFloat64[2*tickerCount:]
			ui.Render(grid)
			tickerCount++
		}
	}
}

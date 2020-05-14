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
	mainStackedBarchart()
}

func mainImage() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/image_input.json"}

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

func mainTable() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dataProvider := termui_ext.File{Path: "/Users/samaddison/GolandProjects/termui-ext/docs/table_input.json"}

	table1 := termui_ext.NewTable(dataProvider)
	table1.RowSeparator = true
	table1.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table1.SetRect(5, 5, 100, 25)
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)

	//ui.Render(table1)

	go table1.Refresh(3 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func mainTree() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	nodes := []*widgets.TreeNode{
		{
			Value: nodeValue("Key 1"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 1.1"),
					Nodes: []*widgets.TreeNode{
						{
							Value: nodeValue("Key 1.1.1"),
							Nodes: nil,
						},
						{
							Value: nodeValue("Key 1.1.2"),
							Nodes: nil,
						},
					},
				},
				{
					Value: nodeValue("Key 1.2"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 2"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 2.1"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.2"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.3"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 3"),
			Nodes: nil,
		},
	}

	l := widgets.NewTree()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetNodes(nodes)

	x, y := ui.TerminalDimensions()

	l.SetRect(0, 0, x, y)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "<Enter>":
			l.ToggleExpand()
		case "G", "<End>":
			l.ScrollBottom()
		case "E":
			l.ExpandAll()
		case "C":
			l.CollapseAll()
		case "<Resize>":
			x, y := ui.TerminalDimensions()
			l.SetRect(0, 0, x, y)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
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

	bc.GoRefresh(5 * time.Second)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<C-d>":
			bc.Shutdown()
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

package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	"log"
	"math"
	"time"
)

func main() {
	mainStackedBarchart()
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

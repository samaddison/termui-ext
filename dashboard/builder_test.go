package dashboard

import (
	_ "fmt"
	ui "github.com/gizak/termui/v3"
	"log"
	"testing"
)

func TestBuildGrid(t *testing.T) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dashboardConfFile := "/Users/samaddison/GolandProjects/termui-ext/dashboard_examples/dashboard.json"
	grid, errs := BuildGrid(dashboardConfFile)
	if errs != nil {
		for _, v := range *errs {
			t.Errorf("Returned Error: %s\n", v.Error())
		}
		return
	}

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
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
		}
	}
}

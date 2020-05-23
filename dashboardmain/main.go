package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"io/ioutil"
	"log"
	"termui-ext/dashboard"
	"time"
)

func main() {
	buildGrid()
}

func readDashboard() {
	dashboardConfFile := "/Users/samaddison/GolandProjects/termui-ext/dashboard_examples/dashboard.json"
	jsonString, err := ioutil.ReadFile(dashboardConfFile)
	if err != nil {
		fmt.Print("Error when reading file: " + err.Error())
	} else {
		dashboardFromJSON, err := dashboard.BuildDashboardDefinitionFromJSON(string(jsonString))
		if err != nil {
			fmt.Print("Error when building dashboardFromJSON:")
		} else {
			fmt.Printf("%+v", dashboardFromJSON)
		}
	}
}

func buildGrid() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	dashboardConfFile := "/Users/samaddison/GolandProjects/termui-ext/dashboard_examples/dashboard.json"
	grid, errs := dashboard.BuildGrid(dashboardConfFile)
	if errs != nil {
		for _, v := range *errs {
			fmt.Printf("Returned Error: %s\n", v.Error())
		}
		return
	}

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	ui.Render(grid)

	go grid.StartRefresh(3 * time.Second)

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

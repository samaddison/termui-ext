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

	dataProvider := termui_ext.File{Path: "./docs/table_input.json"}

	table1 := termui_ext.NewTable(dataProvider)
	table1.RowSeparator = true
	table1.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table1.SetRect(5, 5, 100, 25)
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)

	// Termui doesn't properly handle tables with no rows, and at this point no rows have been created
	// so we need to comment out the following line.
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

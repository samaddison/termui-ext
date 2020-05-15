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

	dataProvider := termui_ext.File{Path: "./docs/paragraph_input.json"}

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

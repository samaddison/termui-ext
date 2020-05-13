package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Table struct {
	*widgets.Table
	BaseWidget
}

func NewTableWithRenderer(provider DataProvider, renderer TableRenderer) *Table {
	widgetTable := widgets.NewTable()

	table := Table{
		Table: widgetTable,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetTable,
		},
	}

	renderer.Table = table
	table.BaseWidget.DataRenderer = renderer

	return &table
}

func NewTable(provider DataProvider) *Table {
	dataRenderer := TableRenderer{}
	table := NewTableWithRenderer(provider, dataRenderer)
	return table
}

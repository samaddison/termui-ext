package dashboard

import (
	termui_ext "termui-ext"
)

func buildGrid(grid Grid, list *[]termui_ext.Refreshable) (*termui_ext.Grid, error) {
	var g = termui_ext.NewGrid()
	block, err := buildBlock(&grid.Common)
	if err != nil {
		return nil, err
	}
	g.Block = *block
	a := build(grid.Items, list)
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	g.Set(b...)
	return g, nil
}

func build(items []Items, list *[]termui_ext.Refreshable) []termui_ext.GridItem {
	gridItems := make([]termui_ext.GridItem, 0)
	for _, v := range items {
		var gridItem termui_ext.GridItem
		widget := findByName(list, v.Widget)
		if v.Type == "column" {
			if widget != nil {
				gridItem = termui_ext.NewCol(v.Ratio, widget)
			} else {
				a := build(v.Items, list)
				b := make([]interface{}, len(a))
				for i := range a {
					b[i] = a[i]
				}
				gridItem = termui_ext.NewCol(v.Ratio, b...)
			}
		} else {
			if widget != nil {
				gridItem = termui_ext.NewRow(v.Ratio, widget)
			} else {
				a := build(v.Items, list)
				b := make([]interface{}, len(a))
				for i := range a {
					b[i] = a[i]
				}
				gridItem = termui_ext.NewRow(v.Ratio, b...)
			}
		}
		gridItems = append(gridItems, gridItem)
	}
	return gridItems
}

func findByName(list *[]termui_ext.Refreshable, name string) termui_ext.Refreshable {
	for _, v := range *list {
		if *v.GetName() == name {
			return v
		}
	}
	return nil
}

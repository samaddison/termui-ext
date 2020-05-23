package dashboard

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"strconv"
	termui_ext "termui-ext"
)

func buildWidgets(widgets []Widgets) (*[]termui_ext.Refreshable, *[]error) {
	list := make([]termui_ext.Refreshable, 0)
	errors := make([]error, 0)
	for _, v := range widgets {
		widget, err := buildWidget(v)
		list = append(list, widget)
		errors = append(errors, err)
	}
	return &list, &errors
}

func buildWidget(widget Widgets) (termui_ext.Refreshable, error) {
	dataProvider, err := buildDataProvider(&widget.Datasource)
	if err != nil {
		return nil, err
	}
	switch {
	case widget.Type == "gauge":
		return buildWidgetGauge(widget, dataProvider)
	case widget.Type == "list":
		return buildWidgetList(widget, dataProvider)
	case widget.Type == "barchart":
		return buildWidgetBarchart(widget, dataProvider)
	case widget.Type == "stacked_barchart":
		return buildWidgetStackedBarchart(widget, dataProvider)
	case widget.Type == "paragraph":
		return buildWidgetParagraph(widget, dataProvider)
	case widget.Type == "piechart":
		return buildWidgetPieChart(widget, dataProvider)
	case widget.Type == "table":
		return buildWidgetTable(widget, dataProvider)
	case widget.Type == "image":
		return buildWidgetImage(widget, dataProvider)
	case widget.Type == "plot":
		return buildWidgetPlot(widget, dataProvider)
	case widget.Type == "sparkline":
		return buildWidgetSparklineGroup(widget, dataProvider)
	}
	return nil, nil
}

func buildWidgetGauge(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.Gauge, error) {
	gauge := termui_ext.NewGaugeDefaultRenderer(provider)
	gauge.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	gauge.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "percent":
			gauge.Percent, _ = strconv.Atoi(v.Value)
		case v.Name == "barColor":
			gauge.BarColor, _ = switchColor(v.Value)
		case v.Name == "label":
			gauge.Label = v.Value
		}
	}
	return gauge, nil
}

func buildWidgetList(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.List, error) {
	list := termui_ext.NewList(provider)
	list.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	list.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "wrapText":
			list.WrapText, _ = strconv.ParseBool(v.Value)
		case v.Name == "selectedRow":
			list.SelectedRow, _ = strconv.Atoi(v.Value)
		}
	}
	return list, nil
}

func buildWidgetBarchart(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.BarChart, error) {
	barChart := termui_ext.NewBarChart(provider)
	barChart.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	barChart.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "barWidth":
			barChart.BarWidth, _ = strconv.Atoi(v.Value)
		case v.Name == "barGap":
			barChart.BarGap, _ = strconv.Atoi(v.Value)
		case v.Name == "maxVal":
			barChart.MaxVal, _ = strconv.ParseFloat(v.Value, 64)
		}
	}
	return barChart, nil
}

func buildWidgetImage(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.Image, error) {
	image := termui_ext.NewImage(provider)
	image.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	image.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "monochrome":
			image.Monochrome, _ = strconv.ParseBool(v.Value)
		case v.Name == "monochromeInvert":
			image.MonochromeInvert, _ = strconv.ParseBool(v.Value)
		}
	}
	return image, nil
}

func buildWidgetParagraph(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.Paragraph, error) {
	paragraph := termui_ext.NewParagraph(provider)
	paragraph.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	paragraph.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "wrapText":
			paragraph.WrapText, _ = strconv.ParseBool(v.Value)
		}
	}
	return paragraph, nil
}

func buildWidgetPieChart(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.PieChart, error) {
	pieChart := termui_ext.NewPieChart(provider)
	pieChart.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	pieChart.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "angleOffset":
			pieChart.AngleOffset, _ = strconv.ParseFloat(v.Value, 64)
		}
	}
	return pieChart, nil
}

func buildWidgetPlot(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.Plot, error) {
	plot := termui_ext.NewPlot(provider)
	plot.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	plot.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "maxVal":
			plot.MaxVal, _ = strconv.ParseFloat(v.Value, 64)
		case v.Name == "showAxes":
			plot.ShowAxes, _ = strconv.ParseBool(v.Value)
		case v.Name == "horizontalScale":
			plot.HorizontalScale, _ = strconv.Atoi(v.Value)
		case v.Name == "plotType":
			plot.PlotType = func(v string) widgets.PlotType {
				if v == "LineChart" {
					return widgets.LineChart
				} else {
					return widgets.ScatterPlot
				}
			}(v.Value)

		case v.Name == "marker":
			plot.Marker = func(v string) widgets.PlotMarker {
				if v == "MarkerBraille" {
					return widgets.MarkerBraille
				} else {
					return widgets.MarkerDot
				}
			}(v.Value)
		}
	}
	return plot, nil
}

func buildWidgetSparklineGroup(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.SparklineGroup, error) {
	sparklineGroup := termui_ext.NewSparklineGroup(provider)
	sparklineGroup.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	sparklineGroup.Block = *block
	return sparklineGroup, nil
}

func buildWidgetStackedBarchart(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.StackedBarChart, error) {
	stackedBarChart := termui_ext.NewStackedBarChart(provider)
	stackedBarChart.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	stackedBarChart.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "barWidth":
			stackedBarChart.BarWidth, _ = strconv.Atoi(v.Value)
		case v.Name == "barGap":
			stackedBarChart.BarGap, _ = strconv.Atoi(v.Value)
		case v.Name == "maxVal":
			stackedBarChart.MaxVal, _ = strconv.ParseFloat(v.Value, 64)
		}
	}
	return stackedBarChart, nil
}

func buildWidgetTable(widget Widgets, provider termui_ext.DataProvider) (*termui_ext.Table, error) {
	table := termui_ext.NewTable(provider)
	table.Name = widget.ID
	block, err := buildBlock(&widget.Common)
	if err != nil {
		return nil, err
	}
	table.Block = *block
	// Deal with WidgetProperties
	for _, v := range widget.WidgetProperties {
		switch {
		case v.Name == "rowSeparator":
			table.RowSeparator, _ = strconv.ParseBool(v.Value)
		case v.Name == "alignment":
			table.TextAlignment = func(v string) termui.Alignment {
				if v == "center" {
					return termui.AlignCenter
				} else if v == "right" {
					return termui.AlignRight
				} else {
					return termui.AlignLeft
				}
			}(v.Value)
		case v.Name == "fillRow":
			table.FillRow, _ = strconv.ParseBool(v.Value)
		}
	}
	return table, nil
}

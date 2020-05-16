# termui-ext

termui-ext is an extension of the termui library (https://github.com/gizak/termui). It provides a set of datasources and renderers that allow the widgets to pull the data periodically from a data source. This effectively enables a "data pull" model for the widget. Rather than the code pushing data into the widget, widgets can be configured to automatically retrieve data from a data provider.

# Installation

To install this library, run the following:

`go get -u github.com/samaddison/termui-ext`

Quick example:

```go
    dataProvider := termui_ext.File{Path: "./termui-ext/docs/stacked_barchart_input.json"}

	bc := termui_ext.NewStackedBarChart(dataProvider)
	bc.Title = "Stacked Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	bc.GoRefresh(5 * time.Second)
```
	
As you can see, you can use the same API provided by the termui widgets, as the termui_ext widgets are essentially a wrapper around those. The only big differences are the call to GoRefresh, which configures the refresh rate, and the introduction of a data provider which needs to be passed in the constructor. 

The library provides 2 data providers: an HTTP/S data provider and a File data provider. With the HTTP data provider, you can configure a url endpoint which the widgets will call periodically. The endpoint will return a JSON string whose format depends on the widget type.

The File data provider reads the JSON from a file. It is probably more of a proof of concept than anything else, but it is certainly usable if you have an application that pushes data periodically into a file.

You can of course create your own data provider. See the source for both the http and file providers for an example. It is actually quite simple. You can for instance create a data provider that will retrieve data from a MySQL database, a time series database like InfluxDB or any other source you can think of.

# Refreshing the Data
After you create and configure a widget, you have to call the Refresh method. You can call the Refresh method directly as a goroutine:

`go bc.Refresh(5 * time.Second)`

or you can call

`bc.GoRefresh(5 * time.Second)`

which will automatically create an internal goroutine to refresh the widget every 5 seconds.

Typically, widgets will run on its own goroutine. This allows you to set up the widgets and let them periodically refresh themselves.

However, you can also use the OneRefresh method which will refresh the widget just once. This allows you to take full control of the widgets.

`bc.OneRefresh()`

You can stop a widget at any time:

`bc.Shutdown()`   

# Data Sources

Data Sources know where the new data source is, connect to the source and retrieve the data. Typically, data would be retrieved from an API, so all you have to do is to configure the widget with the API endpoint and let the data source handle it.



# Renderers
Renderers are the components that receive the JSON from the data source, parse it and refresh the widget data with it. As such, they are specific to each of the widgets, unlike data sources that are general and can be used by any widget.

Renderers are less likely to change, but you can still create your own renderer and assign it to a widget:

// Code


# Hooks
It is unlikely however that existing APIs will return the JSON data in exactly the format required by the widget. For that purpose, Data Sources provide hooks, which you can use to transform the data into a format suitable for the widget. 

Both data sources and renderers provide hooks. If you have multiple widgets of different types which might consume the same data source, it might make sense to use the data source hooks, as you will use the same data source for all the widgets, but you are unlikely to use the same renderer for multiple widget types.

# Licence
MIT




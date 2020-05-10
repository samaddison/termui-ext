package termui_ext

type DataProvider interface {
	PreRetrieve() error
	Retrieve() (*WidgetData, error)
	PostRetrieve(*WidgetData, error) error
}

package dashboard

type Definition struct {
	//Common  Common    `json:"common"`
	Widgets []Widgets `json:"widgets"`
	Grid    Grid      `json:"grid"`
}
type Style struct {
	Fg       string `json:"fg"`
	Bg       string `json:"bg"`
	Modifier string `json:"modifier"`
}
type Title struct {
	Text  string `json:"text"`
	Style Style  `json:"style"`
}
type Border struct {
	Display bool  `json:"display"`
	Left    bool  `json:"left"`
	Right   bool  `json:"right"`
	Top     bool  `json:"top"`
	Bottom  bool  `json:"bottom"`
	Style   Style `json:"style"`
}
type Padding struct {
	Left   int `json:"left"`
	Right  int `json:"right"`
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
}
type Common struct {
	Title   Title   `json:"title"`
	Border  Border  `json:"border"`
	Padding Padding `json:"padding"`
}

type Datasource struct {
	Type string `json:"type"`
	URL  string `json:"url"`
	Path string `json:"path"`
}
type WidgetProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Widgets struct {
	ID               string             `json:"id"`
	Type             string             `json:"type"`
	Common           Common             `json:"common"`
	Datasource       Datasource         `json:"datasource"`
	WidgetProperties []WidgetProperties `json:"widget_properties,omitempty"`
}

type Items struct {
	Type   string  `json:"type"`
	Ratio  float64 `json:"ratio"`
	Widget string  `json:"widget"`
	Items  []Items `json:"items"`
}
type Grid struct {
	ID     string  `json:"id"`
	Common Common  `json:"common"`
	Items  []Items `json:"items"`
}

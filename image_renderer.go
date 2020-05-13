package termui_ext

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

type ImageRenderer struct {
	Image Image
}

type imageData struct {
	Url string
}

func processImageWidgetData(data *WidgetData) (*imageData, error) {
	var chartData = imageData{}
	err := json.Unmarshal(data.Json, &chartData)
	if err != nil {
		return nil, err
	} else {
		return &chartData, nil
	}
}

func (renderer ImageRenderer) PreRender(data *WidgetData) error {
	return nil
}

func (renderer ImageRenderer) Render(data *WidgetData) error {
	chartData, err := processImageWidgetData(data)
	if err != nil {
		return err
	}
	resp, err := http.Get(chartData.Url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	imageBytes, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}
	renderer.Image.Image.Image = imageBytes
	return nil
}

func (renderer ImageRenderer) PostRender(data *WidgetData) error {
	return nil
}

package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Image struct {
	*widgets.Image
	BaseWidget
}

func NewImageWithRenderer(provider DataProvider, renderer ImageRenderer) *Image {
	widgetImage := widgets.NewImage(nil)

	image := Image{
		Image: widgetImage,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetImage,
		},
	}
	renderer.Image = image
	image.BaseWidget.DataRenderer = renderer

	return &image
}

func NewImage(provider DataProvider) *Image {
	dataRenderer := ImageRenderer{}
	image := NewImageWithRenderer(provider, dataRenderer)
	return image
}

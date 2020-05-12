package termui_ext

import "github.com/gizak/termui/v3/widgets"

type Paragraph struct {
	*widgets.Paragraph
	BaseWidget
}

func NewParagraphWithRenderer(provider DataProvider, renderer ParagraphRenderer) *Paragraph {
	widgetParagraph := widgets.NewParagraph()

	paragraph := Paragraph{
		Paragraph: widgetParagraph,
		BaseWidget: BaseWidget{
			DataProvider: provider,
			Drawable:     widgetParagraph,
		},
	}

	renderer.Paragraph = paragraph
	paragraph.BaseWidget.DataRenderer = renderer

	return &paragraph
}

func NewParagraph(provider DataProvider) *Paragraph {
	dataRenderer := ParagraphRenderer{}
	paragraph := NewParagraphWithRenderer(provider, dataRenderer)
	return paragraph
}

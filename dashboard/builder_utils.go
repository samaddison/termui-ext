package dashboard

import (
	"errors"
	"github.com/gizak/termui/v3"
	termui_ext "termui-ext"
)

func buildDataProvider(datasource *Datasource) (termui_ext.DataProvider, error) {
	if datasource.Type == "http" {
		dp := termui_ext.HTTP{Url: datasource.URL}
		return dp, nil
	} else if datasource.Type == "file" {
		dp := termui_ext.File{Path: datasource.Path}
		return dp, nil
	} else {
		return nil, errors.New("invalid datasource type: " + datasource.Type)
	}
}

func buildPadding(padding *Padding, block *termui.Block) {
	block.PaddingBottom = padding.Bottom
	block.PaddingLeft = padding.Left
	block.PaddingRight = padding.Right
	block.PaddingTop = padding.Top
}

func buildBorder(border *Border, block *termui.Block) error {
	block.Border = border.Display
	block.BorderBottom = border.Bottom
	block.BorderLeft = border.Left
	block.BorderTop = border.Bottom
	block.BorderRight = border.Right
	borderStyle, err := buildWidgetStyle(border.Style)
	if err != nil {
		return err
	}
	block.BorderStyle = *borderStyle
	return nil
}

func buildTitle(title *Title, block *termui.Block) error {
	block.Title = title.Text
	style, err := buildWidgetStyle(title.Style)
	if err != nil {
		return err
	}
	block.TitleStyle = *style
	return nil
}

func buildBlock(common *Common) (*termui.Block, error) {
	block := termui.NewBlock()
	err := buildTitle(&common.Title, block)
	if err != nil {
		return nil, err
	}
	err = buildBorder(&common.Border, block)
	if err != nil {
		return nil, err
	}
	buildPadding(&common.Padding, block)
	return block, nil
}

func buildWidgetStyle(style Style) (*termui.Style, error) {
	widgetStyle := termui.Style{}
	var bg, err = switchColor(style.Bg)
	if err != nil {
		return nil, err
	}
	widgetStyle.Bg = bg
	fg, err := switchColor(style.Fg)
	if err != nil {
		return nil, err
	}
	widgetStyle.Fg = fg

	modifier, err := switchModifier(style.Modifier)
	if err != nil {
		return nil, err
	}
	widgetStyle.Modifier = modifier

	return &widgetStyle, nil
}

func switchModifier(modifierName string) (termui.Modifier, error) {
	var modifier termui.Modifier
	switch modifierName {
	case "bold":
		modifier = termui.ModifierBold
	case "underline":
		modifier = termui.ModifierUnderline
	case "reverse":
		modifier = termui.ModifierReverse
	case "":
		modifier = termui.ModifierClear
	default:
		return modifier, errors.New("invalid modifier: " + modifierName)
	}
	return modifier, nil
}

func switchColor(colorName string) (termui.Color, error) {
	var color termui.Color
	switch colorName {
	case "black":
		color = termui.ColorBlack
	case "red":
		color = termui.ColorRed
	case "green":
		color = termui.ColorGreen
	case "yellow":
		color = termui.ColorYellow
	case "blue":
		color = termui.ColorBlue
	case "magenta":
		color = termui.ColorMagenta
	case "cyan":
		color = termui.ColorCyan
	case "white":
		color = termui.ColorRed
	case "":
		color = termui.ColorClear
	default:
		return color, errors.New("invalid color: " + colorName)
	}
	return color, nil
}

package termui_ext

import (
	"io/ioutil"
	"time"
)

type File struct {
	Path string
}

func (dp File) PreRetrieve() error {
	return nil
}

func (dp File) Retrieve() (*WidgetData, error) {
	dat, err := ioutil.ReadFile(dp.Path)
	if err != nil {
		return nil, err
	} else {
		return &WidgetData{
			Time: time.Now(),
			Json: dat,
		}, nil
	}
}

func (dp File) PostRetrieve(*WidgetData, error) error{
	return nil
}

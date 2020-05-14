package termui_ext

import (
	"io/ioutil"
	"net/http"
	"time"
)

type HTTP struct {
	Url string
}

func (dp HTTP) PreRetrieve() {

}

func (dp HTTP) Retrieve() (*WidgetData, error) {
	resp, err := http.Get(dp.Url)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			return &WidgetData{
				Time: time.Now(),
				Json: body,
			}, nil
		}
	}
}

func (dp HTTP) PostRetrieve(*WidgetData, error) {

}

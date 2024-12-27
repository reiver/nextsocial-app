package searchpage

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
	fynewidget    "fyne.io/fyne/v2/widget"
)

type Search struct {
	fynewidget.BaseWidget
}

var _ fyne.CanvasObject = &Search{}
var _ fyne.Widget       = &Search{}

func NewSearch() *Search {
	var result Search
	return &result
}

func (receiver *Search) CreateRenderer() fyne.WidgetRenderer {
	if nil == receiver {
		var nada fyne.WidgetRenderer
		return nada
	}

	var label *fynewidget.Label = fynewidget.NewLabel("dorood")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{
		Bold:true,
		Italic:true,
	}

	var containerptr **fyne.Container

	var input *fynewidget.Entry = fynewidget.NewEntry()
	input.SetPlaceHolder("search...")
	searchFunc := func(str string){
		fmt.Println("searching:", input.Text)

		if nil == containerptr {
			return
		}

		cntnr := *containerptr
		if nil == cntnr {
			return
		}

		if len(cntnr.Objects) <= 0 {
			return
		}

		cntnr.Objects[0] = newProfile()
	}
	input.OnSubmitted = searchFunc

	var container *fyne.Container = fynecontainer.NewBorder(nil, input, nil, nil, label)
	containerptr = &container
	var canvasObject fyne.CanvasObject = container

	return fynewidget.NewSimpleRenderer(canvasObject)
}

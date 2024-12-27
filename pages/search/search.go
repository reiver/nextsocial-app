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

	var input *fynewidget.Entry = fynewidget.NewEntry()
	input.SetPlaceHolder("search...")
	input.OnSubmitted = func(str string){
		fmt.Println("searching:", input.Text)
	}

	var container *fyne.Container = fynecontainer.NewBorder(nil, input, nil, nil, label)
	var canvasObject fyne.CanvasObject = container

	return fynewidget.NewSimpleRenderer(canvasObject)
}

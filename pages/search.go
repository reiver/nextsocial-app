package pages

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynewidget "fyne.io/fyne/v2/widget"
)

func Search(show func()) fyne.CanvasObject {

	var input *fynewidget.Entry = fynewidget.NewEntry()
	input.SetPlaceHolder("search...")
	input.OnSubmitted = func(str string){
		fmt.Println("searching:", input.Text)
		show()
	}

	return input
}

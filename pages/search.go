package pages

import (
	"fyne.io/fyne/v2"
	"github.com/reiver/nextsocial-app/pages/search"
)

func Search() fyne.CanvasObject {
	return searchpage.NewSearch()
}

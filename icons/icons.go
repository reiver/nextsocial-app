package icons

import (
	"fyne.io/fyne/v2"
	fynetheme "fyne.io/fyne/v2/theme"
)

func Account() fyne.Resource {
	return safeFyneIconLookup(fynetheme.IconNameAccount)
}

func Home() fyne.Resource {
	return safeFyneIconLookup(fynetheme.IconNameHome)
}

func Menu() fyne.Resource {
	return safeFyneIconLookup(fynetheme.IconNameMenu)
}

func Notifications() fyne.Resource {
	return safeFyneIconLookup(fynetheme.IconNameHistory)
}

func Search() fyne.Resource {
	return fynetheme.SearchIcon()
}

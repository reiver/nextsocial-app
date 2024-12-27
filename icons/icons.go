package icons

import (
	"fyne.io/fyne/v2"
	fynetheme "fyne.io/fyne/v2/theme"
)

func Account() fyne.Resource {
	return fynetheme.AccountIcon()
}

func Home() fyne.Resource {
	return fynetheme.HomeIcon()
}

func Menu() fyne.Resource {
	return fynetheme.MenuIcon()
}

func Notifications() fyne.Resource {
	return safeFyneIconLookup(fynetheme.IconNameHistory)
}

func Search() fyne.Resource {
	return fynetheme.SearchIcon()
}

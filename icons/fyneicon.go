package icons

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynetheme "fyne.io/fyne/v2/theme"
	"github.com/reiver/go-erorr"
)

var (
	fallBackFyneIcon = &fyne.StaticResource{}
)

func safeFyneIconLookup(name fyne.ThemeIconName) fyne.Resource {

	var icon fyne.Resource = func() fyne.Resource {

		var currentTheme fyne.Theme = fynetheme.Current()
		if nil == currentTheme {
			var err error = erorr.Errorf("could not get fyne current theme when trying to get icon named %q", string(name))
			fyne.LogError("fyne theme.Current() returned nil", err)
			return fallBackFyneIcon
		}

		var ico fyne.Resource = currentTheme.Icon(name)
		if nil == ico {
			var err error = erorr.Errorf("could not load icon named %q from fyne current theme", string(name))
			fyne.LogError(fmt.Sprintf("fyne theme.Current().Icon(%q) returned nil", name), err)
			ico = fallBackFyneIcon
		}

		return ico
	}()

	return icon
}

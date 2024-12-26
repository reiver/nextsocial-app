package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	fyneapp       "fyne.io/fyne/v2/app"
	fynecontainer "fyne.io/fyne/v2/container"
	fynetheme     "fyne.io/fyne/v2/theme"
	fynewidget    "fyne.io/fyne/v2/widget"

	"github.com/reiver/nextsocial-app/icons"
)

func main() {
	fmt.Println("NextSocial 💬")

	var app fyne.App = fyneapp.New()

	var window fyne.Window = app.NewWindow("NextSocial")
	window.Resize(fyne.NewSize(640,480))

	var content *fyne.Container
	{
		var label *fynewidget.Label = fynewidget.NewLabel("NextSocial 💬")

		var toolbar *fynewidget.Toolbar = fynewidget.NewToolbar(
			fynewidget.NewToolbarAction(icons.Account(), func(){
				fmt.Println("account-button activated")
			}),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Search(), func(){
				fmt.Println("search-button activated")
			}),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Menu(), func(){
				fmt.Println("menu-button activated")
			}),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Notifications(), func(){
				fmt.Println("notifications-button activated")
			}),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Home(), func(){
				fmt.Println("home-button activated")
			}),
			fynewidget.NewToolbarSeparator(),
			fynewidget.NewToolbarAction(fynetheme.DocumentCreateIcon(), func() {
				fmt.Println("New document")
			}),
		)

		content = fynecontainer.NewBorder(nil, toolbar, nil, nil, label)
	}

	window.SetContent(content)
	window.ShowAndRun()
}

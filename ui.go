package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
	fynetheme     "fyne.io/fyne/v2/theme"
	fynewidget    "fyne.io/fyne/v2/widget"

	"github.com/reiver/nextsocial-app/icons"
	"github.com/reiver/nextsocial-app/pages"
)

type UI struct {
	container *fyne.Container
	pageAccount       fyne.CanvasObject
	pageHome          fyne.CanvasObject
	pageMenu          fyne.CanvasObject
	pageNew           fyne.CanvasObject
	pageNotifications fyne.CanvasObject
	pageSearch        fyne.CanvasObject
}

func (receiver *UI) Content() fyne.CanvasObject {
	return receiver.container
}

func (receiver *UI) Init() {
	if nil == receiver {
		return
	}

	receiver.pageAccount       = pages.Account()
	receiver.pageHome          = pages.Home()
	receiver.pageMenu          = pages.Menu()
	receiver.pageNew           = pages.New()
	receiver.pageNotifications = pages.Notifications()
	receiver.pageSearch        = pages.Search()

	var content *fyne.Container
	{
		var label *fynewidget.Label = fynewidget.NewLabel("NextSocial 💬")

		var toolbar *fynewidget.Toolbar = fynewidget.NewToolbar(
			fynewidget.NewToolbarAction(icons.Account(), receiver.ShowPageAccount),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Search(), receiver.ShowPageSearch),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Menu(), receiver.ShowPageMenu),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Notifications(), receiver.ShowPageNotifications),
			fynewidget.NewToolbarSpacer(),
			fynewidget.NewToolbarAction(icons.Home(), receiver.ShowPageHome),
			fynewidget.NewToolbarSeparator(),
			fynewidget.NewToolbarAction(fynetheme.DocumentCreateIcon(), receiver.ShowPageNew),
		)


		content = fynecontainer.NewBorder(nil, toolbar, nil, nil, label)
	}

	receiver.container = content
}

func (receiver *UI) show(canvasObject fyne.CanvasObject) {
	if nil == receiver {
		return
	}

	if nil == canvasObject {
		return
	}

	container := receiver.container
	if nil == container {
		return
	}

	if len(container.Objects) <= 0 {
		return
	}

	container.Objects[0] = canvasObject
}

func (receiver *UI) ShowPageAccount() {
	fmt.Println("account-button activated")
	receiver.show(receiver.pageAccount)
}

func (receiver *UI) ShowPageHome() {
	fmt.Println("home-button activated")
	receiver.show(receiver.pageHome)
}

func (receiver *UI) ShowPageMenu() {
	fmt.Println("menu-button activated")
	receiver.show(receiver.pageMenu)
}

func (receiver *UI) ShowPageNew() {
	fmt.Println("new-button activated")
	receiver.show(receiver.pageNew)
}

func (receiver *UI) ShowPageNotifications() {
	fmt.Println("notifications-button activated")
	receiver.show(receiver.pageNotifications)
}

func (receiver *UI) ShowPageSearch() {
	fmt.Println("search-button activated")
	receiver.show(receiver.pageSearch)
}

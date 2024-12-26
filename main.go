package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("NextSocial 💬")

	var app fyne.App = fyneapp.New()

	var window fyne.Window = app.NewWindow("NextSocial")
	window.Resize(fyne.NewSize(640,480))

	var ui UI
	ui.Init()

	var content fyne.CanvasObject = ui.Content()

	window.SetContent(content)
	window.ShowAndRun()
}

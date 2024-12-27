package pages

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
	fynewidget    "fyne.io/fyne/v2/widget"
)

func Profile() fyne.CanvasObject {

	var richtext *fynewidget.RichText
	{
		const name string = "Joe Blow"
		const handle string = "@joe123"
		const summary string = "I like to eat, eat, eat, apples and bananas."

		var markdown string = fmt.Sprintf("# %s (%s)\n\n%s\n", name, handle, summary)

		richtext = fynewidget.NewRichTextFromMarkdown(markdown)
	}

	var scroll *fynecontainer.Scroll = fynecontainer.NewVScroll(richtext)

	return scroll
}

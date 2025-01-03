package searchpage

import (
	"fmt"

	"fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
	fynewidget    "fyne.io/fyne/v2/widget"
)

func newProfile(handle string) fyne.CanvasObject {

	var richtext *fynewidget.RichText
	{
		const name string = "Joe Blow"
		const summary string = "I like to eat, eat, eat, apples and bananas."

		var markdown string = fmt.Sprintf("# %s (%s)\n\n%s\n", name, handle, summary)

		richtext = fynewidget.NewRichTextFromMarkdown(markdown)
	}

	var scroll *fynecontainer.Scroll = fynecontainer.NewVScroll(richtext)

	return scroll
}

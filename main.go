package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xWidget "fyne.io/x/fyne/widget"
	"github.com/tidwall/cities"
)

func main() {

	a := app.New()
	w := a.NewWindow("Entry Completion")

	entry := xWidget.NewCompletionEntry([]string{})

	// When the use typed text, complete the list.
	cardTexts := []string{}

	entry.OnChanged = func(s string) {

		results := []cities.City{}

		for _, value := range cities.Cities {

			if len(value.City) < len(s) {
				continue
			}

			if strings.EqualFold(s, value.City[:len(s)]) {
				results = append(results, value)
			}
		}

		if len(results) == 0 {
			entry.HideCompletion()
			return
		}

		cardTexts = []string{}
		for _, r := range results {
			s := r.City
			cardTexts = append(cardTexts, s)
		}

		entry.SetOptions(cardTexts)
		entry.ShowCompletion()
	}

	button := widget.NewButton("Simulate fast typing", func() {

		for {
			entry.SetText("")
			entry.SetText(entry.Text + "edin")
		}

	})

	w.SetContent(container.NewVBox(button, entry))

	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()

}

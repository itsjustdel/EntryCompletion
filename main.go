package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/x/fyne/widget"
)

func main() {

	entry := widget.NewCompletionEntry([]string{})

	// When the use typed text, complete the list.
	cardTexts := []string{}

	entry.OnChanged = func(s string) {

		// completion start for text length >= 2 Some cities have two letter names
		if len(s) < 2 {
			entry.HideCompletion()
			return
		}

		results := []City{}

		for _, value := range Cities {

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
			//			timezone := timezonemapper.LatLngToTimezoneString(r.Latitude, r.Longitude)

			s := r.City + "--" + r.Country + "--" // + timezone
			cardTexts = append(cardTexts, s)
		}

		// then show them
		entry.SetOptions(cardTexts)
		entry.ShowCompletion()
	}

	a := app.New()
	w := a.NewWindow("Entry Completion")

	w.SetContent(entry)

	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()

}

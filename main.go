package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"time"
)

const WINDOW_TITLE = "Clock"

func main() {
	a := app.New()
	w := a.NewWindow(WINDOW_TITLE)

	kitchenTime := widget.NewLabel("")

	updatingTime(kitchenTime)

	w.SetContent(kitchenTime)

	go func() {
		for range time.Tick(1000 * time.Millisecond) {
			updatingTime(kitchenTime)
		}
	}()

	w.ShowAndRun()
}

func updatingTime(clock *widget.Label) {
	format := "3:04:05 PM"
	formatted := time.Now().Format(format)

	clock.SetText(formatted)
}

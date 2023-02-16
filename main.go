package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
)

const WINDOW_TITLE = "Clock"

func init() {
	// set scale environment variable
	os.Setenv("FYNE_SCALE", "3.5")
}

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

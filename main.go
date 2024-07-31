package main

import (
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fyneApp := app.NewWithID("TestApp")

	win := fyneApp.NewWindow("TestWindow")
	win.Resize(fyne.NewSize(500, 500))
	boundString := binding.NewTree()
	win.SetContent(
		container.NewVBox(
			widget.NewEntryWithData(boundString),
			widget.NewButton("change", func() {
				_ = boundString.Set("1")
				time.Sleep(10 * time.Millisecond)
				_ = boundString.Set("2")
			}),
		),
	)
	win.Show()

	fyneApp.Run()
}

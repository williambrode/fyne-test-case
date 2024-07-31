package main

import (
	"fmt"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("starting")
	fyneApp := app.NewWithID("TestApp")

	win := fyneApp.NewWindow("TestWindow")
	win.Resize(fyne.NewSize(500, 500))
	ids := map[string][]string{"": {"a"}}
	values := map[string]string{"a": "a"}
	treeBinding := binding.NewStringTree()
	treeBinding.Set(ids, values)
	tree := widget.NewTreeWithData(treeBinding,
		// Create Item
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		// Update Item
		func(di binding.DataItem, branch bool, co fyne.CanvasObject) {
			if strBinding, ok := di.(binding.String); !ok {
				fmt.Println("fail")
			} else if val, err := strBinding.Get(); err != nil {
				fmt.Println("fail")
			} else {
				co.(*widget.Label).SetText(val)
			}
		},
	)
	win.SetContent(
		tree,
	)
	win.Show()

	// Update only the value
	go func() {
		time.Sleep(5 * time.Second)
		values = map[string]string{"a": "b"}
		treeBinding.Set(ids, values)
	}()

	fyneApp.Run()
	fmt.Println("ending")
}

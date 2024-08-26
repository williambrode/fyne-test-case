package main

import (
	"fmt"
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	fmt.Println("starting")
	fyneApp := app.NewWithID("TestApp")

	win := fyneApp.NewWindow("TestWindow")
	win.Resize(fyne.NewSize(500, 500))
	ids := map[string][]string{"": {"1"}}
	values := map[string]string{"1": "1"}
	treeBinding := binding.NewStringTree()
	treeBinding.Set(ids, values)
	nextValue := 2
	treeNode := binding.NewInt()
	tree := widget.NewTreeWithData(treeBinding,
		// Create Item
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		// Update Item
		func(di binding.DataItem, branch bool, co fyne.CanvasObject) {
			if strBinding, ok := di.(binding.String); !ok {
				fmt.Println("fail")
			} else {
				co.(*widget.Label).Bind(strBinding)
			}
		},
	)
	treeNodeLabel := widget.NewLabel("")
	treeNode.AddListener(binding.NewDataListener(func() {
		val, _ := treeNode.Get()
		if val == 0 {
			treeNodeLabel.SetText("Root")
			return
		}
		treeNodeLabel.SetText(strconv.Itoa(val))
	}))
	addButton := widget.NewButton("Add", func() {
		val, _ := treeNode.Get()
		treeNode := strconv.Itoa(val)
		newValue := strconv.Itoa(nextValue)
		nextValue++
		if val == 0 {
			ids[""] = append(ids[""], newValue)
		} else {
			ids[treeNode] = append(ids[treeNode], newValue)
		}
		values[newValue] = newValue
		treeBinding.Set(ids, values)
	})
	reproButton := widget.NewButton("Repro", func() {
		ids[""] = append(ids[""], "a")
		ids["a"] = append(ids["a"], "b")
		ids["b"] = append(ids["b"], "c")
		values["a"] = "a"
		values["b"] = "b"
		values["c"] = "c"
		treeBinding.Set(ids, values)
	})
	upButton := widget.NewButton("Up", func() {
		val, _ := treeNode.Get()
		if val > 0 {
			treeNode.Set(val - 1)
		}
	})
	downButton := widget.NewButton("Down", func() {
		val, _ := treeNode.Get()
		treeNode.Set(val + 1)
	})
	selectButton := widget.NewButton("Select", func() {
		tree.Select("c")
	})
	win.SetContent(
		container.NewBorder(container.NewHBox(treeNodeLabel, reproButton, selectButton, addButton, upButton, downButton), nil, nil, nil, tree),
	)
	win.Show()
	fyneApp.Run()
	fmt.Println("ending")
}

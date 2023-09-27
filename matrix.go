package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showMatrixScreen(window fyne.Window, columns, rows int) {
	mainContent := container.New(layout.NewGridLayoutWithColumns(columns))
	for x := 0; x < columns; x++ {
		for y := 0; y < rows; y++ {
			input := widget.NewEntry()
			input.SetPlaceHolder("0")
			mainContent.Add(input)	
		}
	}

	window.SetContent(mainContent)
}

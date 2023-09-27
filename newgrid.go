package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func showNewGridScreen(window fyne.Window) {
	rowsEntry := widget.NewEntry()
	rowsEntry.SetPlaceHolder("Enter the number of rows")

	colsEntry := widget.NewEntry()
	colsEntry.SetPlaceHolder("Enter the number of columns")

	submitButton := widget.NewButton("Submit", func() {
		rows_input := rowsEntry.Text
		cols_input := colsEntry.Text

		// Validate the input
		if rows_input == "" || cols_input == "" {
			dialog.ShowError(fmt.Errorf("both rows and columns must be entered"), window)
			return
		}
		columns, err := strconv.Atoi(cols_input)
		if err != nil {
			dialog.ShowError(fmt.Errorf("columns must be a number"), window)
			return
		}
		rows, err := strconv.Atoi(rows_input)
		if err != nil {
			dialog.ShowError(fmt.Errorf("columns must be a number"), window)
			return
		}

		showMatrixScreen(window, columns, rows)
	})

	content := container.NewVBox(
		widget.NewLabel("Enter the size of the matrix:"),
		rowsEntry,
		colsEntry,
		submitButton,
	)

	window.SetContent(content)
}

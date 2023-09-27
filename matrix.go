package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var currentMatrix [][]uint16

func showMatrixScreen(window fyne.Window, columns, rows int) {
	var entriesMatrix [][]*widget.Entry

	// if currentMatrix is empty, it means we are working with a new grid
	// initialize it empty
	if len(currentMatrix) == 0 {
		for y := 0; y < rows; y++ {
			currentRow := make([]uint16, columns)
			currentMatrix = append(currentMatrix, currentRow)
		}
	}

	mainGrid := container.New(layout.NewGridLayoutWithColumns(columns))
	for y := 0; y < rows; y++ {
		currentRow := make([]*widget.Entry, columns)
		for x := 0; x < columns; x++ {
			input := widget.NewEntry()
			if currentMatrix[y][x] == 0 {
				// currentMatrix is either empty or value of this box is 0
				input.SetPlaceHolder("0")
			} else {
				// currentMatrix has data
				input.Text = fmt.Sprint(currentMatrix[y][x])
			}
			currentRow[x] = input
			mainGrid.Add(input)
		}
		entriesMatrix = append(entriesMatrix, currentRow)
	}
	scrollablegrid := container.NewScroll(mainGrid)

	resetButton := widget.NewButton("Reset Matrix", func() {
		showMatrixScreen(window, columns, rows)
	})

	previewButton := widget.NewButton("Preview", func() {

	})

	validateButton := widget.NewButton("Generate JSON", func() {
		var matrix [][]uint16
		for y := 0; y < rows; y++ {
			row := make([]uint16, columns)
			for x := 0; x < columns; x++ {
				value := entriesMatrix[y][x].Text
				if value == "" {
					value = "0"
				}
				intValue, err := strconv.Atoi(value)
				if err != nil {
					dialog.ShowError(fmt.Errorf("unable to convert value %d/%d %s to uint16", x, y, value), window)
				}
				row[x] = uint16(intValue)
			}
			matrix = append(matrix, row)
		}
		jsonData, err := json.Marshal(matrix)
		if err != nil {
			dialog.ShowError(fmt.Errorf("error encoding JSON: %w", err), window)
			return
		}
		data := fmt.Sprintf("{\"MapMatrix\": %s}", string(jsonData))
		ShowSaveGridScreen(window, data)
	})
	lastLine := container.New(layout.NewGridLayoutWithColumns(3), resetButton, previewButton, validateButton)

	mainContent := container.NewBorder(nil, lastLine, nil, nil, scrollablegrid)

	window.SetContent(mainContent)
}

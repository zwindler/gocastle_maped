package main

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showMatrixScreen(window fyne.Window, columns, rows int) {
	var entriesMatrix [][]*widget.Entry

	mainGrid := container.New(layout.NewGridLayoutWithColumns(columns))
	for y := 0; y < rows; y++ {
		currentRow := make([]*widget.Entry, columns)
		for x := 0; x < columns; x++ {
			input := widget.NewEntry()
			input.SetPlaceHolder("0")
			currentRow[x] = input
			mainGrid.Add(input)
		}
		entriesMatrix = append(entriesMatrix, currentRow)
	}

	resetButton := widget.NewButton("Reset Matrix", func() {
		showMatrixScreen(window, columns, rows)
	})

	previewButton := widget.NewButton("Preview", func() {

	})

	validateButton := widget.NewButton("Generate JSON", func() {
		var matrix [][]string
		for y := 0; y < rows; y++ {
			row := make([]string, columns)
			for x := 0; x < columns; x++ {
				value := entriesMatrix[y][x].Text
				if value == "" {
					value = "0"
				}
				row[x] = value
			}
			matrix = append(matrix, row)
		}
		jsonData, err := json.Marshal(matrix)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	})
	lastLine := container.New(layout.NewGridLayoutWithColumns(3), resetButton, previewButton, validateButton)

	mainContent := container.NewBorder(nil, lastLine, nil, nil, mainGrid)

	window.SetContent(mainContent)
}

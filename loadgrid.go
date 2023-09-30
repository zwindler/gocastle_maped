package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func ShowLoadGridScreen(window fyne.Window) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if reader == nil {
			return
		}

		defer reader.Close()

		currentMatrix, err = loadGridFromFile(reader)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		rows := len(currentMatrix)
		columns := len(currentMatrix[0])
		showMatrixScreen(window, columns, rows)
	}, window)
	// only show .json files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	location, err := getBaseDirectory()

	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)
	fd.Show()
}

func loadGridFromFile(r io.Reader) (data [][]uint16, err error) {
	return data, json.NewDecoder(r).Decode(&data)
}

func getBaseDirectory() (fyne.ListableURI, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	absoluteDirPath := filepath.Dir(executablePath)
	return storage.ListerForURI(storage.NewFileURI(absoluteDirPath))
}

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

		data, err := loadGridFromFile(reader)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		if err := updateLoadedGridData(data); err != nil {
			dialog.ShowError(err, window)
			return
		}
	}, window)
	// only show .sav files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	location, err := getBaseDirectory()

	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)
	fd.Show()
}

func loadGridFromFile(r io.Reader) (data map[string]interface{}, err error) {
	return data, json.NewDecoder(r).Decode(&data)
}

func updateLoadedGridData(data map[string]interface{}) error {
	var loadedData [][]int
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &loadedData)
	if err != nil {
		return err
	}

	return nil
}

func getBaseDirectory() (fyne.ListableURI, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	absoluteDirPath := filepath.Dir(executablePath)
	return storage.ListerForURI(storage.NewFileURI(absoluteDirPath))
}

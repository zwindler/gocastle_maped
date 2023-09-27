package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

// ShowSaveGridScreen is the main function of the save Grid screen.
func ShowSaveGridScreen(window fyne.Window, data string) {

	// Show file save dialog
	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			defer writer.Close()

			// Write Grid data to JSON .sav file
			if _, err := io.WriteString(writer, data); err != nil {
				dialog.ShowError(err, window)
			} else {
				dialog.ShowInformation("Grid Saved", "Grid data has been successfully saved.", window)
			}
		}
	}, window)

	// only allow .json files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))

	location, err := getBaseDirectory()
	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)

	fd.Show()
}

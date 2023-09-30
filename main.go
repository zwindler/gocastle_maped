package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	Maped := app.New()
	Maped.Settings().SetTheme(CustomTheme{})
	if Maped.Settings().ThemeVariant() == 0 {
		// dark theme
		TextColor = color.White
	} else {
		// light theme
		TextColor = color.Black
	}

	window := Maped.NewWindow("GoCastle Maped")

	showMenuScreen(window)
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}

// showMenuScreen is the main function of the main screen.
func showMenuScreen(window fyne.Window) {

	// Create buttons
	newGridButton := widget.NewButton("New Grid", func() {
		showNewGridScreen(window)
	})

	loadGridButton := widget.NewButton("Load Grid", func() {
		ShowLoadGridScreen(window)
	})
	quitButton := widget.NewButton("Quit", func() {
		window.Close()
	})

	buttons := container.New(layout.NewVBoxLayout(),
		newGridButton,
		loadGridButton,
		quitButton,
	)

	menu := container.New(layout.NewCenterLayout(), buttons)

	window.SetContent(menu)
}

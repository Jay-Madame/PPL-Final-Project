package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//non-gui version
/**
func main() {
	StartGame()
}
**/

// test graphics

// to not test this feat
//  feel free to replace with
//  /** these surrounding**/
func main() {
	myApp := app.New()
	window := myApp.NewWindow("Chess with Fyne")

	// Create an 8x8 grid
	grid := container.NewGridWithColumns(8)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			btn := widget.NewButton("", nil)
			btn.Importance = widget.LowImportance // Makes it look flat like a square
			grid.Add(btn)
		}
	}

	window.SetContent(grid)
	window.Resize(fyne.NewSize(400, 400)) // Set window size
	window.ShowAndRun()
}
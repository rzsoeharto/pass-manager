package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	appWindow := myApp.NewWindow("Rizky's Password Manager")

	helloLabel := widget.NewLabel("Hello Fyne!")
	greetButton := widget.NewButton("Greet", func() {
		helloLabel.SetText("Hello Fyne, you clicked the button!")
	})

	appWindow.SetContent(container.NewVBox(
		helloLabel,
		container.NewHBox(
			widget.NewEntry(),
			widget.NewButton("OK", func() {}),
		),
		greetButton,
		layout.NewSpacer(),
	))

	appWindow.ShowAndRun()
}

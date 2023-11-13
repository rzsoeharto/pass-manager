package main

import (
	"fmt"
	"os"
	"pass-manager/database"
	"pass-manager/initialisers"
	"pass-manager/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func init() {
	initialisers.LoadEnv()
	logger.Init(os.Getenv("LOGFILE"))
}

func main() {
	db := database.InitDB()

	err := db.Ping()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ping success, i think")

	myApp := app.New()
	appWindow := myApp.NewWindow("Rizky's Password Manager")
	passwordEntry := widget.NewPasswordEntry()
	greetButton := widget.NewButton("Login", func() {
		fmt.Println(passwordEntry.Text)
	})

	loginContent := container.NewVBox(
		widget.NewLabel("Password"),
		passwordEntry,
		greetButton,
	)

	appWindow.Resize(fyne.NewSize(500, 300))
	appWindow.SetFixedSize(true)

	appWindow.SetContent(loginContent)

	appWindow.ShowAndRun()
}

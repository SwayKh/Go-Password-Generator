package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	encryption "github.com/swaykh/Go-Password-generator/pkg/Encryption"
	fileio "github.com/swaykh/Go-Password-generator/pkg/FileIO"
	input "github.com/swaykh/Go-Password-generator/pkg/Input"
	password "github.com/swaykh/Go-Password-generator/pkg/Password"
)

// * Ask for a default password length - Set the default length to 8
// * Add Password saving
// * Adding name or title for saved password
// * Save to json objects to make getting passwords easier
// * Save some additional info like creation time, name, data, password length
// Getting passwords from Saved File
// Deleting saved passwords
// * Encrypting passwords
// Start with prompt to get password or generate one
// Use some prettier terminal library

// TODO: Somehow all passwords get reset to same time every time the password is
// saved
// OR MAYBE NOT, it seems fine

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Generate Password", "Create a random password or specific length", '1', output).
		AddItem("Get Password", "Select a password from saved passwords list", '2', output).
		AddItem("Delete Password", "Delete a password from saved passwords list", '3', output).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	list.SetMainTextColor(tcell.ColorOrange)
	list.SetSecondaryTextColor(tcell.ColorLightGray)
	list.SetSelectedTextColor(tcell.ColorOrange)
	list.SetSelectedBackgroundColor(tcell.ColorBlack)

	// These are BOX properties, which every widget enherits
	list.SetTitle("Password Generator")
	list.SetBorder(true)
	list.SetTitleColor(tcell.ColorBlue)
	list.SetBackgroundColor(tcell.ColorReset)

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func output() {
	fmt.Print("You clicked on a list item")
}

func Start() {
	// GetPassword returns struct array of PasswordDetail
	// Don't know what I should do with the return value
	// Does running GetPassword put value into a struct into the package

	keyForEncryption := []byte("0123456789ABCDEF0123456789ABCDEF")
	passArray := fileio.GetPassword()
	length := input.GetLength()

	var pass *string = password.GeneratePassord(length)
	fmt.Println("Your Password is: ", *pass)

	encryptedPassword := encryption.Encrypt(*pass, keyForEncryption)

	if input.ShouldSaveInfo() {
		passName := input.GetPassName()
		fileio.Save(&passArray, passName, &encryptedPassword, length)
	} else {
		// Do nothing
	}
}

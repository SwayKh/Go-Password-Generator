package Program

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

func Ui() {
	app := tview.NewApplication()
	textArea := tview.NewTextView()

	list := tview.NewList().
		AddItem("Generate Password",
			"Create a random password or specific length",
			'1',
			func() { NewTextViewString(textArea, "You Clicked Generate Password") }).
		AddItem("Get Password",
			"Select a password from saved passwords list",
			'2',
			func() { NewTextViewString(textArea, "You Clicked Get Password") }).
		AddItem("Delete Password",
			"Delete a password from saved passwords list",
			'3',
			func() { NewTextViewString(textArea, "You Clicked Delete Password") }).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	list.SetMainTextColor(tcell.ColorOrange).
		SetSecondaryTextColor(tcell.ColorLightGray).
		SetSelectedTextColor(tcell.ColorOrange).
		SetSelectedBackgroundColor(tcell.ColorBlack)

	root := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(textArea, 0, 1, false).
		SetDirection(tview.FlexColumn)

	list.SetBorder(true).
		SetTitle("Options").
		SetTitleColor(tcell.ColorBlue).
		SetBackgroundColor(tcell.ColorReset)

	textArea.SetBorder(true).
		SetTitle("Output").
		SetTitleColor(tcell.ColorBlue).
		SetBackgroundColor(tcell.ColorReset)

	root.SetBorder(true).
		SetTitle("Password Generator").
		SetTitleColor(tcell.ColorBlue).
		SetBackgroundColor(tcell.ColorReset)

	if err := app.SetRoot(root, true).Run(); err != nil {
		panic(err)
	}
}

func NewTextViewString(textView *tview.TextView, newString string) {
	textView.Clear()
	fmt.Fprint(textView, newString)
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

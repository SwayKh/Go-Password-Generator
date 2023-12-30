package cmd

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

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

func Start() {
	// GetPassword returns struct array of PasswordDetail
	// Don't know what I should do with the return value
	// Does running GetPassword put value into a struct into the package

	passArray := fileio.LoadPasswords()
	length := input.PromptForLength()

	var pass *string = password.GeneratePassord(length)
	fmt.Println("Your Password is: ", *pass)

	if input.ShouldSaveInfo() {
		passName := input.PromptForPasswordName()
		fileio.Save(&passArray, passName, pass, length)
	} else {
		// Do nothing
	}
}

func Ui() {
	app := tview.NewApplication()

	list := tview.NewList().
		AddItem("Generate Password",
			"Create a random password or specific length",
			'1',
			nil,
		).
		AddItem("Get Password",
			"Select a password from saved passwords list",
			'2',
			nil,
		).
		AddItem("Delete Password",
			"Delete a password from saved passwords list",
			'3',
			nil,
		).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	list.SetMainTextColor(tcell.ColorOrange).
		SetSecondaryTextColor(tcell.ColorLightGray).
		SetSelectedTextColor(tcell.ColorOrange).
		SetSelectedBackgroundColor(tcell.ColorBlack)

	list.SetBackgroundColor(tcell.ColorReset)

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}

func Tui() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorReset
	tview.Styles.TitleColor = tcell.ColorBlue
	tview.Styles.BorderColor = tcell.ColorBlue

	app := tview.NewApplication()
	defaultStyle := tcell.StyleDefault
	textArea := tview.NewTextView().SetTextStyle(defaultStyle)

	list := tview.NewList().
		AddItem("Generate Password",
			"Create a random password or specific length",
			'1',
			func() { newTextViewString(textArea, "You Clicked Generate Password") }).
		AddItem("Get Password",
			"Select a password from saved passwords list",
			'2',
			func() { newTextViewString(textArea, "You Clicked Get Password") }).
		AddItem("Delete Password",
			"Delete a password from saved passwords list",
			'3',
			func() { newTextViewString(textArea, "You Clicked Delete Password") }).
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

	addBoxStyles(root.Box, "Password Generator")
	addBoxStyles(textArea.Box, "Output")
	addBoxStyles(list.Box, "Option")

	if err := app.SetRoot(root, true).Run(); err != nil {
		panic(err)
	}
}

func addBoxStyles(widget *tview.Box, title string) {
	widget.SetBorder(true).
		SetTitle(title).
		SetTitleColor(tcell.ColorBlue).
		SetBackgroundColor(tcell.ColorReset)
}

func newTextViewString(textView *tview.TextView, newString string) {
	textView.Clear()
	fmt.Fprint(textView, newString)
}

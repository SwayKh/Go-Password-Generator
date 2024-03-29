package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PromptForPasswordName() *string {
	scanner := bufio.NewScanner(os.Stdin)
	var passName string
	fmt.Print("Enter a name or description of the password: ")
	scanner.Scan()
	passName = scanner.Text()
	return &passName
}

func PromptForPasswordID() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Choose a password number to show the password: ")
	scanner.Scan()
	passID, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid Data Entered!\n", err)
	}
	return passID
}

func GetOldPassword() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var getChoice string
	fmt.Print("Get Saved Password? [y/n]: ")
	scanner.Scan()

	if getChoice == "y" || getChoice == "Y" {
		return true
	} else if getChoice == "n" || getChoice == "N" {
		return false
	} else {
		fmt.Println("Invalid Choice!, Defaulting to No")
		return false
	}
}

func ShouldSaveInfo() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var saveChoice string
	fmt.Print("Save Password? [y/n]: ")
	scanner.Scan()
	saveChoice = scanner.Text()

	if saveChoice == "y" || saveChoice == "Y" {
		return true
	} else if saveChoice == "n" || saveChoice == "N" {
		return false
	} else {
		fmt.Println("Invalid Choice!, Defaulting to No")
		return false
	}
}

func PromptForLength() *int {
	var length int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Password length: ")
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid Data Entered!\n", err)
	}

	// Handles empty values, and non int values since length is int, so default
	// value is 0
	if length == 0 {
		fmt.Println("Setting password length to Default: 8")
		length = 8
	}
	return &length
}

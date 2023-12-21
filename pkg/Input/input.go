package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetPassName() *string {
	scanner := bufio.NewScanner(os.Stdin)
	var passName string
	fmt.Print("Enter a name or description of the password: ")
	scanner.Scan()
	passName = scanner.Text()
	return &passName
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

func GetLength() *int {
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

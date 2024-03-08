package password

import (
	"math/rand"
	"strings"

	encryption "github.com/swaykh/Go-Password-generator/pkg/Encryption"
	fileio "github.com/swaykh/Go-Password-generator/pkg/FileIO"
)

func GeneratePassord(passwordLength *int) *string {
	letterSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~"}

	var password string
	var strBuilder strings.Builder

	for i := 0; i < *passwordLength; i++ {
		randomIndex := rand.Intn(len(letterSlice))
		strBuilder.WriteString(letterSlice[randomIndex])
	}
	password = strBuilder.String()
	return &password
}

func GetPasswordNames(passwordList []fileio.PasswordDetail) []string {
	passwordNames := []string{}
	for i := range passwordList {
		passwordNames = append(passwordNames, passwordList[i].Name)
	}
	return passwordNames
}

func GetPasswordTime(passwordList []fileio.PasswordDetail) []string {
	passwordSaveTime := []string{}
	for i := range passwordList {
		dateTime := passwordList[i].Date + " " + passwordList[i].Time
		passwordSaveTime = append(passwordSaveTime, dateTime)
	}
	return passwordSaveTime
}

func GetPasswordLength(passwordList []fileio.PasswordDetail) []int {
	passwordLengths := []int{}
	for i := range passwordList {
		passwordLengths = append(passwordLengths, passwordList[i].PasswordLength)
	}
	return passwordLengths
}

func GetPasswordDecrypted(passwordList []fileio.PasswordDetail) []string {
	decryptedPasswordSlice := []string{}

	for i := range passwordList {
		decryptedPassword := encryption.Decrypt(passwordList[i].Key, encryption.KeyForEncryption)
		decryptedPasswordSlice = append(decryptedPasswordSlice, decryptedPassword)
	}
	return decryptedPasswordSlice
}

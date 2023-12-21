package generate

import (
	"math/rand"
	"strings"
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

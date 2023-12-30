package fileio

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	encryption "github.com/swaykh/Go-Password-generator/pkg/Encryption"
)

type PasswordDetail struct {
	Name           string `json:"PasswordName"`
	Key            string `json:"PasswordKey"`
	PasswordLength int    `json:"PasswordLength"`
	Date           string `json:"CreationDate"`
	Time           string `json:"CreationTime"`
}

var filename = "./data/password.json"

func LoadPasswords() []PasswordDetail {
	// Get Password
	var passArray []PasswordDetail
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(file), &passArray)
	if err != nil {
		fmt.Println(err)
	}
	return passArray
}

func Save(passArray *[]PasswordDetail, name *string, key *string, length *int) {
	now := time.Now()

	var pass PasswordDetail

	pass.Name = *name
	pass.Key = encryption.Encrypt(*key, encryption.KeyForEncryption)
	pass.PasswordLength = *length

	pass.Date = now.Format("2 Jan")
	pass.Time = now.Format("03:06 PM")

	*passArray = append(*passArray, pass)

	jsonData, err := json.MarshalIndent(*passArray, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	writeFile(string(jsonData))
}

func writeFile(passwordData string) {
	// Since getPassword loads the json data into passArray, os.Create, truncating
	// the file does matter.
	// file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(passwordData + "\n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Password Saved")
}

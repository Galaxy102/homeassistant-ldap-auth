package homeassistant

import (
	"errors"
	"fmt"
	"os"
)

type UsernameAndPassword struct {
	Username string
	Password string
}

type UserData struct {
	DisplayName string
	Group       string
}

func ReadUsernameAndPassword(target *UsernameAndPassword) error {
	username, uSet := os.LookupEnv("username")
	password, pSet := os.LookupEnv("password")
	if !(uSet && pSet) {
		return errors.New("username or password unset")
	}
	target.Username = username
	target.Password = password
	return nil
}

func PrintEntry(data UserData) {
	fmt.Printf("name = %s\n", data.DisplayName)
	fmt.Printf("group = %s\n", data.Group)
}

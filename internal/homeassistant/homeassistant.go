package homeassistant

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"os"
)

type UsernameAndPassword struct {
	Username string
	Password string
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

func PrintEntry(entry ldap.Entry, displayNameAttr string) {
	fmt.Printf("name = %s\n", entry.GetAttributeValue(displayNameAttr))
	fmt.Printf("group = system-users")
}

package main

import (
	"fmt"
	"github.com/Galaxy102/homeassistant-ldap-auth/internal/homeassistant"
	"github.com/Galaxy102/homeassistant-ldap-auth/internal/ldap"
	"log"
)

func main() {
	credentials := homeassistant.UsernameAndPassword{}
	err := homeassistant.ReadUsernameAndPassword(&credentials)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read credentials: %v", err))
	}
	ldapConfig := ldap.Config{}
	err = ldap.ReadLdapConfig(&ldapConfig)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read ldap config: %v", err))
	}

	ldapResult, err := ldap.ConnectAndReadUser(ldapConfig, credentials)
	if err != nil {
		log.Fatal(fmt.Errorf("could not perform ldap auth: %v", err))
	}

	homeassistant.PrintEntry(homeassistant.UserData{
		DisplayName: ldapResult.GetAttributeValue(ldapConfig.DisplayNameAttr),
		Group:       "system-users",
	})
}

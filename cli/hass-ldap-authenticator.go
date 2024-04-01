package main

import (
	"fmt"
	"hass-ldap/internal/homeassistant"
	"hass-ldap/internal/ldap"
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

	homeassistant.PrintEntry(*ldapResult, ldapConfig.DisplayNameAttr)
}

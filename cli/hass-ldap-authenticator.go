package main

import (
	"fmt"
	"github.com/Galaxy102/homeassistant-ldap-auth/internal/homeassistant"
	"github.com/Galaxy102/homeassistant-ldap-auth/internal/ldap"
	"log"
	"os"
)

func main() {
	credentials := homeassistant.UsernameAndPassword{}
	err := homeassistant.ReadUsernameAndPassword(&credentials)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read credentials: %v", err))
	}
	configCli := ldap.NewConfigCli()
	_ = configCli.Flags.Parse(os.Args[1:]) // Errors are handled by flag
	err = configCli.Validate()
	if err != nil {
		configCli.Flags.Usage()
		log.Fatal(fmt.Errorf("could not read ldap config: %v", err))
	}

	ldapResult, err := ldap.ConnectAndReadUser(configCli.LdapConfig, credentials)
	if err != nil {
		log.Fatal(fmt.Errorf("could not perform ldap auth: %v", err))
	}

	group := "system-users"
	if configCli.LdapConfig.AuthenticateAdmin {
		group = "system-admin"
	}
	homeassistant.PrintEntry(homeassistant.UserData{
		DisplayName: ldapResult.GetAttributeValue(configCli.LdapConfig.DisplayNameAttr),
		Group:       group,
	})
}

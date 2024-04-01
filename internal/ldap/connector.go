package ldap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/Galaxy102/homeassistant-ldap-auth/internal/homeassistant"
	ldapClient "github.com/go-ldap/ldap/v3"
)

func ConnectAndReadUser(ldapConfig Config, credentials homeassistant.UsernameAndPassword) (*ldapClient.Entry, error) {
	l, err := ldapClient.DialURL(ldapConfig.Uri)
	if err != nil {
		return nil, fmt.Errorf("could not dial: %v", err)
	}
	defer l.Close()

	// Reconnect with TLS
	if ldapConfig.UseStartTLS {
		err = l.StartTLS(&tls.Config{})
		if err != nil {
			return nil, fmt.Errorf("could not connect with StartTLS: %v", err)
		}
	}

	// First bind with a read only credentials
	if ldapConfig.UseBind {
		err = l.Bind(ldapConfig.Bind.BindUser, ldapConfig.Bind.Password)
		if err != nil {
			return nil, fmt.Errorf("could not bind: %v", err)
		}
	}

	// Search for the given username
	searchRequest := ldapClient.NewSearchRequest(
		ldapConfig.BaseDN,
		ldapClient.ScopeWholeSubtree, ldapClient.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(ldapConfig.UserFilter, ldapClient.EscapeFilter(credentials.Username)),
		[]string{},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("could not search: %v", err)
	}

	if len(sr.Entries) != 1 {
		return nil, errors.New("user does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the credentials to verify their password
	err = l.Bind(userdn, credentials.Password)
	if err != nil {
		return nil, fmt.Errorf("could not bind with found credentials: %v", err)
	}
	return sr.Entries[0], nil
}

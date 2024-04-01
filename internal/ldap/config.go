package ldap

import (
	"errors"
	"os"
	"strings"
)

type BindConfig struct {
	BindUser string
	Password string
}

type Config struct {
	Uri             string
	UseStartTLS     bool
	UseBind         bool
	Bind            BindConfig
	BaseDN          string
	UserFilter      string
	DisplayNameAttr string
}

func ReadLdapConfig(target *Config) error {
	bindUsername, useBind := os.LookupEnv("LDAP_AUTH_BIND_USER")
	bindPassword := os.Getenv("LDAP_AUTH_BIND_PASSWORD")

	ldapUrl, set := os.LookupEnv("LDAP_AUTH_URI")
	if !set {
		return errors.New("env var LDAP_AUTH_URI must be set")
	}
	baseDn, set := os.LookupEnv("LDAP_AUTH_BASE_DN")
	if !set {
		return errors.New("env var LDAP_AUTH_BASE_DN must be set")
	}

	userFilter, set := os.LookupEnv("LDAP_AUTH_USER_FILTER")
	if !set {
		return errors.New("env var LDAP_AUTH_USER_FILTER must be set")
	}
	if !strings.Contains(userFilter, "%s") {
		return errors.New("env var LDAP_AUTH_USER_FILTER must contain a placeholder '%s' where username is placed, e.g. (uid=%s)")
	}

	startTlsString, set := os.LookupEnv("LDAP_AUTH_STARTTLS")
	useStartTls := strings.ToLower(startTlsString) == "true"

	displayNameAttr, set := os.LookupEnv("LDAP_AUTH_DISPLAY_NAME_ATTR")
	if !set {
		displayNameAttr = "displayName"
	}

	target.Uri = ldapUrl
	target.UseStartTLS = useStartTls
	target.UseBind = useBind
	target.Bind.BindUser = bindUsername
	target.Bind.Password = bindPassword
	target.BaseDN = baseDn
	target.UserFilter = userFilter
	target.DisplayNameAttr = displayNameAttr
	return nil
}

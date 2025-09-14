package ldap

import (
	"errors"
	"flag"
	"strings"
)

type BindConfig struct {
	BindUser string
	Password string
}

type Config struct {
	Uri               string
	UseStartTLS       bool
	UseBind           bool
	Bind              BindConfig
	BaseDN            string
	UserFilter        string
	DisplayNameAttr   string
	AuthenticateAdmin bool
}

type Cli struct {
	Flags      *flag.FlagSet
	LdapConfig Config
}

func NewConfigCli() *Cli {
	result := &Cli{}
	result.Flags = flag.NewFlagSet("", flag.ExitOnError)
	result.Flags.BoolVar(&result.LdapConfig.UseBind, "bind", false, "Enable Bind Authentication. Set bind-user and bind-password accordingly")
	result.Flags.StringVar(&result.LdapConfig.Bind.BindUser, "bind-user", "", "Perform Bind with the given DN")
	result.Flags.StringVar(&result.LdapConfig.Bind.Password, "bind-password", "", "Password for Bind DN")
	result.Flags.StringVar(&result.LdapConfig.Uri, "url", "", "URL of LDAP server, e.g. ldaps://foo.bar:636")
	result.Flags.StringVar(&result.LdapConfig.BaseDN, "base-dn", "", "Base DN to search, e.g. dc=foo,dc=bar")
	result.Flags.StringVar(&result.LdapConfig.UserFilter, "user-filter", "", "User filter to apply. Put %s for username, e.g. (uid=%s)")
	result.Flags.BoolVar(&result.LdapConfig.UseStartTLS, "starttls", false, "Use StartTLS to connect")
	result.Flags.StringVar(&result.LdapConfig.DisplayNameAttr, "displayname-attribute", "displayName", "LDAP attribute containing the user's display name")
	result.Flags.BoolVar(&result.LdapConfig.AuthenticateAdmin, "authenticate-admin", false, "Make any authenticated user an admin")
	return result
}

func (config *Cli) Validate() error {
	if !config.Flags.Parsed() {
		return errors.New("flags have not yet been parsed")
	}

	if "" == config.LdapConfig.Uri {
		return errors.New("url must not be empty")
	}
	if "" == config.LdapConfig.BaseDN {
		return errors.New("base-dn must not be empty")
	}
	if "" == config.LdapConfig.UserFilter {
		return errors.New("user-filter must not be empty")
	}
	if !strings.Contains(config.LdapConfig.UserFilter, "%s") {
		return errors.New("user-filter must contain a placeholder '%s' where username is placed, e.g. (uid=%s)")
	}
	if config.LdapConfig.UseBind && ("" == config.LdapConfig.Bind.BindUser || "" == config.LdapConfig.Bind.Password) {
		return errors.New("if bind is set, bind-user and bind-password must also be set")
	}

	return nil
}

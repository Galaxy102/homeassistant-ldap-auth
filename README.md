# homeassistant-ldap-auth

LDAP authentication command for HomeAssistant

## HowTo

1. Add the compiled result to your HomeAssistant install
2. Configure the command line auth provider in HomeAssistant (
   see [here](https://www.home-assistant.io/docs/authentication/providers/#command-line)) to use this tool.

### Example config
```yaml
homeassistant:
   auth_providers:
      - type: command_line
        command: /opt/hass-ldap-authenticator   # Assuming you copied the binary to this path
        args:
           - "--url"
           - ldaps://auth.foo.bar
           - "--base-dn"
           - dc=auth,dc=foo,dc=bar
           - "--bind"
           - "--bind-user"
           - cn=ldap-bind,ou=service,dc=auth,dc=foo,dc=bar
           - "--bind-password"
           - FooBarTopSecret
           - "--user-filter"
           - "(&(objectClass=user)(uid=%s))"
        meta: true
      - type: homeassistant
```

## Command Line arguments

| Name                             | Meaning                                                                     | Required                            |
|:---------------------------------|:----------------------------------------------------------------------------|:------------------------------------|
| `--url string`                   | URL of LDAP server                                                          | yes                                 |
| `--starttls`                     | Use StartTLS                                                                | no                                  |
| `--base-dn string`               | Base DN to search                                                           | yes                                 |
| `--bind`                         | Use authenticated Bind                                                      | no                                  |
| `--bind-user string`             | DN of Bind User                                                             | When authenticated bind is required |
| `--bind-password string`         | Password of Bind User                                                       | When authenticated bind is required |
| `--user-filter string`           | Filter to apply for search. Put `%s` where you expect `username`            | yes                                 |
| `--displayname-attribute string` | LDAP attribute that contains the name that should be shown in HomeAssistant | defaults to `displayName`           |

## Environment variables

| Name       | Meaning                  | Required             |
|:-----------|:-------------------------|:---------------------|
| `username` | Username to authenticate | Set by HomeAssistant |
| `password` | Password to authenticate | Set by HomeAssistant |

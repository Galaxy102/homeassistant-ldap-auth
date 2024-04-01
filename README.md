# homeassistant-ldap-auth

LDAP authentication command for HomeAssistant

## HowTo

1. Add the compiled result to your HomeAssistant install
2. Set the required environment variables (see below)
3. Configure the command line auth provider in HomeAssistant (
   see [here](https://www.home-assistant.io/docs/authentication/providers/#command-line)) to use this tool.

## Environment variables

| Name                          | Meaning                                                                     | Required                            |
|:------------------------------|:----------------------------------------------------------------------------|:------------------------------------|
| `username`                    | Username to authenticate                                                    | Set by HomeAssistant                |
| `password`                    | Password to authenticate                                                    | Set by HomeAssistant                |
| `LDAP_AUTH_URI`               | URL of LDAP server                                                          | yes                                 |
| `LDAP_AUTH_STARTTLS`          | Use StartTLS                                                                | no                                  |
| `LDAP_AUTH_BASE_DN`           | Base DN to search                                                           | yes                                 |
| `LDAP_AUTH_BIND_USER`         | DN of Bind User                                                             | When authenticated bind is required |
| `LDAP_AUTH_BIND_PASSWORD`     | Password of Bind User                                                       | When authenticated bind is required |
| `LDAP_AUTH_USER_FILTER`       | Filter to apply for search. Put `%s` where you expect `username`            | yes                                 |
| `LDAP_AUTH_DISPLAY_NAME_ATTR` | LDAP attribute that contains the name that should be shown in HomeAssistant | defaults to `displayName`           |

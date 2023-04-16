
# Example YAML:
```
MaxHistoryInMemory: 10
HTTPPort: 3000
HTTPHost: "0.0.0.0"
Gotify:
  GotifyURL:        "http://localhost:8080"
  ApplicationToken: "Aq9Br2z3.2qYcS9"
Targets:
- Name:                  "Test"
  ConnectionInformation: "https://www.pjotrs.nl"
  Checks:
  - Name:     "httpupcheck"
    Schedule: "*/10 * * * * *"
    Options:
      timeoutInSeconds: "3"
    DisableGotifyForSuccessfulCheck: false
  - Name:     "runcmdcheck"
    Schedule: "*/5 * * * * *"
    Options:
      cmd: "ping -c 1 127.0.0.1"
      printoutput: "false"
```

Field description
| Field                                                 | Description                                                                                               |
|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| MaxHistoryInMemory                                    | The amount of previous checks to keep in memory.                                                          |
| HTTPPort                                              | The HTTP port to listen on.                                                                               |
| HTTPHost                                              | The HTTP interface to listen on.                                                                          |
| Gotify.GotifyURL                                      | The URL to connect to Gotify.                                                                             |
| Gotify.ApplicationToken                               | The token used to connect to Gotify.                                                                      |
| Targets[].Name                                        | Name of the target, handy for reading the messages in Gotify and on the UI.                               |
| Targets[].ConnectionInformation                       | Connection information shared with the differend checks.                                                  |
| Targets[].Checks                                      | List of checks to perform on target.                                                                      |
| Targets[].Checks[].Name                               | Name of the check, this is a name of a check in the monitoring service, see below for available checks.   |
| Targets[].Checks[].Schedule                           | Cron schedule format (incl. seconds) defining how often a check needs to run.                             |
| Targets[].Checks[].Options.*                          | Options related to the check.                                                                             |
| Targets[].Checks[].DisableGotifyForSuccessfulCheck    | Optional setting for determining to always sending a message to Gotify, or only failed checks.            |

## Checks
### httpupcheck
Option fields:
- `timeoutInSeconds`, seconds that the HTTP request needs to finish in before timing out.

### runcmdcheck
Option fields:
- `cmd`, command to run.
- `printoutput`, print the results of running this command to stdout (log).
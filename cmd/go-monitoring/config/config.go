package config

type Check struct {
	Name                            string
	Schedule                        string
	Options                         map[string]string
	DisableGotifyForSuccessfulCheck bool
}

type Target struct {
	Name                  string
	ConnectionInformation string
	Checks                []Check
}

type GotifyConfig struct {
	GotifyURL        string
	ApplicationToken string
}

type Config struct {
	Targets []Target
	Gotify  *GotifyConfig
}

func LoadMockConfig() Config {
	return Config{
		// Gotify: &GotifyConfig{
		// 	GotifyURL:        "http://localhost:8080",
		// 	ApplicationToken: "Aq9Br2z3.2qYcS9",
		// },
		Targets: []Target{
			{
				Name:                  "Test",
				ConnectionInformation: "https://www.pjotrs.nl",
				Checks: []Check{
					{
						Name:     "httpupcheck",
						Schedule: "*/10 * * * * *",
						Options: map[string]string{
							"timeoutInSeconds": "3",
						},
						// DisableGotifyForSuccessfulCheck: false,
					},
				},
			},
		},
	}
}

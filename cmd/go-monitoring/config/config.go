package config

type Check struct {
	Name    string
	Options map[string]string
}

type Target struct {
	Name                  string
	ConnectionInformation string
	Checks                []Check
}

type Config struct {
	Targets []Target
}

func LoadMockConfig() Config {
	return Config{
		Targets: []Target{
			{
				Name:                  "Test",
				ConnectionInformation: "https://www.pjotrs.nl",
				Checks: []Check{
					{
						Name: "httpupcheck",
						Options: map[string]string{
							"timeoutInSeconds": "3",
						},
					},
				},
			},
		},
	}
}

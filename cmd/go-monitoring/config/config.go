package config

import (
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

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
	MaxHistoryInMemory uint
	HTTPPort           int
	HTTPHost           string
	Targets            []Target
	Gotify             *GotifyConfig
}

func LoadConfigFromPath(path string) Config {
	var res Config
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error reading config file '%s':\n", path)
		log.Fatal(err)
	}

	err = yaml.Unmarshal(content, &res)
	if err != nil {
		log.Printf("Error parsing config file '%s':\n", path)
		log.Fatal(err)
	}

	return res
}

func LoadMockConfig() Config {
	return Config{
		MaxHistoryInMemory: 5,
		HTTPPort:           3000,
		HTTPHost:           "0.0.0.0",
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

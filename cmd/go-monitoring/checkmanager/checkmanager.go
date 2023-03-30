package checkmanager

import (
	"errors"
	"log"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
)

type CheckManager struct {
	config config.Config
}

func NewCheckManager(config config.Config) CheckManager {
	return CheckManager{config: config}
}

func (cm *CheckManager) getCheckByName(name string) (check.Check, error) {
	for _, check := range getChecks() {
		if check.GetName() == name {
			return check, nil
		}
	}
	return nil, errors.New("Tried to get unkown check '%s'!")
}

func (cm *CheckManager) ValidateConfig() {
	checks := make(map[string]check.Check)
	log.Println("Known checks:")
	for _, check := range getChecks() {
		checks[check.GetName()] = check
		log.Printf(" - %s\n", check.GetName())
	}
	configValid := true
	for _, target := range cm.config.Targets {
		for _, targetCheck := range target.Checks {
			check, ok := checks[targetCheck.Name]
			if !ok {
				log.Fatalf("Unkown check given '%s'!", targetCheck.Name)
			}
			for _, validOption := range check.GetValidOptions() {
				res := validOption.Validate(targetCheck.Options)
				if !res.Success() {
					configValid = false
					res.Log()
				}
			}
		}
	}
	if !configValid {
		log.Fatalln("A not valid config file was provided exiting program.")
	}
}

func (cm *CheckManager) PerformAllChecks() error {
	for _, target := range cm.config.Targets {
		for _, targetCheck := range target.Checks {
			c, err := cm.getCheckByName(targetCheck.Name)
			if err != nil {
				return err
			}
			res := c.Perform(target.ConnectionInformation, targetCheck)
			res.Log()

		}
	}
	return nil
}

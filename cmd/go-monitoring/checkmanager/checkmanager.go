package checkmanager

import (
	"errors"
	"log"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/informer"
	"github.com/robfig/cron/v3"
)

type CheckManager struct {
	config              config.Config
	checkUpdateInformer informer.CheckUpdateInformer
}

func NewCheckManager(config config.Config, cui informer.CheckUpdateInformer) CheckManager {
	return CheckManager{config: config, checkUpdateInformer: cui}
}

func (cm *CheckManager) getCheckByName(name string) (check.Check, error) {
	for _, check := range getChecks() {
		if check.GetName() == name {
			return check, nil
		}
	}
	return nil, errors.New("Tried to get unkown check '%s'!")
}

func (cm *CheckManager) Run() {
	cr := cron.New(cron.WithSeconds())
	cr.Start()

	for _, target := range cm.config.Targets {
		for _, targetCheck := range target.Checks {
			cr.AddFunc(targetCheck.Schedule, func() {
				c, err := cm.getCheckByName(targetCheck.Name)
				if err != nil {
					log.Printf("Check with name '%s' does not exist! Please use a valid check name!", targetCheck.Name)
					return
				}
				res := c.Perform(target.ConnectionInformation, targetCheck)
				res.Log()
				cm.checkUpdateInformer.Inform(res, target, targetCheck)
			})
			_ = targetCheck
		}
	}
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

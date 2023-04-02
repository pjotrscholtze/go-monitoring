package repo

import (
	"errors"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
)

type TargetCheckEntity struct {
	Result entity.Result
	Target config.Target
	Check  config.Check
}

type targetCheckRepoInMemory struct {
	targetCheckEntities map[string]map[string][]TargetCheckEntity
	maxHistory          uint
}
type TargetCheckRepo interface {
	UpdateCheck(result entity.Result, target config.Target, check config.Check)
	List() []TargetCheckEntity
	ListHistory(targetName, checkName string) ([]TargetCheckEntity, error)
}

func (tcr *targetCheckRepoInMemory) ListHistory(targetName, checkName string) ([]TargetCheckEntity, error) {
	if _, ok := tcr.targetCheckEntities[targetName]; !ok {
		return nil, errors.New("Non existing target given")
	}
	if _, ok := tcr.targetCheckEntities[targetName][checkName]; !ok {
		return nil, errors.New("Non existing check given")
	}
	return tcr.targetCheckEntities[targetName][checkName], nil
}

func (tcr *targetCheckRepoInMemory) List() []TargetCheckEntity {
	results := make([]TargetCheckEntity, 0)
	for _, targetMap := range tcr.targetCheckEntities {
		for _, tce := range targetMap {
			results = append(results, tce[len(tce)-1])
		}
	}
	return results
}

func (tcr *targetCheckRepoInMemory) UpdateCheck(result entity.Result, target config.Target, check config.Check) {
	if tcr.targetCheckEntities[target.Name] == nil {
		tcr.targetCheckEntities[target.Name] = make(map[string][]TargetCheckEntity)
		tcr.targetCheckEntities[target.Name][check.Name] = make([]TargetCheckEntity, 0)
	}
	tcr.targetCheckEntities[target.Name][check.Name] = append(
		tcr.targetCheckEntities[target.Name][check.Name],
		TargetCheckEntity{
			Result: result,
			Target: target,
			Check:  check,
		})
	if len(tcr.targetCheckEntities[target.Name][check.Name]) > int(tcr.maxHistory) {
		tcr.targetCheckEntities[target.Name][check.Name] = tcr.targetCheckEntities[target.Name][check.Name][1:]
	}
}
func NewTargetCheckRepoInMemory(maxHistory uint) TargetCheckRepo {
	return &targetCheckRepoInMemory{
		targetCheckEntities: make(map[string]map[string][]TargetCheckEntity),
		maxHistory:          maxHistory,
	}
}

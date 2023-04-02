package main

import (
	"fmt"
	"time"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/checkmanager"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/informer"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/repo"
)

// "github.com/pjotrscholtze/go-buildserver/cmd/go-buildserver/process"

func main() {
	// process.StartProcessWithStdErrStdOutCallback("/bin/sh",
	// 	[]string{path.Join("/home/pjotr/go/src/github.com/pjotrscholtze/go-monitoring", "boot.sh")},
	// 	func(pt process.PipeType, t time.Time, s string) {
	// 		println(s)
	// 	})
	cui := informer.NewCheckUpdateInformer()
	tcr := repo.NewTargetCheckRepoInMemory(5)
	cui.RegisterListenerFunc(func(result entity.Result, target config.Target, check config.Check) {
		tcr.UpdateCheck(result, target, check)
		fmt.Printf("via informer, we got %s, %s, %s\n", result.Message(), target.Name, check.Name)
	})
	cm := checkmanager.NewCheckManager(config.LoadMockConfig(), cui)
	cm.ValidateConfig()
	cm.Run()
	// err := cm.PerformAllChecks()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	for {
		time.Sleep(time.Second)
		for _, tce := range tcr.List() {
			println(" from main")
			tce.Result.Log()
		}
	}
}

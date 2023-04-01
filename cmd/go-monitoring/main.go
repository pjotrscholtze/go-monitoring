package main

import (
	"fmt"
	"time"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/checkmanager"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/informer"
)

// "github.com/pjotrscholtze/go-buildserver/cmd/go-buildserver/process"

func main() {
	// process.StartProcessWithStdErrStdOutCallback("/bin/sh",
	// 	[]string{path.Join("/home/pjotr/go/src/github.com/pjotrscholtze/go-monitoring", "boot.sh")},
	// 	func(pt process.PipeType, t time.Time, s string) {
	// 		println(s)
	// 	})
	cui := informer.NewCheckUpdateInformer()
	cui.RegisterListenerFunc(func(result entity.Result, target config.Target, check config.Check) {
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
	}
}

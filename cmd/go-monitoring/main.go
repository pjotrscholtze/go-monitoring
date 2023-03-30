package main

import (
	"log"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/checkmanager"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
)

// "github.com/pjotrscholtze/go-buildserver/cmd/go-buildserver/process"

func main() {
	// process.StartProcessWithStdErrStdOutCallback("/bin/sh",
	// 	[]string{path.Join("/home/pjotr/go/src/github.com/pjotrscholtze/go-monitoring", "boot.sh")},
	// 	func(pt process.PipeType, t time.Time, s string) {
	// 		println(s)
	// 	})
	cm := checkmanager.NewCheckManager(config.LoadMockConfig())
	cm.ValidateConfig()
	err := cm.PerformAllChecks()
	if err != nil {
		log.Fatal(err.Error())
	}
}

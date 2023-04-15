package checkmanager

import (
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check/httpupcheck"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check/runcmdcheck"
)

func getChecks() []check.Check {
	return []check.Check{
		httpupcheck.NewHttpUpCheck(),
		runcmdcheck.NewRuncmdcheck(),
	}
}

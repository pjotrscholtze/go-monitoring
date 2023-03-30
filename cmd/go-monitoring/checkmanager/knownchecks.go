package checkmanager

import (
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check/httpupcheck"
)

func getChecks() []check.Check {
	return []check.Check{
		httpupcheck.NewHttpUpCheck(),
	}
}

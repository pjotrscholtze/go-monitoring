package httpupcheck

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
)

type httpUpCheck struct {
}

func NewHttpUpCheck() check.Check {
	return &httpUpCheck{}
}
func (hup *httpUpCheck) GetName() string {
	return "httpupcheck"
}
func (hup *httpUpCheck) GetValidOptions() []check.OptionValidator {
	return []check.OptionValidator{
		check.NewOptionValidator("timeoutInSeconds", true, check.Uint16),
	}
}

func (hup *httpUpCheck) Perform(address string, checkConfig config.Check) entity.Result {
	timeoutInSeconds, err := strconv.ParseUint(checkConfig.Options["timeoutInSeconds"], 10, 16)
	client := http.Client{
		Timeout: time.Duration(timeoutInSeconds) * time.Second,
	}
	beforeExec := time.Now()
	resp, err := client.Get(address)
	execDuration := time.Now().Sub(beforeExec)

	if err != nil {
		return entity.NewBadResultWithAttributes(
			hup.GetName(),
			err,
			"Failed to do HTTP request, got an exception.",
			map[string]interface{}{
				"addres": address,
			},
			execDuration,
		)
	}
	if resp.StatusCode != 200 {
		return entity.NewBadResult(
			hup.GetName(),
			nil,
			fmt.Sprintf("Got an error while making an HTTP request, statuscode != 200, statuscode = %d", resp.StatusCode),
			execDuration,
		)
	}

	return entity.NewOkResult(hup.GetName(), "Ok", execDuration)
}

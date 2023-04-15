package runcmdcheck

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pjotrscholtze/go-buildserver/cmd/go-buildserver/process"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/check"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
)

type runcmdcheck struct {
}

func NewRuncmdcheck() check.Check {
	return &runcmdcheck{}
}
func (hup *runcmdcheck) GetName() string {
	return "runcmdcheck"
}
func (hup *runcmdcheck) GetValidOptions() []check.OptionValidator {
	return []check.OptionValidator{
		check.NewOptionValidator("cmd", true, check.String),
		check.NewOptionValidator("printoutput", true, check.Bool),
	}
}

func (hup *runcmdcheck) Perform(address string, checkConfig config.Check) entity.Result {
	command := checkConfig.Options["cmd"]
	printoutput, _ := strconv.ParseBool(checkConfig.Options["printoutput"])

	beforeExec := time.Now()

	output := make([]string, 0)
	hadError := false
	parts := strings.Split(command, " ")
	process.StartProcessWithStdErrStdOutCallback(parts[0],
		parts[1:],
		func(pt process.PipeType, t time.Time, s string) {
			line := fmt.Sprintf("STDOUT: %s", s)
			if pt == process.STDERR {
				line = fmt.Sprintf("STDERR: %s\n", s)
				hadError = true
			}
			if printoutput {
				log.Println(line)
			}
			output = append(output, line)
		})
	execDuration := time.Now().Sub(beforeExec)

	if hadError {
		return entity.NewBadResultWithAttributes(
			hup.GetName(),
			errors.New("Error occured during execution"),
			strings.Join(output, "\n"),
			map[string]interface{}{
				"command": command,
			},
			execDuration,
		)
	}

	output = append([]string{"ok"}, output...)
	return entity.NewOkResult(hup.GetName(), strings.Join(output, "\n"), execDuration)
}

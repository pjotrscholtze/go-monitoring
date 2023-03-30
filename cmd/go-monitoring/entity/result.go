package entity

import (
	"fmt"
	"time"
)

type result struct {
	checkName     string
	error         error
	success       bool
	lastCheck     time.Time
	message       string
	attributes    map[string]interface{}
	executionTime time.Duration
}
type Result interface {
	Error() error
	Success() bool
	LastCheck() time.Time
	Message() string
	Log()
}

func (r *result) Error() error {
	return r.error
}
func (r *result) Success() bool {
	return r.success
}
func (r *result) LastCheck() time.Time {
	return r.lastCheck
}
func (r *result) Message() string {
	return r.message
}
func (r *result) Log() {

	errorMessage := ""
	if r.error != nil {
		errorMessage = " " + r.error.Error()
	}
	success := "did not pass"
	if r.success {
		success = "passed"
	}
	fmt.Printf("Check '%s' %s@%s, with message: %s.%s (took: %s)\n", r.checkName, success, r.lastCheck.Format(time.RFC3339), r.message, errorMessage, r.executionTime.String())
}
func NewOkResult(checkName string, message string, executionTime time.Duration) Result {
	return NewOkResultWithAttributes(checkName, message, map[string]interface{}{}, executionTime)
}
func NewOkResultWithAttributes(checkName string, message string, attributes map[string]interface{}, executionTime time.Duration) Result {
	return NewResult(
		checkName,
		nil,
		true,
		message,
		attributes,
		executionTime,
	)
}
func NewBadResult(checkName string, error error, message string, executionTime time.Duration) Result {
	return NewBadResultWithAttributes(checkName, error, message, map[string]interface{}{}, executionTime)
}
func NewBadResultWithAttributes(checkName string, error error, message string, attributes map[string]interface{}, executionTime time.Duration) Result {
	return NewResult(
		checkName,
		error,
		false,
		message,
		attributes,
		executionTime,
	)
}
func NewResult(checkName string, error error, success bool, message string, attributes map[string]interface{}, executionTime time.Duration) Result {
	return &result{
		checkName:     checkName,
		error:         error,
		success:       success,
		lastCheck:     time.Now(),
		message:       message,
		attributes:    attributes,
		executionTime: executionTime,
	}
}

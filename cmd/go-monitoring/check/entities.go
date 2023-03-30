package check

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
)

type Datatype int64

const (
	Uint8 Datatype = iota
	Uint16
	Uint32
	Uint64
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
	Bool
	String
)

func hasDatatype(needle Datatype, haystack []Datatype) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func isUint(val string, datatype Datatype) bool {
	var err error
	if datatype == Uint8 {
		_, err = strconv.ParseUint(val, 10, 8)
	} else if datatype == Uint16 {
		_, err = strconv.ParseUint(val, 10, 16)
	} else if datatype == Uint32 {
		_, err = strconv.ParseUint(val, 10, 32)
	} else if datatype == Uint64 {
		_, err = strconv.ParseUint(val, 10, 64)
	} else {
		log.Fatal("Unkown data type given! Expected an unsigned integer type!")
	}
	return err == nil
}
func isInt(val string, datatype Datatype) bool {
	var err error
	if datatype == Int8 {
		_, err = strconv.ParseInt(val, 10, 8)
	} else if datatype == Int16 {
		_, err = strconv.ParseInt(val, 10, 16)
	} else if datatype == Int32 {
		_, err = strconv.ParseInt(val, 10, 32)
	} else if datatype == Int64 {
		_, err = strconv.ParseInt(val, 10, 64)
	} else {
		log.Fatal("Unkown data type given! Expected a signed integer type!")
	}
	return err == nil
}
func isFloat(val string, datatype Datatype) bool {
	var err error
	if datatype == Float32 {
		_, err = strconv.ParseFloat(val, 32)
	} else if datatype == Float64 {
		_, err = strconv.ParseFloat(val, 64)
	} else {
		log.Fatal("Unkown data type given! Expected a float type!")
	}
	return err == nil
}
func isBool(val string) bool {
	_, err := strconv.ParseBool(val)
	return err == nil
}

type optionValidator struct {
	name     string
	required bool
	datatype Datatype
}
type OptionValidator interface {
	Validate(option map[string]string) entity.Result
}

func (ov *optionValidator) Validate(option map[string]string) entity.Result {
	val, ok := option[ov.name]
	if ov.required && !ok {
		return entity.NewBadResultWithAttributes(
			"OptionValidationCheck",
			nil,
			fmt.Sprintf("Required value '%s' is missing!", ov.name),
			map[string]interface{}{
				"name":     ov.name,
				"required": ov.required,
				"found":    false,
			},
			time.Duration(0),
		)
	}
	valid := false
	if hasDatatype(ov.datatype, []Datatype{Uint8, Uint16, Uint32, Uint64}) {
		valid = isUint(val, ov.datatype)
	} else if hasDatatype(ov.datatype, []Datatype{Int8, Int16, Int32, Int64}) {
		valid = isInt(val, ov.datatype)
	} else if hasDatatype(ov.datatype, []Datatype{Float32, Float64}) {
		valid = isFloat(val, ov.datatype)
	} else if ov.datatype == Bool {
		valid = isBool(val)
	} else if ov.datatype == String {
		valid = true
	} else {
		log.Fatal("Unkown datatype given!")
	}
	if !valid {
		return entity.NewBadResultWithAttributes(
			"OptionValidationCheck",
			nil,
			fmt.Sprintf("Value '%s' is has incorrect datatype!", ov.name),
			map[string]interface{}{
				"name":              ov.name,
				"required":          ov.required,
				"found":             false,
				"expected_datatype": ov.datatype,
			},
			time.Duration(0),
		)
	}

	return entity.NewOkResult(
		"OptionValidationCheck",
		fmt.Sprintf("Value '%s' has no validation errors!", ov.name),
		time.Duration(0),
	)
}

func NewOptionValidator(name string, required bool, datatype Datatype) OptionValidator {
	return &optionValidator{
		name:     name,
		required: required,
		datatype: datatype,
	}
}

type Check interface {
	GetName() string
	GetValidOptions() []OptionValidator
	Perform(address string, checkConfig config.Check) entity.Result
}
